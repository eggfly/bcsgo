package bcsgo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type BCS struct {
	ak, sk     string
	httpClient *HttpClient
}

func NewBCS(ak, sk string) *BCS {
	return &BCS{ak, sk, NewHttpClient()}
}

func (this *BCS) ListBuckets() ([]*Bucket, error) {
	link := this.getUrl()
	_, data, err := this.httpClient.Get(link)
	if err != nil {
		return nil, err
	} else {
		list := []*Bucket{}
		err := json.Unmarshal(data, &list)
		for i, _ := range list {
			list[i].bcs = this
		}
		return list, err
	}
}
func (this *BCS) Bucket(bucketName string) *Bucket {
	return &Bucket{this, bucketName}
}

func (this *BCS) getUrl() string {
	return this.restUrl(GET, "", "/")
}
func (this *BCS) restUrl(method, bucket, object string) string {
	return this.urlWithSign(method, bucket, object, "", "", "")
}
func (this *BCS) restUrlExtra(method, bucket, object, time, ip, size string) string {
	return this.urlWithSign(method, bucket, object, time, ip, size)
}
func (this *BCS) urlWithSign(method, bucket, object, time, ip, size string) string {
	return fmt.Sprintf("%s?sign=%s", this.urlWithoutSign(bucket, object), this.sign(method, bucket, object, time, ip, size))
}
func (this *BCS) urlWithoutSign(bucket, object string) string {
	return fmt.Sprintf("%s/%s%s", BCS_HOST, bucket, "/"+url.QueryEscape(object[1:]))
	// return fmt.Sprintf("%s/%s%s", BCS_HOST, bucket, "/"+object[1:])
}
func (this *BCS) sign(m, b, o, t, i, s string) string {
	flag := ""
	ss := ""
	flag += "M"
	ss += "Method=" + m + "\n"
	flag += "B"
	ss += "Bucket=" + b + "\n"
	flag += "O"
	ss += "Object=" + o + "\n"
	if t != "" {
		flag += "T"
		ss += "Time=" + t + "\n"
	}
	if i != "" {
		flag += "I"
		ss += "Ip=" + i + "\n"
	}
	if s != "" {
		flag += "S"
		ss += "Size=" + s + "\n"
	}
	ss = flag + "\n" + ss
	h := func(sk, body string) string {
		hash := hmac.New(sha1.New, []byte(sk))
		hash.Write([]byte(body))
		digest := hash.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(digest)
		sign = strings.TrimSpace(sign)
		sign = url.QueryEscape(sign)
		return sign
	}
	sign := h(this.sk, ss)
	final := fmt.Sprintf(
		"%s:%s:%s",
		flag,
		this.ak,
		sign)
	if t != "" {
		final += "&time=" + t
	}
	if i != "" {
		final += "&ip=" + i
	}
	if s != "" {
		final += "&size=" + s
	}
	return final
}
