package api

import (
    "youzango/utils"
)

type PointDecreaseRequest struct {
    // 帐号ID
    AccountId string `json:"account_id"`
    // 帐号类型 与帐户ID配合使用:
    // 2=粉丝(原fansId)
    // 3:手机号
    // 4:三方帐号(原open_user_id)
    // 6:微信open_id
    AccountType int `json:"account_type"`
    // 积分变动值
    Points int `json:"points"`
    // 用于幂等支持（幂等时效三个月, 超过三个月的相同值调用不保证幂等）
    BizValue string `json:"biz_value"`
    // 积分变动原因
    Reason string `json:"reason"`
}

func NewPointDecreaseMethod(request *PointDecreaseRequest, accessToken string) *Method {
    method := NewMethod(accessToken)
    method.Name = "youzan.crm.customer.points.decrease"
    method.Version = "3.1.0"

    jsonData, err := utils.BuildJson(request)
    if err != nil {
        panic(err)
    }
    method.JsonData = jsonData

    return method
}