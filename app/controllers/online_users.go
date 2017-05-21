package controllers

import (
	"strconv"

	"github.com/three-plus-three/web_example/app/models"
	"github.com/three-plus-three/web_example/app/routes"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/forms"
	"github.com/three-plus-three/modules/toolbox"
)

func init() {
}

// OnlineUsers - 控制器
type OnlineUsers struct {
	App
}

// Index 列出所有记录
func (c OnlineUsers) Index() revel.Result {
	var cond orm.Cond
	var name string
	c.Params.Bind(&name, "query")
	if name != "" {
		cond = orm.Cond{"name LIKE": "%" + name + "%"}
	}

	total, err := c.Lifecycle.DB.OnlineUsers().Where().And(cond).Count()
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

	var onlineUsers []models.OnlineUser
	err = c.Lifecycle.DB.OnlineUsers().Where().
		And(cond).
		Offset(pageIndex * pageSize).
		Limit(pageSize).
		All(&onlineUsers)
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render()
	}

	if len(onlineUsers) > 0 {
		var authAccountIDList = make([]int64, 0, len(onlineUsers))
		for idx := range onlineUsers {
			authAccountIDList = append(authAccountIDList, onlineUsers[idx].AuthAccountID)
		}
		var authAccountByID map[int64]models.AuthAccount
		var authAccountList []models.AuthAccount
		err = c.Lifecycle.DB.AuthAccounts().Where().
			And(orm.Cond{"id IN": authAccountIDList}).
			All(&authAccountList)
		if err != nil {
			c.Validation.Error("load AuthAccount fail, " + err.Error())
		} else {
			if authAccountByID == nil {
				authAccountByID = make(map[int64]models.AuthAccount, len(authAccountList))
			}
			for idx := range authAccountList {
				authAccountByID[authAccountList[idx].ID] = authAccountList[idx]
			}
			c.ViewArgs["authAccountByID"] = authAccountByID
		}
	}

	paginator := toolbox.NewPaginator(c.Request.Request, pageSize, total)
	return c.Render(onlineUsers, paginator)
}

// New 编辑新建记录
func (c OnlineUsers) New() revel.Result {
	var err error

	var authAccounts []models.AuthAccount
	err = c.Lifecycle.DB.AuthAccounts().Where().
		All(&authAccounts)
	if err != nil {
		c.Validation.Error("load AuthAccount fail, " + err.Error())
		c.ViewArgs["authAccounts"] = []forms.InputChoice{{
			Value: "",
			Label: revel.Message(c.Request.Locale, "select.empty"),
		}}
	} else {
		var optAuthAccounts = make([]forms.InputChoice, 0, len(authAccounts))
		for _, o := range authAccounts {
			optAuthAccounts = append(optAuthAccounts, forms.InputChoice{
				Value: strconv.FormatInt(int64(o.ID), 10),
				Label: o.Name,
			})
		}
		c.ViewArgs["authAccounts"] = optAuthAccounts
	}
	return c.Render()
}

// Create 创建记录
func (c OnlineUsers) Create(onlineUser *models.OnlineUser) revel.Result {
	if onlineUser.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.New())
	}

	_, err := c.Lifecycle.DB.OnlineUsers().Insert(onlineUser)
	if err != nil {
		if oerr, ok := err.(*orm.Error); ok {
			for _, validation := range oerr.Validations {
				c.Validation.Error(validation.Message).Key(models.KeyForOnlineUsers(validation.Key))
			}
			c.Validation.Keep()
		}
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.New())
	}

	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.OnlineUsers.Index())
}

// Edit 编辑指定 id 的记录
func (c OnlineUsers) Edit(id int64) revel.Result {
	var onlineUser models.OnlineUser
	err := c.Lifecycle.DB.OnlineUsers().Id(id).Get(&onlineUser)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Index())
	}

	var authAccounts []models.AuthAccount
	err = c.Lifecycle.DB.AuthAccounts().Where().
		All(&authAccounts)
	if err != nil {
		c.Validation.Error("load AuthAccount fail, " + err.Error())
		c.ViewArgs["authAccounts"] = []forms.InputChoice{{
			Value: "",
			Label: revel.Message(c.Request.Locale, "select.empty"),
		}}
	} else {
		var optAuthAccounts = make([]forms.InputChoice, 0, len(authAccounts))
		for _, o := range authAccounts {
			optAuthAccounts = append(optAuthAccounts, forms.InputChoice{
				Value: strconv.FormatInt(int64(o.ID), 10),
				Label: o.Name,
			})
		}
		c.ViewArgs["authAccounts"] = optAuthAccounts
	}

	return c.Render(onlineUser)
}

// Update 按 id 更新记录
func (c OnlineUsers) Update(onlineUser *models.OnlineUser) revel.Result {
	if onlineUser.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Edit(int64(onlineUser.ID)))
	}

	err := c.Lifecycle.DB.OnlineUsers().Id(onlineUser.ID).Update(onlineUser)
	if err != nil {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "update.record_not_found"))
		} else {
			if oerr, ok := err.(*orm.Error); ok {
				for _, validation := range oerr.Validations {
					c.Validation.Error(validation.Message).Key(models.KeyForOnlineUsers(validation.Key))
				}
				c.Validation.Keep()
			}
			c.Flash.Error(err.Error())
		}
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Edit(int64(onlineUser.ID)))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	return c.Redirect(routes.OnlineUsers.Index())
}

// Delete 按 id 删除记录
func (c OnlineUsers) Delete(id int64) revel.Result {
	err := c.Lifecycle.DB.OnlineUsers().Id(id).Delete()
	if nil != err {
		if err == orm.ErrNotFound {
			c.Flash.Error(revel.Message(c.Request.Locale, "delete.record_not_found"))
		} else {
			c.Flash.Error(err.Error())
		}
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}
	return c.Redirect(routes.OnlineUsers.Index())
}

// DeleteByIDs 按 id 列表删除记录
func (c OnlineUsers) DeleteByIDs(id_list []int64) revel.Result {
	if len(id_list) == 0 {
		c.Flash.Error("请至少选择一条记录！")
		return c.Redirect(routes.OnlineUsers.Index())
	}
	_, err := c.Lifecycle.DB.OnlineUsers().Where().And(orm.Cond{"id IN": id_list}).Delete()
	if nil != err {
		c.Flash.Error(err.Error())
	} else {
		c.Flash.Success(revel.Message(c.Request.Locale, "delete.success"))
	}
	return c.Redirect(routes.OnlineUsers.Index())
}
