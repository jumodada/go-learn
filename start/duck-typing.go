package main

import "fmt"

type Retriever interface {
	Get(url string) string
	Set(url string)
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

type Poster interface {
	Post(url string, form map[string]string)
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{
		"name":  "1",
		"child": "mike",
	})
}

func session() {

}

func main() {
	var r Retriever
	fmt.Println(download(r))
}
