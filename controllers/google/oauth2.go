package controllersGoogle

import (
	"example-go-oauth-app/lib/google"

	"github.com/astaxie/beego"
)

type Oauth2Controller struct {
	beego.Controller
}

func (c *Oauth2Controller) Get() {
	config := google.GetConnect()

	url := config.AuthCodeURL("")

	c.Redirect(url, 302)
}
