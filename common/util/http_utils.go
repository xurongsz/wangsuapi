package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"openApi-authentication/common/model"
	"strings"
)

func Call(requestMsg model.HttpRequestMsg) string{
	client := &http.Client{}
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		// handle error
	}
	for key := range requestMsg.Headers {
		req.Header.Set(key, requestMsg.Headers [key])
	}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(resp)
	return string(body)
}
