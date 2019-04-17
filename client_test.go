package youzango

import (
    "github.com/smartystreets/goconvey/convey"
    "testing"
)

var code = "4c52379b1e80ecc7cd554843ecf49588"
var clientId = "be9a06911147e350b0"
var clientSecret = "dfee7097a9b77c394c2439c838e95ef8"
var redirectUri = "http://devbbdapi.ffuture.cn/api/yz/newcallback"

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