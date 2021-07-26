package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

type Adapter func(next http.Handler) http.Handler
type HttpRouterAdapter func(next httprouter.Handle)  httprouter.Handle

func Log() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before")
			defer fmt.Println("After")
			next.ServeHTTP(w, r)
		})
	}
}


func NotifyForHttpRouter() HttpRouterAdapter {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			fmt.Println("Before")
			defer fmt.Println("After")
			next(w,r,ps)
		}
	}
}

type Logger struct{

}

func NewAuth() *Logger {
	return &Logger{}
}

func (auth *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Before Auth")
	defer fmt.Println("After Auth")

	next(w,r)
}
