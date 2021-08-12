package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"log"
	"reflect"
	"strings"
	"sync"
)

var Debug = false
var Endpoint = "https://api-tencent.xiaoshouyi.com"
var GrantType = "password"

type API struct {
	clientId     string
	clientSecret string
	username     string
	password     string
	redirectUri  string
	Tokener      *Tokener
}

func (a *API) fetchToken() (string, int64) {
	var result TokenInfo
	resp, _ := resty.New().R().SetResult(&result).
		SetQueryParams(map[string]string{
			"client_id":     a.clientId,
			"client_secret": a.clientSecret,
			"grant_type":    GrantType,
			"username":      a.username,
			"password":      a.password,
			"redirect_uri":  a.redirectUri,
		}).
		Get(Endpoint + "/oauth2/token")
	log.Println(resp.String())
	return result.AccessToken, result.IssuedAt
}

func New(clientId, clientSecret, username, password, redirectUri string) *API {
	api := &API{
		clientId:     clientId,
		clientSecret: clientSecret,
		username:     username,
		password:     password,
		redirectUri:  redirectUri,
	}
	token := &Tokener{
		mutex:        &sync.RWMutex{},
		tokenFetcher: api.fetchToken,
	}
	api.Tokener = token
	return api
}

func (a *API) XOQLRecords(page int, t, where, v interface{}) error {
	table := strings.ToLower(reflect.TypeOf(t).Name())
	field := strings.Join(getFieldName(reflect.TypeOf(t)), ",")
	sql := fmt.Sprintf("SELECT %s FROM %s %s LIMIT %d,%d", field, table, cast.ToString(where), (page-1)*100, 100)
	return a.XOQLQuery(sql, v)
}

func (a *API) XOQLQuery(sql string, v interface{}) error {
	var resp XOQLResp
	response, err := resty.New().R().SetResult(&resp).SetHeader("Authorization", a.Tokener.Token()).
		SetQueryParam("q", sql).
		Get(Endpoint + "/rest/data/v2/query")
	if Debug {
		log.Printf("sql = %s", sql)
		log.Printf("response = %s", response.String())
	}
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Msg)
	}
	bt, _ := json.Marshal(resp.Result.Records)
	err = json.Unmarshal(bt, v)
	return err
}

func getFieldName(t reflect.Type) []string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Type.Kind() == reflect.Struct {
			for _, it := range getFieldName(t.Field(i).Type) {
				result = append(result, it)
			}
		} else {
			result = append(result, t.Field(i).Name)
		}
	}
	return result
}
