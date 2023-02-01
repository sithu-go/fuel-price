package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"fuel-price/cmd/handler"
	_ "fuel-price/conf"
	"fuel-price/pkg/ds"
)

func main() {

	// to get file line and path when print
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := flag.String("port", "8080", "default port is 8080")
	flag.Parse()
	addr := fmt.Sprintf(":%s", *port)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	ds, _ := ds.NewDataSource()

	h := handler.NewHandler(
		&handler.HConfig{
			R:  router,
			DS: ds,
		},
	)
	h.Register()

	server := http.Server{
		Addr:           addr,
		Handler:        h.R,
		ReadTimeout:    time.Duration(time.Minute * 3),
		WriteTimeout:   time.Duration(time.Minute * 3),
		MaxHeaderBytes: 10 << 20, //10MB
	}

	go func() {
		fmt.Println("server started listening on port : ", *port)
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("server failed to initialized  on port : ", *port)
			log.Fatalf("error on listening :%v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	// shutdown close
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("Failed to shutdown server: ", err.Error())
	}
}
