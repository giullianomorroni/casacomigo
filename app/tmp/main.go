// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	_ "casacomigo/app"
	controllers1 "casacomigo/app/controllers"
	controllerBase "casacomigo/app/controllers/base"
	controllers0 "github.com/robfig/revel/modules/static/app/controllers"
	_ "github.com/robfig/revel/modules/testrunner/app"
	controllers "github.com/robfig/revel/modules/testrunner/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllerBase.BaseController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Before",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "After",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Panic",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"error",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Account)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Profile",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "url_name", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Shopping)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "FindCouple",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					18: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "ProductDetail",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					22: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Checkout",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					30: []string{ 
						"token",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Home)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					21: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Test)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "PaymentApproved",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "param", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					18: []string{ 
						"param",
					},
				},
			},
			&revel.MethodType{
				Name: "PaymentCanceled",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "param", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					22: []string{ 
						"param",
					},
				},
			},
			&revel.MethodType{
				Name: "PayPalAccountConfirm",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "PaymentReturn",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "param", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					33: []string{ 
						"param",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Signature)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "nomeNoivo", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "telefoneNoivo", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailNoivo", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "nomeNoiva", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "telefoneNoiva", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "emailNoiva", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "apelido", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "senha", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "dataCasamento", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "NewAccount",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					67: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "ChooseSite",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Payment",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					75: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RegisterPayment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "plano", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					89: []string{ 
						"token",
					},
				},
			},
			&revel.MethodType{
				Name: "RegisterSite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "site_id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"casacomigo/app/controllers.Signature.Register": { 
			25: "nomeNoivo",
			26: "nomeNoiva",
			28: "telefoneNoivo",
			29: "telefoneNoiva",
			31: "emailNoivo",
			32: "emailNoiva",
			34: "apelido",
			35: "dataCasamento",
			37: "senha",
		},
	}
	revel.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
