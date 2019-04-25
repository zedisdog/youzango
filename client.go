package youzango

import (
    "errors"
    "fmt"
    "github.com/json-iterator/go"
    "github.com/tidwall/gjson"
    "youzango/api"
)

type Client struct {
    accessToken string
    refreshToken string
    IsLog bool
}

func (c *Client) SetAccessToken(accessToken string) {
    c.accessToken = accessToken
}

func (c *Client) SetRefreshToken(refreshToken string) {
    c.refreshToken = refreshToken
}

func (c *Client) GetToken(tokenRequest *TokenRequest) (*gjson.Result, error) {
    var response *gjson.Result
    json := jsoniter.ConfigCompatibleWithStandardLibrary
    jsonData, err := json.Marshal(tokenRequest)

    if err != nil {
        return nil, err
    }

    response, err = request(tokenApi, "", "", nil, jsonData, c.IsLog)
    
    if err != nil {
        return nil, err
    }

    if response.Get("code").Int() != 200 {
        return nil, errors.New(fmt.Sprintf("code:%d, message:%s", response.Get("code").Int(), response.Get("message").String()))
    }

    result := response.Get("data")
    return &result, nil
}

func (c *Client) RefreshToken(refreshTokenRequest *RefreshTokenRequest) (*gjson.Result, error) {
    var response *gjson.Result
    json := jsoniter.ConfigCompatibleWithStandardLibrary
    jsonData, err := json.Marshal(refreshTokenRequest)

    if err != nil {
        return nil, err
    }

    response, err = request(tokenApi, "", "", nil, jsonData, c.IsLog)

    if err != nil {
        return nil, err
    }

    if response.Get("code").Int() != 200 {
        return nil, errors.New(fmt.Sprintf("code:%d, message:%s", response.Get("code").Int(), response.Get("message").String()))
    }

    result := response.Get("data")
    return &result, nil
}

func (c *Client) Trade(request *api.TradeRequest) (*gjson.Result, error) {
    result, err := c.requestApi(api.NewTradeMethod(request, c.accessToken))

    if err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) DecreasePoint(request *api.PointDecreaseRequest) (*gjson.Result, error) {
    result, err := c.requestApi(api.NewPointDecreaseMethod(request, c.accessToken))

    if err != nil {
        return nil, err
    }

    return result, nil
}

func (c *Client) requestApi(method *api.Method) (*gjson.Result, error) {
    response,err := request(normalApi, method.Name, method.Version, method.Query, method.JsonData, c.IsLog)

    if err != nil {
        return nil, err
    }

    if response.Get("code").Int() != 200 {
        return nil, errors.New(fmt.Sprintf("code:%d, message:%s", response.Get("code").Int(), response.Get("message").String()))
    }

    result := response.Get("data")
    return &result, nil
}