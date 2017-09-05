package routers

import (
	"github.com/astaxie/beego/context"
	"strings"
)

var filterLoggedInUser = func(ctx *context.Context) {
	url := ctx.Input.URL()
	if strings.HasPrefix(url, "/session") || strings.HasPrefix(url, "/graphql") || strings.HasPrefix(url, "/query") {
		return
	}
	_, ok := ctx.Input.Session("current_user").(map[string]string)
	if !ok {
		ctx.Redirect(302, "/session/login")
	}
}
