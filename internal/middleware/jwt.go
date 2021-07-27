package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-prac-site/internal/router/response"
	"go-prac-site/internal/services"
	"net/http"
	"strings"
)

func JWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authorizations := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorizations) != 2 {
			response.ReturnError(w, r, fmt.Errorf("authorization not found"))
			return
		}

		tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]
		claim, err := services.Check(tokenStr)

		if err != nil {
			response.ReturnError(w, r, err)
			return
		}

		ps = append(ps, httprouter.Param{
			Key: "username",
			Value: claim.Username,
		})
		next(w,r,ps)
	}
}
