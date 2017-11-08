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

	revel.TemplateFuncs["sex_format"] = func(value string) string {
		switch value {
		case "male":
			return "男"
		case "female":
			return "女"
		default:
			return value
		}
	}
}

// AuthAccounts - 控制器
type AuthAccounts struct {
	App
}

// Index 列出所有记录
func (c AuthAccounts) Index() revel.Result {
	var page = c.PagingParams()

	var cond orm.Cond
	var query string
	c.Params.Bind(&query, "query")
	if query != "" {
		cond = orm.Cond{"name LIKE": "%" + query + "%"}
	}

	total, err := c.Lifecycle.DB.AuthAccounts().Where().And(cond).Count()
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render(err)
	}

	var authAccounts []models.AuthAccount
	err = c.Lifecycle.DB.AuthAccounts().Where().
		And(cond).
		Offset(page.Offset()).
		Limit(page.Limit()).
		All(&authAccounts)
	if err != nil {
		c.Validation.Error(err.Error())
		return c.Render()
	}

	var authaccountList = make([]map[string]interface{}, 0, len(authAccounts))
	for idx := range authAccounts {
		authaccountList = append(authaccountList, map[string]interface{}{
			"authaccount": authAccounts[idx],
		})
	}

	if len(authAccounts) > 0 {
		var authAccountIDList = make([]int64, 0, len(authAccounts))
		for idx := range authAccounts {
			authAccountIDList = append(authAccountIDList, authAccounts[idx].ManagerID)
		}
		var authAccountList []models.AuthAccount
		err = c.Lifecycle.DB.AuthAccounts().Where().
			And(orm.Cond{"id IN": authAccountIDList}).
			All(&authAccountList)
		if err != nil {
			c.Validation.Error("load AuthAccount fail, " + err.Error())
		} else {

			for idx := range authAccountList {
				for vidx := range authAccounts {
					if authAccountList[idx].ID == authAccounts[vidx].ManagerID {
						authaccountList[vidx]["authAccount"] = authAccountList[idx]
					}
				}
			}
		}
		authAccountIDList = authAccountIDList[:0]
		for idx := range authAccounts {
			authAccountIDList = append(authAccountIDList, authAccounts[idx].LeaderID)
		}
		authAccountList = nil
		err = c.Lifecycle.DB.AuthAccounts().Where().
			And(orm.Cond{"id IN": authAccountIDList}).
			All(&authAccountList)
		if err != nil {
			c.Validation.Error("load AuthAccount fail, " + err.Error())
		} else {

			for idx := range authAccountList {
				for vidx := range authAccounts {
					if authAccountList[idx].ID == authAccounts[vidx].LeaderID {
						authaccountList[vidx]["authAccount"] = authAccountList[idx]
					}
				}
			}
		}
	}

	paginator := page.Get(total)
	c.ViewArgs["authAccounts"] = authaccountList
	return c.Render(paginator)
}

func (c AuthAccounts) withAuthAccounts() ([]models.AuthAccount, error) {
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
	optAuthAccounts = append(optAuthAccounts, forms.InputChoice{
		Value: "",
		Label: revel.Message(c.Request.Locale, "select.empty"),
	})
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
func (c AuthAccounts) New() revel.Result {
	c.withAuthAccounts()
	c.withAuthAccounts()

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
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.New())
	}

	c.Flash.Success(revel.Message(c.Request.Locale, "insert.success"))
	return c.Redirect(routes.AuthAccounts.Index())
}

// Edit 编辑指定 id 的记录
func (c AuthAccounts) Edit(id int64) revel.Result {
	var authAccount models.AuthAccount
	err := c.Lifecycle.DB.AuthAccounts().ID(id).Get(&authAccount)
	if err != nil {
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Index())
	}

	c.withAuthAccounts()
	c.withAuthAccounts()
	return c.Render(authAccount)
}

// Update 按 id 更新记录
func (c AuthAccounts) Update(id int64, authAccount *models.AuthAccount) revel.Result {
	if authAccount.Validate(c.Validation) {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Edit(id))
	}

	err := c.Lifecycle.DB.AuthAccounts().ID(id).Update(authAccount)
	if err != nil {
		c.ErrorToFlash(err)
		c.FlashParams()
		return c.Redirect(routes.AuthAccounts.Edit(id))
	}
	c.Flash.Success(revel.Message(c.Request.Locale, "update.success"))
	return c.Redirect(routes.AuthAccounts.Index())
}

// Delete 按 id 删除记录
func (c AuthAccounts) Delete(id int64) revel.Result {
	err := c.Lifecycle.DB.AuthAccounts().ID(id).Delete()
	if nil != err {
		c.ErrorToFlash(err, "delete.record_not_found")
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
