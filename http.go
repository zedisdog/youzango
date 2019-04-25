package youzango

import (
    "bytes"
    "github.com/pasztorpisti/qs"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "log"
    "net/http"
)

const tokenApi string = "https://open.youzanyun.com/auth/token"
const normalApi string = "https://open.youzanyun.com/api"

func httpPost(url string, jsonData []byte) ([]byte, error) {
    rsp, err := http.Post(
        url,
        "application/json",
        bytes.NewReader(jsonData),
    )
    defer rsp.Body.Close()

    if err != nil {
        return []byte{}, err
    }

    result, err := ioutil.ReadAll(rsp.Body)

    if err != nil {
        return []byte{}, err
    }

    return result, nil
}

func request(baseUrl string, method string, version string, query map[string]string, jsonData []byte, isLog bool) (*gjson.Result, error) {
    var rsp []byte
    var err error
    if baseUrl == tokenApi {
        if isLog {
            log.Println("+request url:", baseUrl, "data:", jsonData)
        }
        rsp, err = httpPost(baseUrl, jsonData)
        if err != nil {
            return nil, err
        }
    } else {

        var url string
        url, err = buildUrl(baseUrl, method, version, query)
        if err != nil {
            return nil, err
        }

        if isLog {
            log.Println("+request url:", url, "data:", jsonData)
        }
        rsp, err = httpPost(url, jsonData)
        if err != nil {
            return nil, err
        }
    }

    if isLog {
        log.Println("-response", string(rsp))
    }

    result := gjson.ParseBytes(rsp)
    return &result, nil
}

func buildUrl(baseUrl string, method string, version string, query map[string]string) (string, error) {
    var buffer bytes.Buffer
    buffer.WriteString(baseUrl)
    buffer.WriteString("/")
    buffer.WriteString(method)
    buffer.WriteString("/")
    buffer.WriteString(version)

    queryStr, err := qs.Marshal(query)
    if err != nil {
        return "", err
    }

    buffer.WriteString("?")
    buffer.WriteString(queryStr)

    return buffer.String(), nil
}