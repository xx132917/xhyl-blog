package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:ArticleTypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:FileController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:FileController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/download/:filename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:FileController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:FileController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:MessageController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Email",
            Router: `/email`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xhylblog/controllers:UserController"] = append(beego.GlobalControllerRouter["xhylblog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
