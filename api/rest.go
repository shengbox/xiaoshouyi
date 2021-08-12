package api

import (
	"github.com/go-resty/resty/v2"
	"log"
)

func (a *API) Description(objectName string) (*Description, error) {
	var result Description
	response, err := resty.New().R().
		SetResult(&result).
		SetHeader("Authorization", a.Tokener.Token()).
		Get(Endpoint + "/rest/data/v2.0/xobjects/" + objectName + "/description")
	if Debug {
		log.Printf("response = %s", response.String())
	}
	return &result, err
}

//CreateAccount 创建客户
func (a *API) CreateAccount(account *Account) (*AccountResp, error) {
	body := map[string]interface{}{
		"data": account,
	}
	var result AccountResp
	response, err := resty.New().R().SetBody(body).
		SetResult(&result).
		SetHeader("Authorization", a.Tokener.Token()).
		Post(Endpoint + "/rest/data/v2.0/xobjects/account")
	if err != nil {
		return nil, err
	}
	if Debug {
		log.Printf("response = %s", response.String())
	}
	return &result, err
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
