package api

import (
    "youzango/utils"
)

func NewTradeMethod(request *TradeRequest, accessToken string) *Method {
    jsonData, err := utils.BuildJson(request)
    if err != nil {
        panic(err)
    }
    method := NewMethod(accessToken)
    method.Name = "youzan.trade.get"
    method.Version = "4.0.0"
    method.JsonData = jsonData
    return method
}

type TradeRequest struct {
    Tid string `json:"tid"`
}