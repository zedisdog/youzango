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
        token, err := client.GetToken(map[string]string{
            "client_id": clientId,
            "client_secret": clientSecret,
            "code": code,
            "redirect_uri": redirectUri,
            "authorize_type": "authorization_code",
        })
        convey.So(err, convey.ShouldBeNil)
        convey.So(token, convey.ShouldNotBeNil)

        token, err = client.RefreshToken(map[string]string{
            "client_id": clientId,
            "client_secret": clientSecret,
            "refresh_token": token.Get("refresh_token").String(),
            "authorize_type": "refresh_token",
        })
        convey.So(err, convey.ShouldBeNil)
        convey.So(token, convey.ShouldNotBeNil)
        convey.So(token.Get("access_token").String(), convey.ShouldNotBeEmpty)
    })
}