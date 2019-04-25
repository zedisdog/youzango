package utils

import (
    "github.com/json-iterator/go"
)

func BuildJson(data interface{}) ([]byte, error) {
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    return json.Marshal(data)
}