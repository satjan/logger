package main

import (
	"fmt"
	"github.com/satjan/logger"
)

func main() {
	logger.Init(fmt.Sprintf("%s/main.log", "."))
	logger.Error("error message")
	logger.Info("info message")
}
