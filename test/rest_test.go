package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shengbox/xiaoshouyi/api"
)

const (
	client       = ""
	clientSecret = ""
	username     = ""
	password     = ""
)

func TestCreateAccount(t *testing.T) {
	api.Debug = true
	xsy := api.New(client, clientSecret, username, password, "")
	account := &api.Account{
		OwnerId:     int64(1481294232240872),
		AccountName: "测试测试有限公司",
		EntityType:  int64(1449331636765394),
		DimDepart:   int64(1466747499905749),
		CreatedAt:   time.Now().Unix() * int64(1000),
		UpdatedAt:   time.Now().Unix() * int64(1000),
		HighSeaId:   int64(1449333235581632),
		Phone:       "18666666666",
		Address:     "华润大厦",
		Comment:     "呵呵",
		Level:       1,
	}
	_, _ = xsy.CreateAccount(account)
}

func TestDescription(t *testing.T) {
	api.Debug = true
	xsy := api.New(client, clientSecret, username, password, "")
	data, _ := xsy.Description("account")
	for _, v := range data.Data.Fields {
		fmt.Println(v.ApiKey, v.ItemType, "//", v.Label)
	}
}
