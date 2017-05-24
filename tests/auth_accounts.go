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

//  AuthAccountsTest 测试
type AuthAccountsTest struct {
	BaseTest
}

func (t AuthAccountsTest) TestIndex() {
	t.ClearTable("tpt_auth_accounts")
	t.LoadFiles("tests/fixtures/auth_accounts.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)

	t.Get(t.ReverseUrl("AuthAccounts.Index"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
	//t.AssertContains("这是一个规则名,请替换成正确的值")

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().Id(ruleId).Get(&authAccount)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertContains(fmt.Sprint(authAccount.ManagerID))
	t.AssertContains(fmt.Sprint(authAccount.LeaderID))
	t.AssertContains(fmt.Sprint(authAccount.Name))
	t.AssertContains(fmt.Sprint(authAccount.Email))
	t.AssertContains(fmt.Sprint(authAccount.Sex))
	t.AssertContains(fmt.Sprint(authAccount.Level))
	t.AssertContains(fmt.Sprint(authAccount.Description))
	t.AssertContains(fmt.Sprint(authAccount.Birthday))
}

func (t AuthAccountsTest) TestNew() {
	t.ClearTable("tpt_auth_accounts")
	t.Get(t.ReverseUrl("AuthAccounts.New"))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t AuthAccountsTest) TestCreate() {
	t.ClearTable("tpt_auth_accounts")
	v := url.Values{}

	v.Set("authAccount.ManagerID", "7639537466609394841")

	v.Set("authAccount.LeaderID", "5628780035462113946")

	v.Set("authAccount.Name", "t46")

	v.Set("authAccount.Password", "m8u9m38si")

	v.Set("authAccount.Email", "Rerum ullam rerum facere ut alias nobis est.")

	v.Set("authAccount.Sex", "Voluptates perferendis corrupti sapiente aut quae harum.")

	v.Set("authAccount.Level", "Eos placeat nisi cum neque temporibus voluptatibus neque.")

	v.Set("authAccount.Description", "Quasi sed nisi aut saepe eos est impedit.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "1981-02-14T23:31:34+08:00")

	v.Set("authAccount.UpdatedAt", "1991-01-21T14:17:23+08:00")

	t.Post(t.ReverseUrl("AuthAccounts.Create"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().Id(ruleId).Get(&authAccount)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(authAccount.ManagerID), v.Get("authAccount.ManagerID"))
	t.AssertEqual(fmt.Sprint(authAccount.LeaderID), v.Get("authAccount.LeaderID"))
	t.AssertEqual(fmt.Sprint(authAccount.Name), v.Get("authAccount.Name"))
	t.AssertEqual(fmt.Sprint(authAccount.Password), v.Get("authAccount.Password"))
	t.AssertEqual(fmt.Sprint(authAccount.Email), v.Get("authAccount.Email"))
	t.AssertEqual(fmt.Sprint(authAccount.Sex), v.Get("authAccount.Sex"))
	t.AssertEqual(fmt.Sprint(authAccount.Level), v.Get("authAccount.Level"))
	t.AssertEqual(fmt.Sprint(authAccount.Description), v.Get("authAccount.Description"))
	t.AssertEqual(fmt.Sprint(authAccount.Birthday), v.Get("authAccount.Birthday"))
}

func (t AuthAccountsTest) TestEdit() {
	t.ClearTable("tpt_auth_accounts")
	t.LoadFiles("tests/fixtures/auth_accounts.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)
	t.Get(t.ReverseUrl("AuthAccounts.Edit", ruleId))
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().Id(ruleId).Get(&authAccount)
	if err != nil {
		t.Assertf(false, err.Error())
	}
	fmt.Println(string(t.ResponseBody))

	t.AssertContains(fmt.Sprint(authAccount.ManagerID))
	t.AssertContains(fmt.Sprint(authAccount.LeaderID))
	t.AssertContains(fmt.Sprint(authAccount.Name))
	t.AssertContains(fmt.Sprint(authAccount.Email))
	t.AssertContains(fmt.Sprint(authAccount.Sex))
	t.AssertContains(fmt.Sprint(authAccount.Level))
	t.AssertContains(fmt.Sprint(authAccount.Description))
	t.AssertContains(fmt.Sprint(authAccount.Birthday))
}

func (t AuthAccountsTest) TestUpdate() {
	t.ClearTable("tpt_auth_accounts")
	t.LoadFiles("tests/fixtures/auth_accounts.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)
	v := url.Values{}
	v.Set("_method", "PUT")
	v.Set("authAccount.ID", strconv.FormatInt(ruleId, 10))

	v.Set("authAccount.ManagerID", "3148472061957132038")

	v.Set("authAccount.LeaderID", "5151738548693216516")

	v.Set("authAccount.Name", "o90")

	v.Set("authAccount.Password", "twmdoajo4")

	v.Set("authAccount.Email", "Et consectetur est nihil sunt.")

	v.Set("authAccount.Sex", "Optio excepturi rem assumenda delectus illum.")

	v.Set("authAccount.Level", "Voluptas quasi quas ipsa.")

	v.Set("authAccount.Description", "Ex enim et et cupiditate pariatur blanditiis.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "1995-05-07T13:44:42+08:00")

	v.Set("authAccount.UpdatedAt", "1993-06-15T17:27:30+08:00")

	t.Post(t.ReverseUrl("AuthAccounts.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().Id(ruleId).Get(&authAccount)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(authAccount.ManagerID), v.Get("authAccount.ManagerID"))

	t.AssertEqual(fmt.Sprint(authAccount.LeaderID), v.Get("authAccount.LeaderID"))

	t.AssertEqual(fmt.Sprint(authAccount.Name), v.Get("authAccount.Name"))

	t.AssertEqual(fmt.Sprint(authAccount.Password), v.Get("authAccount.Password"))

	t.AssertEqual(fmt.Sprint(authAccount.Email), v.Get("authAccount.Email"))

	t.AssertEqual(fmt.Sprint(authAccount.Sex), v.Get("authAccount.Sex"))

	t.AssertEqual(fmt.Sprint(authAccount.Level), v.Get("authAccount.Level"))

	t.AssertEqual(fmt.Sprint(authAccount.Description), v.Get("authAccount.Description"))

	t.AssertEqual(fmt.Sprint(authAccount.Birthday), v.Get("authAccount.Birthday"))

}

func (t AuthAccountsTest) TestDelete() {
	t.ClearTable("tpt_auth_accounts")
	t.LoadFiles("tests/fixtures/auth_accounts.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)
	t.Delete(t.ReverseUrl("AuthAccounts.Delete", ruleId))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("tpt_auth_accounts", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}

func (t AuthAccountsTest) TestDeleteByIDs() {
	t.ClearTable("tpt_auth_accounts")
	t.LoadFiles("tests/fixtures/auth_accounts.yaml")
	//conds := EQU{"name": "这是一个规则名,请替换成正确的值"}
	conds := EQU{}
	ruleId := t.GetIDFromTable("tpt_auth_accounts", conds)
	t.Delete(t.ReverseUrl("AuthAccounts.DeleteByIDs", []interface{}{ruleId}))
	t.AssertStatus(http.StatusOK)
	//t.AssertContentType("application/json; charset=utf-8")
	count := t.GetCountFromTable("tpt_auth_accounts", nil)
	t.Assertf(count == 0, "count != 0, actual is %v", count)
}
