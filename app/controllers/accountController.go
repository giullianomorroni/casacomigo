package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/account"
	"casacomigo/app/models/account"
	"github.com/robfig/revel"
)

type (
	Account struct {
		cb.BaseController
	}
)

func init() {}

func (this *Account) Site() revel.Result {
	return this.Render();
}
