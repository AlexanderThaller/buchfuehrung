package main

import (
	"net/http"
	"strconv"
	"time"

	data "github.com/AlexanderThaller/buchfuehrung"
	"github.com/AlexanderThaller/httphelper"
	"github.com/jinzhu/now"
	"github.com/juju/errgo"
	"github.com/julienschmidt/httprouter"
)

func pageAPIV0AccountAdd(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	l := httphelper.NewHandlerLogEntry(r)

	name := r.PostFormValue("name")
	comment := r.PostFormValue("comment")
	accounttype := r.PostFormValue("type")

	account := data.NewAccount(name, comment, accounttype)

	l.Debug("New Account: ", account)

	query := Database.Create(&account)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not insert account into database"))
	}

	return httphelper.MarshalCompactJsonToWriter(w, account)
}

func pageAPIV0AccountGetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not convert id parameter to integer"))
	}

	var account data.Account
	query := Database.Find(&account, id)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not find account with this id"))
	}

	return httphelper.MarshalCompactJsonToWriter(w, account)
}

func pageAPIV0AccountGetByName(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	l := httphelper.NewHandlerLogEntry(r)
	l.Debug("Name: ", p.ByName("name"))

	var account data.Account

	query := Database.Where(&data.Account{Name: p.ByName("name")}).First(&account)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not find account with this name"))
	}

	return httphelper.MarshalCompactJsonToWriter(w, account)
}

func pageAPIV0AccountList(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	var accounts []data.Account
	Database.Find(&accounts)

	return httphelper.MarshalCompactJsonToWriter(w, accounts)
}

func pageAPIV0TransactionAdd(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	l := httphelper.NewHandlerLogEntry(r)

	comment := r.PostFormValue("comment")
	l.Debug("Comment: ", comment)

	account := data.Account{Name: r.PostFormValue("account")}
	l.Debug("Account: ", account)

	category := data.Category{Name: r.PostFormValue("category")}
	l.Debug("Category: ", category)

	payee := data.Payee{Name: r.PostFormValue("payee")}
	l.Debug("Payee: ", payee)

	inflow, err := strconv.ParseFloat(r.PostFormValue("inflow"), 10)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not parse inflow from parameter"))
	}
	l.Debug("Inflow: ", inflow)

	outflow, err := strconv.ParseFloat(r.PostFormValue("outflow"), 10)
	if err != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not parse outflow from parameter"))
	}
	l.Debug("Outflow: ", outflow)

	timestamp := time.Now()
	if r.PostFormValue("timestamp") != "" {
		timestamp, err = now.Parse(r.PostFormValue("timestamp"))
		if err != nil {
			return httphelper.NewHandlerErrorDef(errgo.Notef(err, "can not parse timestamp from parameter"))
		}
	}
	l.Debug("TimeStamp: ", timestamp)

	query := Database.Create(&account)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not insert account into database"))
	}

	transaction := data.Transaction{
		Comment:   comment,
		Category:  category,
		Payee:     payee,
		Inflow:    inflow,
		Outflow:   outflow,
		TimeStamp: timestamp,
	}

	query = Database.FirstOrCreate(&account, account)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not get or insert account from/into database"))
	}

	account.Transactions = append(account.Transactions, transaction)

	query = Database.Save(&account)
	if query.Error != nil {
		return httphelper.NewHandlerErrorDef(errgo.Notef(query.Error, "can not insert transaction into database"))
	}

	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}

func pageAPIV0TransactionGetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}

func pageAPIV0TransactionList(w http.ResponseWriter, r *http.Request, p httprouter.Params) *httphelper.HandlerError {
	return httphelper.NewHandlerErrorDef(errgo.New("not implemented"))
}
