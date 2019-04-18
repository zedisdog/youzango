package api

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestCreatePointDecreaseResponse(t *testing.T) {
    convey.Convey("测试json解析", t, func() {
        data := `{"code":200,"data":{"is_success":"true"},"success":true,"message":"successful"}`
        json := jsoniter.ConfigCompatibleWithStandardLibrary
        var response PointDecreaseResponse
        err := json.Unmarshal([]byte(data), &response)

        convey.So(err, convey.ShouldBeNil)
        convey.So(response.Data.IsSuccess, convey.ShouldEqual, "true")
        convey.So(response.Success, convey.ShouldBeTrue)
    })
}