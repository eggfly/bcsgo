package bcsgo

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Object struct {
	bucket       *Bucket
	VersionKey   string `json:"version_key"`
	AbsolutePath string `json:"object"`
	Superfile    string `json:"superfile"`
	Size         uint64 `json:"size,string"`
	ParentDir    string `json:"parent_dir"`
	IsDir        string `json:"is_dir"`
	MDatetime    string `json:"mdatetime"`
	RefKey       string `json:"ref_key"`
	ContentMD5   string `json:"content_md5"`
}

func (this *Object) getUrl() string {
	return this.bucket.bcs.restUrl(GET, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) putUrl() string {
	return this.bucket.bcs.restUrl(PUT, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) deleteUrl() string {
	return this.bucket.bcs.restUrl(DELETE, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) Link() string {
	return this.getUrl()
}
func (this *Object) PublicLink() string {
	return this.bucket.bcs.urlWithoutSign(this.bucket.Name, this.AbsolutePath)
}
func (this *Object) Head() error {
	return nil
}
func (this *Object) PutFile(localFile string, acl string) (*Object, error) {
	link := this.putUrl()
	file, err := os.Open(localFile)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var modifyHeader func(header *http.Header) = nil
	if acl != "" {
		modifyHeader = func(header *http.Header) {
			header.Set(HEADER_ACL, acl)
		}
	}
	resp, _, err := this.bucket.bcs.httpClient.Put(link, file, fileInfo.Size(), modifyHeader)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	} else {
		this.ContentMD5 = resp.Header["Content-Md5"][0]
		this.VersionKey = resp.Header["X-Bs-Version"][0] // TODO check version json and this
		this.Size, _ = strconv.ParseUint(resp.Header["X-Bs-File-Size"][0], 10, 64)
		return this, err
	}
}
func (this *Object) Delete() error {
	link := this.deleteUrl()
	resp, _, err := this.bucket.bcs.httpClient.Delete(link)
	if resp.StatusCode != http.StatusOK {
		err = errors.New("request not ok, status: " + strconv.Itoa(resp.StatusCode))
	}
	return err
}
func (this *Object) refStr() string {
	return fmt.Sprintf(`bs://%s%s`, this.bucket.Name, this.AbsolutePath)
}
