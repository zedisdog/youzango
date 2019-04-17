package youzango

import "youzango/api"

type Client struct {
    ClientId string
    ClientSecret string
}

func (c *Client) GetToken(code string) (*Token, error) {
    var tokenResponse TokenResponse
    err := request(tokenApi, "", "", nil, []byte{}, &tokenResponse)
    
    if err != nil {
        return nil, err
    }
    
    return &tokenResponse.Data, nil
}

func (c *Client) RefreshToken(refreshToken string) (*Token, error) {
    var tokenResponse TokenResponse
    err := request(tokenApi, "", "", nil, []byte{}, &tokenResponse)

    if err != nil {
        return nil, err
    }

    return &tokenResponse.Data, nil
}

func (c *Client) Trade(method *api.Method) (*api.Trade, error) {
    var tradeResponse api.TradeResponse
    err := request(normalApi, method.Name, method.Version, method.Query, method.JsonData, &tradeResponse)

    if err != nil {
        return nil, err
    }

    return &tradeResponse.Data, nil
}