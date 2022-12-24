package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From 127.17.0.4."))
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(":20001", nil)
	}()
	err := <-errChan
	if err != nil {
		fmt.Println("Hello server stop running.")
	}

}
