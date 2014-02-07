package bcsgo

import (
	// "fmt"
	"testing"
)

var bucketForSuperfileTest *Bucket

func TestSuperfileInit(t *testing.T) {
	bucket := createBucketTempForTest(t)
	bucketForSuperfileTest = bucket

	createTestFile(_TEST_NAME, 1024)
}

func TestSuperfilePutAndDelete(t *testing.T) {
	DEBUG = true
	DEBUG_REQUEST_BODY = true
	bucket := bucketForSuperfileTest
	// todo file name with blank char
	putFile := func(path, localFile string) *Object {
		testObj := bucket.Object(path)
		testObj, err := testObj.PutFile(localFile, ACL_PUBLIC_READ)
		if err != nil {
			t.Error(err)
		}
		if testObj.AbsolutePath != path {
			t.Error("testObj.AbsolutePath != path", testObj.AbsolutePath, path)
		}
		return testObj
	}
	deleteFile := func(testObj *Object) {
		deleteErr := testObj.Delete()
		if deleteErr != nil {
			t.Error(deleteErr)
		}
	}
	obj1 := putFile("/testDir/part1.txt", _TEST_NAME)
	obj2 := putFile("/testDir/part2.txt", _TEST_NAME)
	s := bucket.Superfile("/testDir/merged.txt", []*Object{obj1, obj2})
	err := s.Put()
	if err != nil {
		t.Error(err)
	}

	deleteFile(obj1)
	deleteFile(obj2)
	deleteFile(&s.Object)
}

func TestSuperfileFinalize(t *testing.T) {
	deleteBucketForTest(t, bucketForSuperfileTest)
	bucketForSuperfileTest = nil
	deleteTestFile(_TEST_NAME)
}
