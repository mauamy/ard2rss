package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func readParam(r *http.Request, name string) string {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(name)
}
