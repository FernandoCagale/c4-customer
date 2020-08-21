package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type NotifyHTTP struct {
	address string
	client  *http.Client
}

func NewNotifyHTTP() *NotifyHTTP {
	return &NotifyHTTP{
		address: os.Getenv("ADDRESS_TYPE_HTTP"),
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (notify *NotifyHTTP) GetNotify(headers map[string]string, ) ([]string, error) {
	fmt.Println("TYPE - HTTP")

	req, err := http.NewRequest("GET", notify.address+"/types/001", nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	response, err := notify.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var types *Types
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&types); err != nil {
		return nil, err
	}

	fmt.Println(types)

	return types.Types, nil
}
