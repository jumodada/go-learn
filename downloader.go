package main

import (
	"go-learn/infra"
)

func main() {
	retriever := infra.Retriever{}
	retriever.Get("https://baidu.com")
}
