package test

import (
	"github.com/shengbox/xiaoshouyi/api"
	"log"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	xsy := api.New("", "", "", "", "")
	account := &api.Account{
		OwnerId:     1,
		AccountName: "测试",
		EntityType:  int64(1449331636765394),
		DimDepart:   int64(1466747499905749),
		CreatedAt:   time.Now().Unix() * 1000,
		UpdatedAt:   time.Now().Unix() * 1000,
		HighSeaId:   int64(1449333235581632),
		Phone:       "18666666666",
		Address:     "华润大厦",
		Comment:     "呵呵",
		Level:       int(1),
	}
	err := xsy.CreateAccount(account)
	log.Panicln(err)
}
