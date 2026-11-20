package main

import "github.com/khoand3012/bookmark-management/internal/api"

// @title           Bookmark Management API
// @version         1.0.0
// @description     Swagger doc for Bookmark Management API.
// @BasePath        /
func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
