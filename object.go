package bcsgo

import (
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
	Size         int64  `json:"size,string"`
	ParentDir    string `json:"parent_dir"`
	IsDir        string `json:"is_dir"`
	MDatetime    string `json:"mdatetime"`
	RefKey       string `json:"ref_key"`
	ContentMD5   string `json:"content_md5"`
}

func (this *Object) getUrl() string {
	return this.bucket.bcs.restUrl(GET, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) getACLUrl() string {
	return this.getUrl() + "&acl=1"
}
func (this *Object) putUrl() string {
	return this.bucket.bcs.restUrl(PUT, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) putACLUrl() string {
	return this.putUrl() + "&acl=1"
}
func (this *Object) headUrl() string {
	return this.bucket.bcs.restUrl(HEAD, this.bucket.Name, this.AbsolutePath)
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
	link := this.headUrl()
	resp, _, err := this.bucket.bcs.httpClient.Head(link)
	err = mergeResponseError(err, resp)
	if err != nil {
		return err
	} else {
		this.Size = resp.ContentLength
		this.ContentMD5 = resp.Header.Get(HEADER_CONTENT_MD5)
		this.VersionKey = resp.Header.Get(HEADER_VERSION)
		return nil
	}
}
func (this *Object) PutFile(localFile string) (*Object, error) {
	return this.putFileInner(localFile, "")
}
func (this *Object) PutFileWithACL(localFile, acl string) (*Object, error) {
	return this.putFileInner(localFile, acl)
}
func (this *Object) putFileInner(localFile string, acl string) (*Object, error) {
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
	err = mergeResponseError(err, resp)
	if err != nil {
		return nil, err
	} else {
		this.ContentMD5 = resp.Header.Get(HEADER_CONTENT_MD5)
		this.VersionKey = resp.Header.Get(HEADER_VERSION) // TODO check version json and this
		this.Size, _ = strconv.ParseInt(resp.Header.Get(HEADER_FILESIZE), 10, 64)
		return this, err
	}
}
func (this *Object) Delete() error {
	link := this.deleteUrl()
	resp, _, err := this.bucket.bcs.httpClient.Delete(link)
	return mergeResponseError(err, resp)
}
func (this *Object) refStr() string {
	return fmt.Sprintf(`bs://%s%s`, this.bucket.Name, this.AbsolutePath)
}
func (this *Object) GetACL() (string, error) {
	link := this.getACLUrl()
	resp, data, err := this.bucket.bcs.httpClient.Get(link)
	err = mergeResponseError(err, resp)
	return string(data), err
}
func (this *Object) SetACL(acl string) error {
	link := this.putACLUrl()
	modifyHeader := func(header *http.Header) {
		header.Set(HEADER_ACL, acl)
	}
	resp, _, err := this.bucket.bcs.httpClient.Put(link, nil, 0, modifyHeader)
	return mergeResponseError(err, resp)
}
func (this *Object) CopyTo(target *Object) (*Object, error) {
	// take care of this, target put url
	link := target.putUrl()
	modifyHeader := func(header *http.Header) {
		header.Set(HEADER_COPY_SOURCE, this.refStr())
	}
	resp, _, err := this.bucket.bcs.httpClient.Put(link, nil, 0, modifyHeader)
	err = mergeResponseError(err, resp)
	if err != nil {
		return nil, err
	} else {
		target.ContentMD5 = resp.Header.Get(HEADER_CONTENT_MD5)
		target.VersionKey = resp.Header.Get(HEADER_VERSION) // TODO check version json and this
		target.Size, _ = strconv.ParseInt(resp.Header.Get(HEADER_FILESIZE), 10, 64)
		return target, err
	}
}
