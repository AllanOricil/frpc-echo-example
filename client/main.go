package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/allanoricil/frpc-echo-example/echo"
)

func main() {
	c, err := echo.NewClient(nil, nil)
	if err != nil {
		panic(err)
	}

	err = c.Connect("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	req := echo.NewRequest()
	i := 0
	for {
		select {
		case <-stop:
			err = c.Close()
			if err != nil {
				panic(err)
			}
			return
		default:
			req.Message = fmt.Sprintf("#%d", i)
			log.Printf("Sending Request %s\n", req.Message)
			res, err := c.EchoService.Echo(context.Background(), req)
			if err != nil {
				panic(err)
			}
			log.Printf("Received Response %s\n", res.Message)
			time.Sleep(time.Second)
			i++
		}
	}
}
