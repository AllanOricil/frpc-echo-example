package main

import (
	"context"

	"github.com/allanoricil/frpc-echo-example/echo"
)

type svc struct{}

func (s *svc) Echo(_ context.Context, req *echo.Request) (*echo.Response, error) {
	res := new(echo.Response)
	res.Message = req.Message
	return res, nil
}

func main() {
	frpcServer, err := echo.NewServer(new(svc), nil, nil)
	if err != nil {
		panic(err)
	}

	err = frpcServer.Start(":8080")
    if err != nil {
        panic(err)
    }
}