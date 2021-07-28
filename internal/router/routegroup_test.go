package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteGroupOfARouteGroup(t *testing.T) {
	var get bool
	router := httprouter.New()
	foo := NewRouteGroup(router, "/foo")
	bar := foo.NewGroup("/bar")

	bar.GET("/GET", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		get = true
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar/GET", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, get, true)
}

func TestRouteGroupAPI(t *testing.T) {
	var get, head, options, post, put, patch, delete bool

	//httpHandler := handlerStruct{&handler}
	router := httprouter.New()
	group := NewRouteGroup(router, "/foo")

	group.GET("/GET", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		get = true
	})
	group.HEAD("/GET", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		head = true
	})
	group.OPTIONS("/GET", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		options = true
	})
	group.POST("/POST", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		post = true
	})
	group.PUT("/PUT", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		put = true
	})
	group.PATCH("/PATCH", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		patch = true
	})
	group.DELETE("/DELETE", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		delete = true
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/GET", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, get, true)

	r, _ = http.NewRequest("HEAD", "/foo/GET", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, head, true)

	r, _ = http.NewRequest("OPTIONS", "/foo/GET", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, options, true)

	r, _ = http.NewRequest("POST", "/foo/POST", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, post, true)

	r, _ = http.NewRequest("PUT", "/foo/PUT", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, put, true)

	r, _ = http.NewRequest("PATCH", "/foo/PATCH", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, patch, true)

	r, _ = http.NewRequest("DELETE", "/foo/DELETE", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, delete, true)

}
