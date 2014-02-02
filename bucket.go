package bcsgo

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

func (this *Bucket) Object(absolutePath string) *Object {
	o := Object{}
	o.bucket = this
	o.AbsolutePath = absolutePath
	return &o
}

func (this *Bucket) ListObjects(prefix string, start, limit int) (*ObjectCollection, error) {
	params := url.Values{}
	params.Set("start", string(start))
	params.Set("limit", string(limit))
	link := this.getUrl() + "&" + params.Encode()
	data, err := this.bcs.httpClient.Get(link)
	fmt.Println(string(data))
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
