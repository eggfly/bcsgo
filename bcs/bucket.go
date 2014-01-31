package bcs

import (
// "fmt"
)

type Bucket struct {
	bcs  *BCS
	Name string `json:"bucket_name"`
}

func (this Bucket) ListObjects() []Object {
	return nil
}
