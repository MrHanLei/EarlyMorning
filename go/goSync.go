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
	g, gctx := errgroup.WithContext(ctx) //初始化
	//var  g  errgroup.Group  //可以直接使用g，不用context

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	g.Go(func() error { //启动一个线程
		//阻塞启动http监听 当Shutdown方法执行后，会立即返回err
		return server.ListenAndServe()
	})

	g.Go(func() error { //启动一个线程
		s := <-signalChan
		fmt.Println("收到退出信号，优雅退出", s)
		err := server.Shutdown(gctx)
		if err != nil {
			return err
		}
		return nil
	})

	//阻塞 等待所有线程执行完，但是当有线程返回错误时会立即返回。不会停止已经在执行中的线程
	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait", err)
	} else {
		fmt.Println("successfully!")
	}

}
