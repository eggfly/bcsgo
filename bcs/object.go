package bcs

import (
// "fmt"
)

type Object struct {
	bucket *Bucket
	Name   string
}

func newObject(bucket *Bucket, objectName string) *Object {
	return &Object{bucket, objectName}
}
