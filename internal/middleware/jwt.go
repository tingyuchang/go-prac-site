package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-prac-site/internal/models"
	"go-prac-site/internal/services"
	"net/http"
	"strings"
)

func JWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username := ps.ByName("username")
		authorizations := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorizations) != 2 {
			models.ResponseError(w, r, fmt.Errorf("authorization not found"))
			return
		}

		tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]

		err := services.Check(username, tokenStr)

		if err != nil {
			models.ResponseError(w, r, err)
			return
		}
		next(w,r,ps)
	}
}
