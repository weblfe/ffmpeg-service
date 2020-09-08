package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tal-tech/go-zero/core/logx"
)

func init() {
	logx.SetLevel(logx.ErrorLevel)
}
