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

	v.Set("authAccount.ManagerID", "5770681614450777694")

	v.Set("authAccount.LeaderID", "5874364607433165177")

	v.Set("authAccount.Name", "5tu")

	v.Set("authAccount.Password", "venczsp0m")

	v.Set("authAccount.Email", "Maiores dolor asperiores voluptates dolorem corporis aliquid necessitatibus.")

	v.Set("authAccount.Sex", "Est sed possimus.")

	v.Set("authAccount.Level", "Nisi qui voluptas ut corporis id.")

	v.Set("authAccount.Description", "Velit at iste ut sit.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "2004-11-09T20:07:46+08:00")

	v.Set("authAccount.UpdatedAt", "1982-04-17T06:48:42+08:00")

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

	v.Set("authAccount.ManagerID", "8249874301148402771")

	v.Set("authAccount.LeaderID", "5873861573245505049")

	v.Set("authAccount.Name", "ton")

	v.Set("authAccount.Password", "13gqst015")

	v.Set("authAccount.Email", "Qui necessitatibus vel ad quia possimus repudiandae.")

	v.Set("authAccount.Sex", "Aut fugit qui ut ut reiciendis et est.")

	v.Set("authAccount.Level", "Aut et consectetur.")

	v.Set("authAccount.Description", "Fugit et labore suscipit aut sed provident id.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "2009-10-27T04:22:39+08:00")

	v.Set("authAccount.UpdatedAt", "1984-05-29T06:16:17+08:00")

	t.Post(t.ReverseUrl("AuthAccounts.Update", ruleId), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().ID(ruleId).Get(&authAccount)
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
