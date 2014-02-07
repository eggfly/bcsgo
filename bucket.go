package bcsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Bucket struct {
	bcs  *BCS
	Name string `json:"bucket_name"`
}

func (this *Bucket) getUrl() string {
	return this.bcs.restUrl(GET, this.Name, "/")
}
func (this *Bucket) getACLUrl() string {
	return this.getUrl() + "&acl=1"
}
func (this *Bucket) putUrl() string {
	return this.bcs.restUrl(PUT, this.Name, "/")
}
func (this *Bucket) putACLUrl() string {
	return this.putUrl() + "&acl=1"
}
func (this *Bucket) deleteUrl() string {
	return this.bcs.restUrl(DELETE, this.Name, "/")
}
func (this *Bucket) CreateWithACL(acl string) error {
	return this.createInner(acl)
}
func (this *Bucket) Create() error {
	return this.createInner("")
}
func (this *Bucket) createInner(acl string) error {
	link := this.putUrl()
	var modifyHeader func(*http.Header)
	if acl != "" {
		modifyHeader = func(header *http.Header) {
			header.Set(HEADER_ACL, acl)
		}
	}
	resp, _, err := this.bcs.httpClient.Put(link, nil, 0, modifyHeader)
	if resp.StatusCode != http.StatusOK {
		err = errors.New("request not ok, status: " + strconv.Itoa(resp.StatusCode))
	}
	return err
}
func (this *Bucket) Delete() error {
	link := this.deleteUrl()
	resp, _, err := this.bcs.httpClient.Delete(link)
	if resp.StatusCode != http.StatusOK {
		err = errors.New("request not ok, status: " + strconv.Itoa(resp.StatusCode))
	}
	return err
}
func (this *Bucket) Object(absolutePath string) *Object {
	if absolutePath[0] != '/' {
		panic("object name (aka absolute path) must start with '/'")
	}
	o := Object{}
	o.bucket = this
	o.AbsolutePath = absolutePath
	return &o
}
func (this *Bucket) Superfile(absolutePath string, objects []*Object) *Superfile {
	if absolutePath[0] != '/' {
		panic("object name (aka absolute path) must start with '/'")
	}
	s := Superfile{}
	s.bucket = this
	s.AbsolutePath = absolutePath
	s.Objects = objects
	return &s
}

func (this *Bucket) ListObjects(prefix string, start, limit int) (*ObjectCollection, error) {
	params := url.Values{}
	params.Set("start", string(start))
	params.Set("limit", string(limit))
	link := this.getUrl() + "&" + params.Encode()
	_, data, err := this.bcs.httpClient.Get(link)
	if err != nil {
		return nil, err
	} else {
		var objectsInfo ObjectCollection
		err := json.Unmarshal(data, &objectsInfo)
		fmt.Println(objectsInfo)
		if err != nil {
			return nil, err
		} else {
			for i, _ := range objectsInfo.Objects {
				objectsInfo.Objects[i].bucket = this
			}
			return &objectsInfo, nil
		}
	}
}
func (this *Bucket) GetACL() (string, error) {
	link := this.getACLUrl()
	_, data, err := this.bcs.httpClient.Get(link)
	return string(data), err
}
func (this *Bucket) SetACL(acl string) error {
	link := this.putACLUrl()
	modifyHeader := func(header *http.Header) {
		header.Set(HEADER_ACL, acl)
	}
	_, _, err := this.bcs.httpClient.Put(link, nil, 0, modifyHeader)
	return err
}
