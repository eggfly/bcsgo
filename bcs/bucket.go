package bcs

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Bucket struct {
	bcs  *BCS
	Name string `json:"bucket_name"`
}

func (this *Bucket) getUrl() string {
	return this.bcs.simpleSign(GET, this.Name, "/")
}

func (this *Bucket) ListObjects(prefix string, start, limit int) ([]Object, error) {
	params := url.Values{}
	params.Set("start", string(start))
	params.Set("limit", string(limit))
	link := this.getUrl() + "&" + params.Encode()
	data, err := this.bcs.httpClient.Get(link)
	fmt.Println(string(data))
	if err != nil {
		return nil, err
	} else {
		pList := &[]Object{}
		err := json.Unmarshal(data, pList)
		list := *pList
		for i, _ := range list {
			list[i].bucket = this
		}
		return list, err
	}
	return nil, nil
}
