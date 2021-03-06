package main

import (
	"encoding/json"
	"fmt"
	dproxy "github.com/koron/go-dproxy"
	"io/ioutil"
	"net/http"
)

type ThingName string

type ThingWorxClient struct {
	URL    string
	AppKey string
}

func (tw *ThingWorxClient) Properties(name ThingName) (dproxy.Proxy, error) {
	url := fmt.Sprintf("%s/Things/%s/Properties/", tw.URL, string(name))
	if tw.AppKey != "" {
		url += "?appKey=" + tw.AppKey
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	//client := http.Client{Timeout: 10}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	js, err := ioutil.ReadAll(res.Body)

	var v interface{}
	json.Unmarshal(js, &v)

	return dproxy.New(v).M("rows").A(0), nil
}
