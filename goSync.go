package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	g, gctx := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		s := <-signalChan
		fmt.Println("收到退出信号", s)
		err := server.Shutdown(gctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait", err)
	} else {
		fmt.Println("successfully!")
	}

}
