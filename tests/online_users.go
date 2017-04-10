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

	t.AssertContains(fmt.Sprint(onlineUser.AuthAccountID))
	t.AssertContains(fmt.Sprint(onlineUser.Hostaddress))
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

	v.Set("onlineUser.AuthAccountID", "abc")

	v.Set("onlineUser.Hostaddress", "Est aut consequuntur aut quo voluptates modi voluptatem.")

	v.Set("onlineUser.Macaddress", "Totam sapiente et porro in debitis ullam harum.")

	v.Set("onlineUser.CreatedAt", "1995-06-16T06:48:20+08:00")

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

	t.AssertEqual(fmt.Sprint(onlineUser.AuthAccountID), v.Get("onlineUser.AuthAccountID"))
	t.AssertEqual(fmt.Sprint(onlineUser.Hostaddress), v.Get("onlineUser.Hostaddress"))
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

	t.AssertContains(fmt.Sprint(onlineUser.AuthAccountID))
	t.AssertContains(fmt.Sprint(onlineUser.Hostaddress))
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

	v.Set("onlineUser.AuthAccountID", "abc")

	v.Set("onlineUser.Hostaddress", "Id accusantium veniam velit culpa dicta consectetur dolore.")

	v.Set("onlineUser.Macaddress", "Dolorum qui omnis dolorum dolores commodi aliquid.")

	v.Set("onlineUser.CreatedAt", "1998-11-08T03:04:55+08:00")

	t.Post(t.ReverseUrl("OnlineUsers.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var onlineUser models.OnlineUser
	err := app.Lifecycle.DB.OnlineUsers().Id(ruleId).Get(&onlineUser)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(onlineUser.AuthAccountID), v.Get("onlineUser.AuthAccountID"))

	t.AssertEqual(fmt.Sprint(onlineUser.Hostaddress), v.Get("onlineUser.Hostaddress"))

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
