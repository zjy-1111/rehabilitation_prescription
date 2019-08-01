package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rehabilitation_prescription/models"
	"rehabilitation_prescription/pkg/ali_oss"
	"rehabilitation_prescription/pkg/gredis"
	"rehabilitation_prescription/pkg/logging"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/routers"
	"rehabilitation_prescription/util"
	"time"
)

func init() {
	setting.InitConf()
	models.InitDB()
	logging.InitLogger()
	util.InitUtil()
	gredis.InitRedis()
	ali_oss.InitOssBucket()
}

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 收到中断信号不退出
	quit := make(chan os.Signal)      // 信号量通道
	signal.Notify(quit, os.Interrupt) // 捕获信号
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}

	log.Println("Server exiting")
}
