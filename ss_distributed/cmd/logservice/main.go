package main

import (
	"context"
	"fmt"
	stlog "log"
	"ss_distributed/log"
	"ss_distributed/registry"
	"ss_distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "log Service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down server")

}
