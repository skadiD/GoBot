package request

import (
	"bytes"
	"encoding/json"
	"github.com/fexli/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Post POST请求
func Post(url string, data map[string]interface{}, headers map[string]string) (content []byte) {
	client := &http.Client{}
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("请求网络失败：", err))
	}
	defer req.Body.Close()
	if headers != nil {
		for key, header := range headers {
			req.Header.Set(key, header)
		}
	}
	resp, err2 := client.Do(req)
	if err2 != nil {
		logger.RootLogger.Error(logger.WithContent("请求网络失败：", err2))
	}
	defer resp.Body.Close()
	content, _ = ioutil.ReadAll(resp.Body)
	return
}
func PostEncode(urls string, data map[string]string, headers map[string]string) (content []byte) {
	client := &http.Client{}
	q := url.Values{}
	if data != nil {
		for key, val := range data {
			q.Add(key, val)
		}
	}
	req, _ := http.NewRequest("POST", urls, strings.NewReader(q.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, _ := client.Do(req)
	content, _ = ioutil.ReadAll(resp.Body)
	return
}
