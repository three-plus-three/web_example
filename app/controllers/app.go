package controllers

import (
	"github.com/three-plus-three/modules/web_ext"
	"github.com/three-plus-three/web_example/app"
	"github.com/three-plus-three/web_example/app/libs"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	Lifecycle *libs.Lifecycle
}

func (c *App) CurrentUser() web_ext.User {
	return c.Lifecycle.CurrentUser(c.Controller)
}

func (c *App) initLifecycle() revel.Result {
	c.Lifecycle = app.Lifecycle
	c.ViewArgs["menuList"] = c.Lifecycle.MenuList
	user := c.CurrentUser()
	if user != nil {
		c.ViewArgs["currentUsername"] = user.Name()
		c.ViewArgs["currentUser"] = user
	}
	return nil
}

func (c App) beforeInvoke() revel.Result {
	c.ViewArgs["controller"] = c.Name
	return nil
}

func (c *App) IsAjax() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// func (c *ApplicationController) checkUser() revel.Result {
// 	return c.Lifecycle.CheckUser(c.Controller)
// }

func init() {
	revel.InterceptMethod((*App).initLifecycle, revel.BEFORE)
	revel.InterceptMethod(func(c interface{}) revel.Result {
		if check, ok := c.(interface {
			CheckUser() revel.Result
		}); ok {
			return check.CheckUser()
		}
		return nil
	}, revel.BEFORE)

	revel.InterceptMethod((App).beforeInvoke, revel.BEFORE)
}
