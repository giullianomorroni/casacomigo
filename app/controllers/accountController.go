package controllers

import (
	cb "casacomigo/app/controllers/base"
	"casacomigo/app/services/project"
	"github.com/robfig/revel"
    "fmt"
)

//** TYPES
type (
	Account struct {
		cb.BaseController
	}
)

//** INIT FUNCTION

// init is called when the first request into the controller is made
func init() {

}

//** CONTROLLER FUNCTIONS
func (this *Account) Register(url_name string) revel.Result {
	project, err := project.FindProject(url_name)
	fmt.Println(err)
	if err != nil {
		return this.Render(err.Error())
	}
	return this.Render(project);
}

func (this *Account) NewAccount(url_name string) revel.Result {
	return this.Render();
}
