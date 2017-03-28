package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/three-plus-three/web_example/app"
	"github.com/three-plus-three/web_example/app/models"
)

//  OnlineUsersTest 测试
type OnlineUsersTest struct {
	BaseTest
}

func (t OnlineUsersTest) TestIndex() {
	t.ClearTable("tpt_online_users")
	t.LoadFiles("tests/fixtures/online_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)

	t.Get(t.ReverseUrl("OnlineUsers.Index"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	//t.AssertContains("这是一个规则名,请替换成正确的值")

	var onlineUser models.OnlineUser
	err := app.Lifecycle.DB.OnlineUsers().Id(ruleId).Get(&onlineUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertContains(fmt.Sprint(onlineUser.Name))
	t.AssertContains(fmt.Sprint(onlineUser.AuthAccountID))
	t.AssertContains(fmt.Sprint(onlineUser.Ipaddress))
	t.AssertContains(fmt.Sprint(onlineUser.Macaddress))
}

func (t OnlineUsersTest) TestNew() {
	t.ClearTable("tpt_online_users")
	t.Get(t.ReverseUrl("OnlineUsers.New"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t OnlineUsersTest) TestCreate() {
	t.ClearTable("tpt_online_users")
	v := url.Values{}

	v.Set("onlineUser.Name", "37i")

	v.Set("onlineUser.AuthAccountID", "abc")

	v.Set("onlineUser.Ipaddress", "Rerum nihil architecto officiis placeat voluptate.")

	v.Set("onlineUser.Macaddress", "Hic consequatur officiis.")

	v.Set("onlineUser.CreatedAt", "2002-01-19T22:16:25+08:00")

	t.Post(t.ReverseUrl("OnlineUsers.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)

	var onlineUser models.OnlineUser
	err := app.Lifecycle.DB.OnlineUsers().Id(ruleId).Get(&onlineUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(onlineUser.Name), v.Get("onlineUser.Name"))
	t.AssertEqual(fmt.Sprint(onlineUser.AuthAccountID), v.Get("onlineUser.AuthAccountID"))
	t.AssertEqual(fmt.Sprint(onlineUser.Ipaddress), v.Get("onlineUser.Ipaddress"))
	t.AssertEqual(fmt.Sprint(onlineUser.Macaddress), v.Get("onlineUser.Macaddress"))
}

func (t OnlineUsersTest) TestEdit() {
	t.ClearTable("tpt_online_users")
	t.LoadFiles("tests/fixtures/online_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)
	t.Get(t.ReverseUrl("OnlineUsers.Edit", ruleId))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	var onlineUser models.OnlineUser
	err := app.Lifecycle.DB.OnlineUsers().Id(ruleId).Get(&onlineUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}
	fmt.Println(string(t.ResponseBody))

	t.AssertContains(fmt.Sprint(onlineUser.Name))
	t.AssertContains(fmt.Sprint(onlineUser.AuthAccountID))
	t.AssertContains(fmt.Sprint(onlineUser.Ipaddress))
	t.AssertContains(fmt.Sprint(onlineUser.Macaddress))
}

func (t OnlineUsersTest) TestUpdate() {
	t.ClearTable("tpt_online_users")
	t.LoadFiles("tests/fixtures/online_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)
	v := url.Values{}
	v.Set("_method", "PUT")
	v.Set("onlineUser.ID", strconv.FormatInt(ruleId, 10))

	v.Set("onlineUser.Name", "65c")

	v.Set("onlineUser.AuthAccountID", "abc")

	v.Set("onlineUser.Ipaddress", "Eos quo voluptate.")

	v.Set("onlineUser.Macaddress", "In at labore fuga quos.")

	v.Set("onlineUser.CreatedAt", "1985-07-20T20:19:27+08:00")

	t.Post(t.ReverseUrl("OnlineUsers.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var onlineUser models.OnlineUser
	err := app.Lifecycle.DB.OnlineUsers().Id(ruleId).Get(&onlineUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(onlineUser.Name), v.Get("onlineUser.Name"))

	t.AssertEqual(fmt.Sprint(onlineUser.AuthAccountID), v.Get("onlineUser.AuthAccountID"))

	t.AssertEqual(fmt.Sprint(onlineUser.Ipaddress), v.Get("onlineUser.Ipaddress"))

	t.AssertEqual(fmt.Sprint(onlineUser.Macaddress), v.Get("onlineUser.Macaddress"))

}

func (t OnlineUsersTest) TestDelete() {
	t.ClearTable("tpt_online_users")
	t.LoadFiles("tests/fixtures/online_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)
	t.Delete(t.ReverseUrl("OnlineUsers.Delete", ruleId))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("tpt_online_users", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}

func (t OnlineUsersTest) TestDeleteByIDs() {
	t.ClearTable("tpt_online_users")
	t.LoadFiles("tests/fixtures/online_users.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_online_users", conds)
	t.Delete(t.ReverseUrl("OnlineUsers.DeleteByIDs", []interface{}{ruleId}))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("tpt_online_users", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}
