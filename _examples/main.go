package main

import (
	"fmt"
	"github.com/satjan/logger"
)

func main() {
	logger.Init(fmt.Sprintf("%s/main.log", "."))
	logger.Info("info message")
	logger.Error("error message")
}
