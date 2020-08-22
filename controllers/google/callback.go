package controllersGoogle

import (
	"errors"
	"example-go-oauth-app/lib/google"
	"fmt"

	"golang.org/x/net/context"

	"github.com/astaxie/beego"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackController struct {
	beego.Controller
}

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func (c *CallbackController) Get() {
	req := CallbackRequest{}

	if err := c.ParseForm(&req); err != nil {
		panic(err)
	}

	config := google.GetConnect()

	context := context.Background()

	token, err := config.Exchange(context, req.Code)
	if err != nil {
		panic(err)
	}

	if !token.Valid() {
		panic(errors.New("valid token"))
	}

	service, _ := v2.New(config.Client(context, token))
	tokenInfo, _ := service.Tokeninfo().AccessToken(token.AccessToken).Context(context).Do()

	c.Data["ID"] = tokenInfo.UserId
	c.Data["Email"] = tokenInfo.Email
	fmt.Println(c.Data)
	c.TplName = "google/callback.tpl"
}
