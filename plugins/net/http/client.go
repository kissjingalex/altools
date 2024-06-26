package nhttp

import (
	cjson "altools/plugins/json"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	nurl "net/url"
	"strings"
	"time"
)

var (
	// NoKeepAliveClient is http client without keep alive
	NoKeepAliveClient = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	defaultHTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 2048,
			IdleConnTimeout:     time.Minute * 5,
		},
	}
	defaultTimeout    = 500
	defaultRetryCount = 2
)

// PostRaw PostRaw
// params []{"毫秒超时","重试次数"}
func PostRaw(client *http.Client, url string, header http.Header, reqBody interface{}, params ...int) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	timeOut, retryCount := genDefaultParams(params...)
	for i := 0; i < retryCount; i++ {
		data, _, _, err = do(client, http.MethodPost, url, header, reqBody, timeOut)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}
	return data, err
}

// PostWithUnmarshal do http get with unmarshal
func PostWithUnmarshal(client *http.Client, url string, header http.Header, reqBody interface{}, resp interface{}, params ...int) error {
	data, err := PostRaw(client, url, header, reqBody, params...)
	if err != nil {
		return err
	}
	// for no resp needed request.
	if resp == nil {
		return nil
	}
	// for big int
	decoder := cjson.JSON.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	err = decoder.Decode(resp)
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}
	return err
}

// GetRaw get http raw
func GetRaw(client *http.Client, url string, header http.Header, reqBody interface{}, params ...int) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	timeOut, retryCount := genDefaultParams(params...)
	for i := 0; i < retryCount; i++ {
		data, _, _, err = do(client, http.MethodGet, url, header, reqBody, timeOut)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}
	return data, err
}

// ReqGet get http raw
func ReqGet(client *http.Client, requestUrl string, header http.Header, requestBody interface{}, params ...int) (data []byte, rspHeader http.Header, httpCode int, err error) {
	timeout, retry := genDefaultParams(params...)
	for i := 0; i < retry; i++ {
		data, rspHeader, httpCode, err = do(client, http.MethodGet, requestUrl, header, requestBody, timeout)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Printf("ReqGet [%s] err: %s\n", err)
	}
	return data, rspHeader, httpCode, err
}

// Head head 请求
func Head(client *http.Client, url string, header http.Header, params ...int) (http.Header, int, error) {
	var (
		httpStatus int
		rspHeader  http.Header
		err        error
	)
	timeOut, retryCount := genDefaultParams(params...)
	for i := 0; i < retryCount; i++ {
		_, rspHeader, httpStatus, err = do(client, http.MethodHead, url, header, nil, timeOut)
		if err == nil {
			break
		}
	}
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}
	return rspHeader, httpStatus, err
}

// GetWithUnmarshal do http get with unmarshal
func GetWithUnmarshal(client *http.Client, url string, header http.Header, reqBody interface{}, resp interface{}, params ...int) error {
	data, err := GetRaw(client, url, header, reqBody, params...)
	if err != nil {
		return err
	}
	// for no resp needed request.
	if resp == nil {
		return nil
	}
	// for big int
	decoder := cjson.JSON.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	err = decoder.Decode(resp)
	if err != nil {
		fmt.Printf("err:%s, url:%s, respData:%s\n", err, url, string(data))
	}
	return err
}

func genDefaultParams(params ...int) (int, int) {
	timeOut, retryCount := defaultTimeout, defaultRetryCount
	switch {
	case len(params) >= 2:
		timeOut, retryCount = params[0], params[1]
	case len(params) >= 1:
		timeOut = params[0]
	}
	return timeOut, retryCount
}

func do(client *http.Client, method string, url string, header http.Header, reqBody interface{}, timeOut int) ([]byte, http.Header, int, error) {
	if client == nil {
		client = defaultHTTPClient
	}
	var reader io.Reader
	switch v := reqBody.(type) {
	case nurl.Values:
		reader = strings.NewReader(v.Encode())
	case []byte:
		reader = bytes.NewBuffer(v)
	case string:
		reader = strings.NewReader(v)
	case io.Reader:
		reader = v
	default:
		buff := &bytes.Buffer{}
		err := cjson.JSON.NewEncoder(buff).Encode(v)
		if err != nil {
			return nil, nil, -1, err
		}
		reader = buff
	}
	if (method == http.MethodGet || method == http.MethodHead) && reqBody == nil {
		reader = nil
	}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, nil, -1, err
	}
	if header != nil {
		req.Header = header
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeOut))
	defer cancelFunc()
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, -1, err // TODO maybe should define ctx timeout in package errs
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, -1, err
	}

	return data, resp.Header, resp.StatusCode, nil
}
