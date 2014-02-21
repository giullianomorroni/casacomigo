package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/product"
	"github.com/robfig/revel"
	//"fmt"
)

type (
	Product struct {
		cb.BaseController
	}
)

func init() {}

func (this *Product) List(limite, pular int) revel.Result {	
	list := srv.List(limite, pular);
	return this.RenderJson(list);
}

