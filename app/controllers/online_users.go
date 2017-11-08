package controllers

import (
	"fmt"
	"strconv"

	"github.com/three-plus-three/web_example/app/models"
	"github.com/three-plus-three/web_example/app/routes"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/forms"
)

func init() {
}

// OnlineUsers - 控制器
type OnlineUsers struct {
	App
}

// Index 列出所有记录
func (c OnlineUsers) Index() revel.Result {
	var page = c.PagingParams()

	var cond orm.Cond
	var query string
	c.Params.Bind(&query, "query")
	if query != "" {
		cond = orm.Cond{"name LIKE": "%" + query + "%"}
	}

	total, err := c.Lifecycle.DB.OnlineUsers().Where().And(cond).Count()
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render(err)
	}

	var onlineUsers []models.OnlineUser
	err = c.Lifecycle.DB.OnlineUsers().Where().
		And(cond).
		Offset(page.Offset()).
		Limit(page.Limit()).
		All(&onlineUsers)
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render()
	}

	var onlineuserList = make([]map[string]interface{}, 0, len(onlineUsers))
	for idx := range onlineUsers {
		onlineuserList = append(onlineuserList, map[string]interface{}{
			"onlineuser": onlineUsers[idx],
		})
	}

	if len(onlineUsers) > 0 {
		var authAccountIDList = make([]int64, 0, len(onlineUsers))
		for idx := range onlineUsers {
			authAccountIDList = append(authAccountIDList, onlineUsers[idx].AuthAccountID)
		}
		var authAccountList []models.AuthAccount
		err = c.Lifecycle.DB.AuthAccounts().Where().
			And(orm.Cond{"id IN": authAccountIDList}).
			All(&authAccountList)
		if err != nil {
			c.Validation.Error("load AuthAccount fail, " + err.Error())
		} else {

			for idx := range authAccountList {
				for vidx := range onlineUsers {
					if authAccountList[idx].ID == onlineUsers[vidx].AuthAccountID {
						onlineuserList[vidx]["authAccount"] = authAccountList[idx]
					}
				}
			}
		}
	}

	paginator := page.Get(total)
	c.ViewArgs["onlineUsers"] = onlineuserList
	return c.Render(paginator)
}

func (c OnlineUsers) withAuthAccounts() ([]models.AuthAccount, error) {
	var authAccounts []models.AuthAccount
	err := c.Lifecycle.DB.AuthAccounts().Where().
		All(&authAccounts)
	if err != nil {
		c.Validation.Error("load AuthAccount fail, " + err.Error())
		c.ViewArgs["authAccounts"] = []forms.InputChoice{{
			Value: "",
			Label: revel.Message(c.Request.Locale, "select.empty"),
		}}
		return nil, err
	}

	var optAuthAccounts = make([]forms.InputChoice, 0, len(authAccounts))
	for _, o := range authAccounts {
		optAuthAccounts = append(optAuthAccounts, forms.InputChoice{
			Value: strconv.FormatInt(int64(o.ID), 10),
			Label: fmt.Sprint(o.Name),
		})
	}
	c.ViewArgs["authAccounts"] = optAuthAccounts
	return authAccounts, nil
}

// New 编辑新建记录
func (c OnlineUsers) New() revel.Result {
	c.withAuthAccounts()

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
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.New())
	}

	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.OnlineUsers.Index())
}

// Edit 编辑指定 id 的记录
func (c OnlineUsers) Edit(id int64) revel.Result {
	var onlineUser models.OnlineUser
	err := c.Lifecycle.DB.OnlineUsers().ID(id).Get(&onlineUser)
	if err != nil {
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Index())
	}

	c.withAuthAccounts()
	return c.Render(onlineUser)
}

// Update 按 id 更新记录
func (c OnlineUsers) Update(id int64, onlineUser *models.OnlineUser) revel.Result {
	if onlineUser.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Edit(id))
	}

	err := c.Lifecycle.DB.OnlineUsers().ID(id).Update(onlineUser)
	if err != nil {
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.OnlineUsers.Edit(id))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	return c.Redirect(routes.OnlineUsers.Index())
}

// Delete 按 id 删除记录
func (c OnlineUsers) Delete(id int64) revel.Result {
	err := c.Lifecycle.DB.OnlineUsers().ID(id).Delete()
	if nil != err {
		c.ErrorToFlash(err, "delete.record_not_found")
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
