package youzango

import (
    "youzango/utils"
)

type Method struct {
    Name string
    Version string
    Query map[string]string
    JsonData []byte
}

func NewMethod(name string, version string, accessToken string, request interface{}) *Method {

    method := &Method{
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

func NewPointDecreaseMethod(request map[string]interface{}, accessToken string) *Method {
    return NewMethod(
        "youzan.crm.customer.points.decrease",
        "3.1.0",
        accessToken,
        request,
    )
}

func NewSalesmanAccountsMethod(request map[string]interface{}, accessToken string) *Method {
    return NewMethod(
        "youzan.salesman.accounts.get",
        "3.0.0",
        accessToken,
        request,
    )
}

func NewTradeMethod(request map[string]interface{}, accessToken string) *Method {
    return NewMethod(
        "youzan.trade.get",
        "4.0.0",
        accessToken,
        request,
    )
}

func NewGetOpenIdByMobileMethod(request map[string]interface{}, accessToken string) *Method  {
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