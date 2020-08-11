package main

import (
	"context"
	"encoding/json"
	"github.com/egsam98/rusprofile/api/rusprofile"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	logger := initLogger()

	mux := runtime.NewServeMux(runtime.WithProtoErrorHandler(errorHandler))
	err := rusprofile.RegisterServiceHandlerServer(context.Background(), mux, rusprofile.NewRusprofileServer(logger))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "swagger/ui.html")
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("swagger"))))
	http.Handle("/api/", mux)

	if err := http.ListenAndServe(":8080", nil); err != nil {
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
	encoder := json.NewEncoder(w)
	err = encoder.Encode(rusprofile.Error{
		Error:    st.Message(),
		GRPCCode: uint32(st.Code()),
	})
	if err != nil {
		w.Write([]byte(`{"error": "failed to marshal error"}`))
	}
}
