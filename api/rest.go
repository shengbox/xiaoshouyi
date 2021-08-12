package api

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"log"
)

func (a *API) Description(objectName string) (*Description, error) {
	var result Description
	response, err := resty.New().R().SetResult(&result).
		SetHeader("Authorization", a.Tokener.Token()).
		Get(Endpoint + "/rest/data/v2.0/xobjects/" + objectName + "/description")
	if Debug {
		log.Printf("response = %s", response.String())
	}
	return &result, err
}

//CreateAccount 创建客户
func (a *API) CreateAccount(account *Account) (*CommonResp, error) {
	body := map[string]interface{}{"data": account}
	var result CommonResp
	response, err := resty.New().R().SetBody(body).SetResult(&result).
		SetHeader("Authorization", a.Tokener.Token()).
		Post(Endpoint + "/rest/data/v2.0/xobjects/account")
	if err != nil {
		return nil, err
	}
	if Debug {
		bt, _ := json.Marshal(body)
		log.Printf("request url  = %s", "/rest/data/v2.0/xobjects/account")
		log.Printf("request body = %s", string(bt))
		log.Printf("response = %s", response.String())
	}
	return &result, err
}

//UpdateAccount 更新客户信息
func (a *API) UpdateAccount(account *Account) (*CommonResp, error) {
	body := map[string]interface{}{
		"data": account,
	}
	var result CommonResp
	response, err := resty.New().R().SetResult(&result).SetBody(body).
		SetHeader("Authorization", a.Tokener.Token()).
		Patch(Endpoint + "/rest/data/v2.0/xobjects/account/" + cast.ToString(account.Id))
	if err != nil {
		return nil, err
	}
	if Debug {
		bt, _ := json.Marshal(body)
		log.Printf("request url  = %s", "/rest/data/v2.0/xobjects/account/"+cast.ToString(account.Id))
		log.Printf("request body = %s", string(bt))
		log.Printf("response = %s", response.String())
	}
	if result.Code != "200" {
		return nil, errors.New(result.Msg)
	}
	return &result, err
}

//CreateActivityRecord 创建活动记录
func (a *API) CreateActivityRecord(record *ActivityRecord) (*CommonResp, error) {
	body := map[string]interface{}{"data": record}
	var result CommonResp
	response, err := resty.New().R().SetResult(&result).SetBody(body).
		SetHeader("Authorization", a.Tokener.Token()).
		Post(Endpoint + "/rest/data/v2.0/xobjects/activityrecord")
	if err != nil {
		return nil, err
	}
	if Debug {
		bt, _ := json.Marshal(body)
		log.Printf("request url  = %s", "/rest/data/v2.0/xobjects/activityrecord")
		log.Printf("request body = %s", string(bt))
		log.Printf("response = %s", response.String())
	}
	if result.Code != "200" {
		return nil, errors.New(result.Msg)
	}
	return &result, err
}

//Transfer 转移客户
func (a *API) Transfer(accountId, targetOwnerId string) error {
	body := map[string]interface{}{
		"id":            accountId,
		"targetOwnerId": targetOwnerId,
		"highSeaStatus": "4",
	}
	response, err := resty.New().R().SetBody(body).
		SetHeader("Authorization", a.Tokener.Token()).
		Post(Endpoint + "/data/v1/objects/account/transfer")
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("request url  = %s", "/data/v1/objects/account/transfer")
		log.Printf("response = %s", response.String())
	}
	return err
}

// Release 客户退回公海池
func (a *API) Release(accountId string) error {
	body := map[string]interface{}{
		"reasonType": 1388121,
		"accountIds": accountId,
	}
	response, err := resty.New().R().SetBody(body).Post(Endpoint + "/rest/data/v2.0/xobjects/account/actions/release")
	if Debug {
		bt, _ := json.Marshal(body)
		log.Printf("request url  = %s", "/rest/data/v2.0/xobjects/activityrecord")
		log.Printf("request body = %s", string(bt))
		log.Printf("response = %s", response.String())
	}
	return err
}

//GetDocument 文档信息
func (a *API) GetDocument(dataId string) ([]Document, error) {
	var result DocumentResp
	response, err := resty.New().R().SetHeader("Authorization", a.Tokener.Token()).
		SetResult(&result).
		SetQueryParam("belongId", "1").
		SetQueryParam("dataId", dataId).
		Get(Endpoint + "/data/v1/objects/document/list")
	if Debug {
		log.Printf("response = %s", response.String())
	}
	if err != nil {
		return nil, err
	}
	return result.Record, nil
}
