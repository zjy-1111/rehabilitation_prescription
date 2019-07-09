package main

import (
	"fmt"
	"net/http"
	"rehabilitation-training/pkg/setting"
	"rehabilitation-training/routers"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
