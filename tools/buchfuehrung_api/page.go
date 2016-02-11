package main

import (
	"net/http"

	"github.com/AlexanderThaller/httphelper"
	"github.com/julienschmidt/httprouter"
)

func pageRoot(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	http.Redirect(w, r, "/gallery", http.StatusMovedPermanently)
	return nil
}
