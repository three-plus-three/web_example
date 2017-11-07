package controllers

import (
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/modules/web_ext"
	"github.com/three-plus-three/web_example/app"
	"github.com/three-plus-three/web_example/app/libs"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	Lifecycle *libs.Lifecycle
}

func (c *App) ErrorToFlash(err error) {
	if err == orm.ErrNotFound {
		c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
	} else {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).
					Key(validation.Key)
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
	}
}

func (c *App) CurrentUser() web_ext.User {
	return c.Lifecycle.CurrentUser(c.Controller)
}

func (c *App) init() revel.Result {
	c.Lifecycle = app.Lifecycle
	c.ViewArgs["menuList"] = c.Lifecycle.Menus()
	c.ViewArgs["controller"] = c.Name
	user := c.CurrentUser()
	if user != nil {
		c.ViewArgs["currentUsername"] = user.Name()
		c.ViewArgs["currentUser"] = user
	}
	return nil
}

func (c *App) IsAJAX() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// Pagination 分页参数
type Pagination struct {
	c           *revel.Controller
	Index, Size int
}

// Offset 偏移值
func (p Pagination) Offset() int {
	return p.Index * p.Size
}

// Limit 限制值
func (p Pagination) Limit() int {
	return p.Size
}

// Get 获取分页对象
func (p Pagination) Get(nums interface{}) *toolbox.Paginator {
	return toolbox.NewPaginator(p.c.Request.Request, p.Size, nums)
}

func (c *App) Pagination() Pagination {
	var pageIndex, pageSize int
	c.Params.Bind(&pageIndex, "pageIndex")
	c.Params.Bind(&pageSize, "pageSize")
	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	return Pagination{c: c.Controller, Index: pageIndex, Size: pageSize}
}

// func (c *ApplicationController) checkUser() revel.Result {
// 	return c.Lifecycle.CheckUser(c.Controller)
// }

func init() {
	revel.InterceptMethod((*App).init, revel.BEFORE)
	revel.InterceptFunc(func(c *revel.Controller) revel.Result {
		if check, ok := c.AppController.(interface {
			checkUser() revel.Result
		}); ok {
			return check.checkUser()
		}
		return nil
	}, revel.BEFORE, revel.AllControllers)
}
