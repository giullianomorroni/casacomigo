package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
)

type (
	Shopping struct {
		cb.BaseController
	}
)

func init() {}

func (this *Shopping) FindCouple() revel.Result {
	return this.Render();
}
