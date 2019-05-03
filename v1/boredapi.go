package boredapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ApiRoot = "https://www.boredapi.com/api/"
)

var (
	defaultClient = http.Client{
		Timeout: time.Second * 3,
	}
)

func SetHttpClient(client *http.Client) {
	defaultClient = *client
}

type query interface {
	Endpoint() string
	Params() map[string]string
}

type apiError struct {
	Error *string `json:"error"`
}

func queryApi(query query, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, ApiRoot+query.Endpoint(), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	for key, value := range query.Params() {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	res, err := defaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	var apiError apiError
	if err := json.Unmarshal(body, &apiError); err != nil {
		return err
	}
	if apiError.Error != nil {
		return errors.New(*apiError.Error)
	}
	if err := json.Unmarshal(body, out); err != nil {
		return err
	}
	return nil
}
