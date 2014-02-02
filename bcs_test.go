package bcsgo

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	bcs := NewBCS("vYlphQiwbhVz67jjW48ddY3C", "mkfr0AYygGjgm4MIC7KBc7qzFOtz9Nha")
	url := bcs.Sign("GET", "", "/", "", "", "")
	url_ex := "http://bcs.duapp.com//?sign=MBO:vYlphQiwbhVz67jjW48ddY3C:yf27Oy6JVtK6nxRtIASKX6H%2BR4I%3D"
	fmt.Println("test sign", url == url_ex)
	b, e := bcs.ListBuckets()

	fmt.Println(e)
	// for _, pBucket := range b {
	// 	fmt.Println(pBucket)
	// 	o, e := pBucket.ListObjects("", 0, 5)
	// 	fmt.Println(e)
	// 	for _, pObject := range o.Objects {
	// 		fmt.Println(pObject)
	// 	}
	// }

	testObj := b[0].Object("/test.txt")
	testObj, err := testObj.PutFile("test.txt")
	fmt.Println(testObj, err)
}
