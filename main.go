package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
	"os"
	rusprofile2 "rusproile/api/rusprofile"
)

var PORT = os.Getenv("PORT")

func main() {
	logger := initLogger()

	mux := runtime.NewServeMux()
	err := rusprofile2.RegisterServiceHandlerServer(context.Background(), mux, rusprofile2.NewRusprofileServer(logger))
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":"+PORT, mux); err != nil {
		panic(err)
	}
}

func initLogger() grpclog.LoggerV2 {
	f, err := os.OpenFile("logs/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, f)
}
