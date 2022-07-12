package main

import (
	"fmt"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	fmt.Println(1)
	logger.Warn("warning")

}
