package main

import (
	"net/http"

	"github.com/AlexanderThaller/httphelper"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

func pageAPIV0TransactionAdd(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}

func pageAPIV0TransactionGetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}

func pageAPIV0TransactionList(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}
