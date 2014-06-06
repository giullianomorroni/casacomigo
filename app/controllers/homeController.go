package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
)

type (
	Home struct {
		cb.BaseController
	}
)

func init() {
	revel.InterceptMethod((*Home).Before, revel.BEFORE)
	revel.InterceptMethod((*Home).After, revel.AFTER)
	revel.InterceptMethod((*Home).Panic, revel.PANIC)
}

func (this *Home) Index() revel.Result {
	return this.Render()
}
