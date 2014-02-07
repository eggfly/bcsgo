package bcsgo

import (
	// "fmt"
	"testing"
	"time"
)

var bucketForSuperfileTest *Bucket

func TestSuperfileInit(t *testing.T) {
	bucket := createBucketTempForTest(t)
	bucketForSuperfileTest = bucket

	createTestFile(_TEST_NAME, 1*1024*1024)
}

func TestSuperfilePutAndDelete(t *testing.T) {
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
	dupFileTimes := func(absPath string, obj *Object, times int) *Superfile {
		repeats := make([]*Object, 0)
		for i := 0; i < times; i++ {
			repeats = append(repeats, obj)
		}
		s := bucket.Superfile(absPath, repeats)
		err := s.Put()
		if err != nil {
			t.Error(err)
		}
		return s
	}
	obj := putFile("/testDir/test.txt", _TEST_NAME)
	// DEBUG_REQUEST_BODY = true
	s := dupFileTimes("/testDir/test.txt", obj, 1024)

	deleteFile(&s.Object)
}

func TestSuperfileFinalize(t *testing.T) {
	time.Sleep(time.Second)
	deleteBucketForTest(t, bucketForSuperfileTest)
	bucketForSuperfileTest = nil
	deleteTestFile(_TEST_NAME)
}
