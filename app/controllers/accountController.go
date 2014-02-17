package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
)

type (
	Account struct {
		cb.BaseController
	}
)

func init() {}

func (this *Account) Login() revel.Result {
	return this.Render();
}

func (this *Account) Profile(url_name string) revel.Result {
	return this.Render();
}

