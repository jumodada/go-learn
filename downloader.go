package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", bytes)
}
