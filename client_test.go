package youzango

import (
    "github.com/smartystreets/goconvey/convey"
    "testing"
)

var code = ""
var clientId = ""
var clientSecret = ""
var redirectUri = ""

func TestClient_GetToken(t *testing.T) {
    convey.Convey("测试获取token和刷新token", t, func() {
        client := &Client{IsLog: true}

        token, err := client.GetToken(NewTokenRequest(clientId, clientSecret, code, redirectUri))
        convey.So(err, convey.ShouldBeNil)
        convey.So(token, convey.ShouldNotBeNil)

        token, err = client.RefreshToken(NewRefreshTokenRequest(clientId, clientSecret, token.RefreshToken))
        convey.So(err, convey.ShouldBeNil)
        convey.So(token, convey.ShouldNotBeNil)
        convey.So(token.AccessToken, convey.ShouldNotBeEmpty)
    })
}