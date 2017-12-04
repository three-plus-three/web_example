package controllers

import (
	"strconv"

	"github.com/runner-mei/orm"
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

func (c *App) ErrorToFlash(err error, notFoundKey ...string) {
	if err == orm.ErrNotFound {
		if len(notFoundKey) >= 1 && notFoundKey[0] != "" {
			c.Flash.Error(revel.Message(c.Request.Locale, notFoundKey[0]))
		} else {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		}
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

// PagingParams 分页参数
type PagingParams struct {
	c           *revel.Controller
	Index, Size int
}

// Offset 偏移值
func (p PagingParams) Offset() int {
	return p.Index * p.Size
}

// Limit 限制值
func (p PagingParams) Limit() int {
	return p.Size
}

// Get 获取分页对象
func (p PagingParams) Get(total interface{}) *toolbox.Paginator {
	form, _ := p.c.Request.GetForm()
	if form != nil {
		pageIndex, _ := strconv.Atoi(form.Get("pageIndex"))
		return toolbox.NewPaginatorWith(p.c.Request.URL, pageIndex, p.Size, total)
	}
	return toolbox.NewPaginatorWith(p.c.Request.URL, 0, p.Size, total)
}

func (c *App) pagingParams() PagingParams {
	var pageIndex, pageSize int
	c.Params.Bind(&pageIndex, "pageIndex")
	c.Params.Bind(&pageSize, "pageSize")
	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	return PagingParams{c: c.Controller, Index: pageIndex, Size: pageSize}
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
