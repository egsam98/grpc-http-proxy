package main

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"os"
	rusprofile2 "rusproile/api/rusprofile"
)

var PORT = os.Getenv("PORT")

func main() {
	logger := initLogger()

	mux := runtime.NewServeMux(runtime.WithProtoErrorHandler(errorHandler))
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

func errorHandler(_ context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	st := status.Convert(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(runtime.HTTPStatusFromCode(st.Code()))
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"error":     st.Message(),
		"gRPC-code": st.Code(),
	})
	if err != nil {
		w.Write([]byte(`{"error": "failed to marshal error"}`))
	}
}
