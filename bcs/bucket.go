package bcs

import (
// "fmt"
)

type bucketStruct struct {
	bcs  *BCS
	Name string `json:"bucket_name"`
}
type Bucket *bucketStruct

func newBucket(bcs *BCS, bucketName string) Bucket {
	return &bucketStruct{bcs, bucketName}
}
