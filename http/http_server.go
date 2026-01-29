package http

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func MainServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Сервер стартовал на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}
