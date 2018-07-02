package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello world, go http server!!!!!")) // 브라우저의 응답
	}) // 경로 접속시 실행

	// http://localhost:9999 포트에서 실행
	http.ListenAndServe(":9999", nil)
}
