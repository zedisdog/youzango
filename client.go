package youzango

import (
    "errors"
    "fmt"
    "github.com/json-iterator/go"
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

func (c *Client) GetToken(tokenRequest *TokenRequest) (*Token, error) {
    json := jsoniter.ConfigCompatibleWithStandardLibrary
    jsonData, err := json.Marshal(tokenRequest)

    if err != nil {
        return nil, err
    }

    var tokenResponse TokenResponse
    err = request(tokenApi, "", "", nil, jsonData, &tokenResponse, c.IsLog)
    
    if err != nil {
        return nil, err
    }

    if tokenResponse.Code != 200 {
        return nil, errors.New(fmt.Sprintf("code:%d, message:%s", tokenResponse.Code, tokenResponse.Message))
    }
    
    return &tokenResponse.Data, nil
}

func (c *Client) RefreshToken(refreshTokenRequest *RefreshTokenRequest) (*Token, error) {
    json := jsoniter.ConfigCompatibleWithStandardLibrary
    jsonData, err := json.Marshal(refreshTokenRequest)

    if err != nil {
        return nil, err
    }

    var tokenResponse TokenResponse
    err = request(tokenApi, "", "", nil, jsonData, &tokenResponse, c.IsLog)

    if err != nil {
        return nil, err
    }

    if tokenResponse.Code != 200 {
        return nil, errors.New(fmt.Sprintf("code:%d, message:%s", tokenResponse.Code, tokenResponse.Message))
    }

    return &tokenResponse.Data, nil
}

func (c *Client) Trade(request *api.TradeRequest) (*api.Trade, error) {
    var tradeResponse api.TradeResponse
    err := c.requestApi(api.NewTradeMethod(request, c.accessToken), &tradeResponse)

    if err != nil {
        return nil, err
    }

    return &tradeResponse.Data, nil
}

func (c *Client) DecreasePoint(request *api.PointDecreaseRequest) (*api.PointDecreaseResult, error) {
    var response api.PointDecreaseResponse
    err := c.requestApi(api.NewPointDecreaseMethod(request, c.accessToken), &response)

    if err != nil {
        return nil, err
    }

    return &response.Data, nil
}

func (c *Client) requestApi(method *api.Method, response interface{}) error {
    return request(normalApi, method.Name, method.Version, method.Query, method.JsonData, response, c.IsLog)
}