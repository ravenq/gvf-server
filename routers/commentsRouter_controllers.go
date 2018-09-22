package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:MyloveController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:PostController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UploadController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ravenq/gvf-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}