package youzango

import (
    "github.com/smartystreets/goconvey/convey"
    "testing"
)

var code = ""
var clientId = ""
var clientSecret = ""

func TestClient_GetToken(t *testing.T) {
    convey.Convey("测试获取token和刷新token", t, func() {
        var refresh_token string
        client := &Client{
            ClientId: clientId,
            ClientSecret: clientSecret,
        }
        convey.Convey("测试获取token", func() {
            token, err := client.GetToken(code)
            convey.So(err, convey.ShouldBeNil)
            convey.So(token, convey.ShouldNotBeNil)
            refresh_token = token.RefreshToken
        })
        convey.Convey("测试刷新token", func() {
            token, err := client.RefreshToken(refresh_token)
            convey.So(err, convey.ShouldBeNil)
            convey.So(token, convey.ShouldNotBeNil)
        })
    })
}