package main

import (
	"fmt"
	"github.com/eggfly/baidu-bcs-sdk-go/bcs"
)

func main() {
	b := bcs.NewBCS("vYlphQiwbhVz67jjW48ddY3C", "mkfr0AYygGjgm4MIC7KBc7qzFOtz9Nha")
	url := b.Sign("GET", "", "/", "", "", "")
	url_ex := "http://bcs.duapp.com//?sign=MBO:vYlphQiwbhVz67jjW48ddY3C:yf27Oy6JVtK6nxRtIASKX6H%2BR4I%3D"
	fmt.Println("test sign", url == url_ex)
	b.ListBuckets()
}
