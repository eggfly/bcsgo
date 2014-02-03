package bcsgo

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&http.Client{}}
}

func (this *HttpClient) Get(url string) (*http.Response, []byte, error) {
	resp, err := this.client.Get(url)
	respData, err := this.handleResponseContent(resp, err)
	return resp, respData, err
}
func (this *HttpClient) Put(url string, data io.Reader, size int64) (*http.Response, []byte, error) {
	req, err := http.NewRequest(PUT, url, data)
	if err != nil {
		return nil, nil, err
	}
	req.ContentLength = size
	dump, dumpErr := httputil.DumpRequest(req, false)
	log.Println(string(dump), dumpErr)
	old := time.Now()
	resp, err := this.client.Do(req)
	log.Println(time.Now().Sub(old))
	respData, err := this.handleResponseContent(resp, err)
	log.Println(string(respData), err)
	return resp, respData, err
}

func (this *HttpClient) handleResponseContent(resp *http.Response, err error) ([]byte, error) {
	dump, dumpErr := httputil.DumpResponse(resp, true)
	log.Println(string(dump), dumpErr)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
}
