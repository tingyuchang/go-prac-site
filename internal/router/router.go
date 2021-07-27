package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"go-prac-site/internal/middleware"
	"go-prac-site/internal/router/api"
	"go-prac-site/internal/router/response"
	"net/http"
)

var router *httprouter.Router

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.BasicAuth()
	response.DefaultResponse(w, map[string]interface{}{
		"msg": "success",
	})
}

func NewRouter() *httprouter.Router {
	router = httprouter.New()
	router.GET("/", middleware.NotifyForHttpRouter()(Index))
	router.POST("/user", api.CreateUser)
	router.POST("/login", api.Auth)
	router.GET("/user/:id", middleware.JWT(api.GetUserInfomation))
	return router
}

func NewNegroni() *negroni.Negroni {
	router = httprouter.New()
	router.GET("/", Index)
	router.POST("/user", api.CreateUser)
	router.POST("/login", api.Auth)
	router.GET("/user/:uid", middleware.JWT(api.GetUserInfomation))
	n := negroni.Classic()
	n.UseHandler(router)

	return n
}





