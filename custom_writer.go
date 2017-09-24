package main

import (
	"io"
	"net/http"
	"fmt"
	"os"
)

type logWriter struct {

}

func main() {
	resp, err := http.Get("http://vetkom.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1);
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte)(int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
