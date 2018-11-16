package main

import (
	"os"

	flag "github.com/spf13/pflag"

	"cloudgo/service"
)

const (
	// PORT 配置端口号
	PORT string = "8080"
)

func main() {
	// 从环境变量中获取端口号
	port := os.Getenv("PORT")
	// 端口号不存在，将端口号设置为预设的端口号
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
