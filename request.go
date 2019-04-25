package youzango

import (
    "youzango/utils"
)

type Request struct {
    Name string
    Version string
    Query map[string]string
    JsonData []byte
}

func NewMethod(name string, version string, accessToken string, request interface{}) *Request {

    method := &Request{
        Name: name,
        Version: version,
        Query: map[string]string{"access_token": accessToken},
    }
    jsonData, err := utils.BuildJson(request)
    if err != nil {
        panic(err)
    }
    method.JsonData = jsonData

    return method
}

func NewPointDecreaseRequest(request map[string]interface{}, accessToken string) *Request {
    return NewMethod(
        "youzan.crm.customer.points.decrease",
        "3.1.0",
        accessToken,
        request,
    )
}

func NewSalesmanAccountsRequest(request map[string]interface{}, accessToken string) *Request {
    return NewMethod(
        "youzan.salesman.accounts.get",
        "3.0.0",
        accessToken,
        request,
    )
}

func NewTradeRequest(request map[string]interface{}, accessToken string) *Request {
    return NewMethod(
        "youzan.trade.get",
        "4.0.0",
        accessToken,
        request,
    )
}

func NewGetOpenIdByMobileRequest(request map[string]interface{}, accessToken string) *Request {
    if request["country_code"] == nil {
        request["country_code"] = "86"
    }
    return NewMethod(
        "youzan.user.weixin.openid.get",
        "3.0.0",
        accessToken,
        request,
    )
}