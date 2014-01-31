package bcs

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&http.Client{}}
}

func (this *HttpClient) Get(url string) (string, error) {
	resp, err := this.client.Get(url)
	return this.handleResponse(resp, err)
}

func (this *HttpClient) handleResponse(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		return string(body), err
	}
}
