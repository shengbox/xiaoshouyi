package api

import (
	"github.com/go-resty/resty/v2"
	"log"
)

func (a *API) CreateAccount(account *Account) error {
	body := map[string]interface{}{
		"data": account,
	}
	response, err := resty.New().R().SetBody(body).
		SetHeader("Authorization", a.Tokener.AccessToken).
		Post(Endpoint + "/rest/data/v2.0/xobjects/account")
	if err != nil {
		return nil
	}
	log.Printf("response = %s", response.String())
	return err
}

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
