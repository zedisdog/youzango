package youzango

import (
    "github.com/zedisdog/youzango/utils"
)

type Request struct {
    Name string
    Version string
    Query map[string]string
    JsonData []byte
}

func newMethod(name string, version string, accessToken string, request interface{}) *Request {

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

func newPointDecreaseRequest(request map[string]interface{}, accessToken string) *Request {
    return newMethod(
        "youzan.crm.customer.points.decrease",
        "3.1.0",
        accessToken,
        request,
    )
}

func newSalesmanAccountsRequest(request map[string]interface{}, accessToken string) *Request {
    return newMethod(
        "youzan.salesman.accounts.get",
        "3.0.0",
        accessToken,
        request,
    )
}

func newTradeRequest(request map[string]interface{}, accessToken string) *Request {
    return newMethod(
        "youzan.trade.get",
        "4.0.0",
        accessToken,
        request,
    )
}

func newGetOpenIdByMobileRequest(request map[string]interface{}, accessToken string) *Request {
    if request["country_code"] == nil {
        request["country_code"] = "86"
    }
    return newMethod(
        "youzan.user.weixin.openid.get",
        "3.0.0",
        accessToken,
        request,
    )
}

func newUsersWeixinFollowerRequest(request map[string]interface{}, accessToken string) *Request {
    return newMethod(
        "youzan.users.weixin.follower.get",
        "3.0.0",
        accessToken,
        request,
    )
}