package api

import (
    "github.com/uniplaces/carbon"
)

type CanGetCarbon interface {
    ToCarbon(prop string) (*carbon.Carbon, error)
}

type CanGetPrice interface {
    GetPrice(prop string) (int, error)
}

type Method struct {
    Name string
    Version string
    Query map[string]string
    JsonData []byte
}

func NewMethod(accessToken string) *Method {
    return &Method{Query: map[string]string{"access_token": accessToken}}
}