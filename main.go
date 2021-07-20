package main

import (
	"20210703/internal/router"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", router.NewRouter())
	if err != nil {
		panic(err)
	}
}
