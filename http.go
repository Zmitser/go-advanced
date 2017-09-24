package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("http://vetkom.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1);
	}
	//bs := make([]byte, 9999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}
