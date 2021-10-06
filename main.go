package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// flag.Set("v", "4")
	// glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler) // 注册function
	http.ListenAndServe(":8000", nil)
	// mux := http.NewServeMux()
	http.HandleFunc("/healthz", healthz)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, fmt.Sprintf("====================================question 1========================================\n"))
	io.WriteString(w, fmt.Sprintf("==================打印header, 并将request中的header写入response header:=================\n"))
	header := w.Header()
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
			header[k] = v
		}
	}

	io.WriteString(w, "============================== 现在的Response header:========================\n")
	if len(header) > 0 {
		for i, j := range header {
			io.WriteString(w, fmt.Sprintf("%s=%s\n", i, j))
		}
	}

	io.WriteString(w, fmt.Sprintf("====================================question 2========================================\n"))
	io.WriteString(w, fmt.Sprintf("===========================读取当前环境变量VERSION设置，并写入response header:=======================\n"))
	VERSION := os.Getenv("VERSION")
	GOPATH := os.Getenv("GOPATH")

	header["GOPATH"] = []string{GOPATH}
	io.WriteString(w, fmt.Sprintf("GOPATH=%s\n", header["GOPATH"]))

	header["VERSION"] = []string{VERSION}
	io.WriteString(w, fmt.Sprintf("VERSION=%s\n", header["VERSION"]))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("200"))
}
