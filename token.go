package youzango

import "youzango/utils"

type CanBuildJson interface {
    ToJson() ([]byte, error)
}

// TokenRequest token请求结构体
// ClientId: 应用id
// ClientSecret: 应用secret
// AuthorizeType: 授权类型
// Code: 商家code
// RedirectUri: 工具型应用接收授权码的回调地址
type TokenRequest struct {
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    AuthorizeType string `json:"authorize_type"`
    Code string `json:"code"`
    RedirectUri string `json:"redirect_uri"`
}

func (t *TokenRequest) ToJson() ([]byte, error) {
    return utils.BuildJson(t)
}

func NewTokenRequest(clientId string, clientSecret string, code string, redirectUri string) *TokenRequest {
    return &TokenRequest{
        ClientId: clientId,
        ClientSecret: clientSecret,
        Code: code,
        RedirectUri: redirectUri,
        AuthorizeType: "authorization_code",
    }
}

// RefreshTokenRequest 刷新token请求结构体
// AuthorizeType: 授权方式（固定为 “refresh_token”）
// RefreshToken: 用于刷新令牌 access_token 的refresh_token（过期时间：28 天）
// ClientId: 有赞云颁发给开发者的应用ID
// ClientSecret: 有赞云颁发给开发者的应用Secret
type RefreshTokenRequest struct {
    AuthorizeType string `json:"authorize_type"`
    RefreshToken string `json:"refresh_token"`
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
}

func NewRefreshTokenRequest(clientId string, clientSecret string, refreshToken string) *RefreshTokenRequest {
    return &RefreshTokenRequest{
        AuthorizeType: "refresh_token",
        RefreshToken: refreshToken,
        ClientId: clientId,
        ClientSecret: clientSecret,
    }
}

// TokenRequest token请求结构体
// Success: 是否成功获取 token
// Code: 成功固定为:200,错误码不同
// Data: token 信息
// Data.AccessToken: 用于调用 API 的 access_token，有效7天；access_token失效前可通过refresh_token刷新获取新的access_token，有效期仍是7天
// Data.ExpiresIn: access_token 的有效时长，时间戳（过期时间：7天）
// Data.Scope: access_token 最终的访问范围
// Data.RefreshToken: 用于延长 access_token 有效时间的刷新令牌（过期时间：28 天），在刷新后access_token会返回新的refresh_token
// Code: 商家code
// Message: 出现异常时，返回错误信息
type TokenResponse struct {
    Success bool `json:"success"`
    Code int `json:"code"`
    Data Token `json:"data"`
    Message string `json:"message"`
}

type Token struct {
    AccessToken string `json:"access_token"`
    ExpiresIn string `json:"expires_in"`
    Scope string `json:"scope"`
    RefreshToken string `json:"refresh_token"`
}