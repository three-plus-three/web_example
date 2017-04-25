package controllers

import (
	"github.com/three-plus-three/web_example/app/models"
	"github.com/three-plus-three/web_example/app/routes"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/toolbox"
)

// AuthAccounts - 控制器
type AuthAccounts struct {
	App
}

// Index 列出所有记录
func (c AuthAccounts) Index() revel.Result {
	var cond orm.Cond
	var name string
	c.Params.Bind(&name, "query")
	if name != "" {
		cond = orm.Cond{"name LIKE": "%" + name + "%"}
	}

	total, err := c.Lifecycle.DB.AuthAccounts().Where().And(cond).Count()
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render(err)
	}

	var pageIndex, pageSize int
	c.Params.Bind(&pageIndex, "pageIndex")
	c.Params.Bind(&pageSize, "pageSize")
	if pageSize <= 0 {
		pageSize = toolbox.DEFAULT_SIZE_PER_PAGE
	}

	var authAccounts []models.AuthAccount
	err = c.Lifecycle.DB.AuthAccounts().Where().
		And(cond).
		Offset(pageIndex * pageSize).
		Limit(pageSize).
		All(&authAccounts)
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render()
	}

	paginator := toolbox.NewPaginator(c.Request.Request, pageSize, total)
	return c.Render(authAccounts, paginator)
}

// New 编辑新建记录
func (c AuthAccounts) New() revel.Result {

	return c.Render()
}

// Create 创建记录
func (c AuthAccounts) Create(authAccount *models.AuthAccount) revel.Result {
	if authAccount.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.New())
	}

	_, err := c.Lifecycle.DB.AuthAccounts().Insert(authAccount)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(models.KeyForAuthAccounts(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.New())
	}

	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.AuthAccounts.Index())
}

// Edit 编辑指定 id 的记录
func (c AuthAccounts) Edit(id int64) revel.Result {
	var authAccount models.AuthAccount
	err := c.Lifecycle.DB.AuthAccounts().Id(id).Get(&authAccount)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Index())
	}

	return c.Render(authAccount)
}

// Update 按 id 更新记录
func (c AuthAccounts) Update(authAccount *models.AuthAccount) revel.Result {
	if authAccount.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Edit(int64(authAccount.ID)))
	}

	err := c.Lifecycle.DB.AuthAccounts().Id(authAccount.ID).Update(authAccount)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(models.KeyForAuthAccounts(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Edit(int64(authAccount.ID)))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	return c.Redirect(routes.AuthAccounts.Index())
}

// Delete 按 id 删除记录
func (c AuthAccounts) Delete(id int64) revel.Result {
	err := c.Lifecycle.DB.AuthAccounts().Id(id).Delete()
	if nil != err {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "delete.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}
	return c.Redirect(routes.AuthAccounts.Index())
}

// DeleteByIDs 按 id 列表删除记录
func (c AuthAccounts) DeleteByIDs(id_list []int64) revel.Result {
	if len(id_list) == 0 {
		c.Flash.Error("请至少选择一条记录！")
		return c.Redirect(routes.AuthAccounts.Index())
	}
	_, err := c.Lifecycle.DB.AuthAccounts().Where().And(orm.Cond{"id IN": id_list}).Delete()
	if nil != err {
		c.Flash.Error(err.Error())
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}
	return c.Redirect(routes.AuthAccounts.Index())
}
