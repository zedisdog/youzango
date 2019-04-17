package utils

import (
    "errors"
    "github.com/json-iterator/go"
    "github.com/uniplaces/carbon"
    "reflect"
    "strconv"
    "time"
)

func BuildJson(data interface{}) ([]byte, error) {
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    return json.Marshal(data)
}

func ParseJson(data []byte, s interface{}) error {
    var json = jsoniter.ConfigCompatibleWithStandardLibrary
    return json.Unmarshal(data, s)
}

func toCarbon(s string) (*carbon.Carbon, error) {
    loc, _ := time.LoadLocation("Local")
    t, err := carbon.CreateFromFormat("2006-01-02 15:04:05", s, loc.String())
    return t, err
}

func GetCarbon(s interface{}, prop string) (*carbon.Carbon, error) {
    rv := reflect.ValueOf(s)
    pv := rv.FieldByName(prop)
    if pv.IsValid() && !pv.IsNil() && pv.Type().String() == "string" {
        return toCarbon(pv.String())
    }

    return nil, errors.New("prop is invalid")
}

func convertPrice(s string) (int, error) {
    f, err := strconv.ParseFloat(s, 32)
    if err != nil {
        return 0, err
    }

    return int(f*100), nil
}

func GetPrice(s interface{}, prop string) (int, error) {
    rv := reflect.ValueOf(s)
    pv := rv.FieldByName(prop)
    if pv.IsValid() && !pv.IsNil() && pv.Type().String() == "string" {
        return convertPrice(pv.String())
    }

    return 0, errors.New("prop is invalid")
}