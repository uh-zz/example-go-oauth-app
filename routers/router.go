package routers

import (
	"example-go-oauth-app/controllers"
	controllersGoogle "example-go-oauth-app/controllers/google"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/google/oauth2", &controllersGoogle.Oauth2Controller{})
	beego.Router("/google/callback", &controllersGoogle.CallbackController{})
	beego.Router("/", &controllers.MainController{})
}
