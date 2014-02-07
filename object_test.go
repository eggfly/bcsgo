package bcsgo

import (
	// "fmt"
	"os"
	"strings"
	"testing"
)

var bucketForObjectTest *Bucket

const (
	_LARGER_NAME = "256KB.data"
	_TEST_NAME   = "test.txt"
)

func createTestFile(filename string, size int) {
	file, _ := os.Create(filename)
	file.WriteString(strings.Repeat(" ", size))
	file.Close()
}
func deleteTestFile(filename string) {
	os.Remove(filename)
}

func TestObjectInit(t *testing.T) {
	bucket := createBucketTempForTest(t)
	bucketForObjectTest = bucket

	createTestFile(_LARGER_NAME, 256*1024)
	createTestFile(_TEST_NAME, 1024)
}

func TestObjectPutAndDeleteObject(t *testing.T) {
	bucket := bucketForObjectTest
	// todo file name with blank char
	path := "/testDir/testwithblank.txt"
	testObj := bucket.Object(path)
	testObj, err := testObj.PutFileWithACL(_TEST_NAME, ACL_PUBLIC_READ)
	if err != nil {
		t.Error(err)
	}
	if testObj.AbsolutePath != path {
		t.Error("testObj.AbsolutePath != path", testObj.AbsolutePath, path)
	}

	// expectedPublicLink := fmt.Sprintf("%s/%s%s", BCS_HOST, bucket.Name, path)
	// publicLink := testObj.PublicLink()
	// if expectedPublicLink != publicLink {
	// 	t.Error("expectedPublicLink != publicLink", expectedPublicLink, publicLink)
	// }

	headErr := testObj.Head()
	if headErr != nil {
		t.Error(headErr)
	}
	if testObj.ContentMD5 == "" || testObj.VersionKey == "" {
		t.Error("Info after HEAD is not ok!")
	}

	deleteErr := testObj.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestObjectLargerSingleFile(t *testing.T) {
	bucket := bucketForObjectTest
	obj := bucket.Object("/larger.data")
	obj, err := obj.PutFile(_LARGER_NAME)
	if err != nil {
		t.Error(err)
	}
	deleteErr := obj.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestObjectFinalize(t *testing.T) {
	deleteBucketForTest(t, bucketForObjectTest)
	bucketForObjectTest = nil
	deleteTestFile(_LARGER_NAME)
	deleteTestFile(_TEST_NAME)
}
