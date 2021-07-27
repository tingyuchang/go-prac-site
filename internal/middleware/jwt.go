package middleware

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-prac-site/e"
	"go-prac-site/internal/models"
	"go-prac-site/internal/router/response"
	"go-prac-site/internal/services"
	"net/http"
	"strings"
)

func JWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authorizations := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorizations) != 2 {
			response.ReturnError(w, r, models.ResponseError{
				Err: fmt.Errorf("authorization not found"),
				Code: e.ERROR_AUTH_PARSEFORM,
				Desc: e.GetErrorMsg(e.ERROR_AUTH_PARSEFORM),

			})
			return
		}

		tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]
		claim, err := services.Check(tokenStr)

		if err != nil {
			response.ReturnError(w, r, models.ResponseError{
				Err: err,
				Code: e.ERROR_AUTH_INVALID_TOKEN,
				Desc: e.GetErrorMsg(e.ERROR_AUTH_INVALID_TOKEN),
			})
			return
		}

		ps = append(ps, httprouter.Param{
			Key: "username",
			Value: claim.Username,
		})
		next(w,r,ps)
	}
}
