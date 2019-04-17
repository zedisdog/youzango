package youzango

import (
    "errors"
    "fmt"
    "github.com/json-iterator/go"
    "youzango/api"
)

type Client struct {
    IsLog bool
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

func (c *Client) Trade(method *api.Method) (*api.Trade, error) {
    var tradeResponse api.TradeResponse
    err := request(normalApi, method.Name, method.Version, method.Query, method.JsonData, &tradeResponse, c.IsLog)

    if err != nil {
        return nil, err
    }

    return &tradeResponse.Data, nil
}