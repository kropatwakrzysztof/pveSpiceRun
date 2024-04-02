package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

func sendApiRequest(method string, url string, headers map[string]string, data []byte, insecure bool) (map[string]interface{}, error) {

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	var tr *http.Transport = nil
	if insecure {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		if err == nil {
			err = errors.New(response.Status)
		}
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if len(body) > 0 {
		var responseData ResponseData
		json.Unmarshal(body, &responseData)
		if responseData.Data == nil {
			return nil, errors.New("no json in response body")
		}
		return responseData.Data.(map[string]interface{}), nil
	}
	// everything OK - returning empty map
	return map[string]interface{}{}, nil
}
