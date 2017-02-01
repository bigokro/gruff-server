package main

import (
	"fmt"
	"github.com/bigokro/gruff-server/api"
	"github.com/bigokro/gruff-server/config"
	"github.com/labstack/echo/engine/fasthttp"
	"os"
	"time"
)

func main() {

	config.Init()
	api.RW_DB_POOL = config.InitDB()
	api.RW_DB_POOL.LogMode(true)
	api.RW_DB_POOL.DB().SetMaxIdleConns(100)
	api.RW_DB_POOL.DB().SetMaxIdleConns(1000)

	root := api.SetUpRouter(false, api.RW_DB_POOL)

	fmt.Printf("Starting %s server on port %s at %s\n", os.Getenv("GRUFF_NAME"), os.Getenv("GRUFF_PORT"), time.Now().String())
	root.Run(fasthttp.New(":" + os.Getenv("GRUFF_PORT")))
}
