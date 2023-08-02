package main

import (
	"log"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
)

// main PG_URL=1;RMQ_URL=1
func main() {
	// 加载配置
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// 启动服务
	app.Run(cfg)
}
