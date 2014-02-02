package bcsgo

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

func (this *HttpClient) Get(url string) ([]byte, error) {
	resp, err := this.client.Get(url)
	return this.handleResponse(resp, err)
}

func (this *HttpClient) handleResponse(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}
