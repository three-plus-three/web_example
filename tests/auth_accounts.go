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

	t.AssertContains(fmt.Sprint(authAccount.Name))
	t.AssertContains(fmt.Sprint(authAccount.Email))
	t.AssertContains(fmt.Sprint(authAccount.Sex))
	t.AssertContains(fmt.Sprint(authAccount.Description))
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

	v.Set("authAccount.Name", "uz9")

	v.Set("authAccount.Password", "ktzfplixz")

	v.Set("authAccount.Email", "Magnam recusandae est assumenda tenetur.")

	v.Set("authAccount.Sex", "Et est id tempora quia quod sapiente et.")

	v.Set("authAccount.Description", "Itaque sunt temporibus alias.")

	v.Set("authAccount.CreatedAt", "2002-08-02T09:56:40+08:00")

	v.Set("authAccount.UpdatedAt", "1999-12-11T22:28:45+08:00")

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

	t.AssertEqual(fmt.Sprint(authAccount.Name), v.Get("authAccount.Name"))
	t.AssertEqual(fmt.Sprint(authAccount.Password), v.Get("authAccount.Password"))
	t.AssertEqual(fmt.Sprint(authAccount.Email), v.Get("authAccount.Email"))
	t.AssertEqual(fmt.Sprint(authAccount.Sex), v.Get("authAccount.Sex"))
	t.AssertEqual(fmt.Sprint(authAccount.Description), v.Get("authAccount.Description"))
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

	t.AssertContains(fmt.Sprint(authAccount.Name))
	t.AssertContains(fmt.Sprint(authAccount.Email))
	t.AssertContains(fmt.Sprint(authAccount.Sex))
	t.AssertContains(fmt.Sprint(authAccount.Description))
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

	v.Set("authAccount.Name", "ndf")

	v.Set("authAccount.Password", "humdt8v3q")

	v.Set("authAccount.Email", "Necessitatibus non atque non accusantium alias quis.")

	v.Set("authAccount.Sex", "A vel natus inventore reprehenderit ratione totam.")

	v.Set("authAccount.Description", "Eum est molestias culpa aliquam sunt.")

	v.Set("authAccount.CreatedAt", "2001-03-03T16:43:18+08:00")

	v.Set("authAccount.UpdatedAt", "2009-11-14T11:23:55+08:00")

	t.Post(t.ReverseUrl("AuthAccounts.Update"), "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	t.AssertOk()

	var authAccount models.AuthAccount
	err := app.Lifecycle.DB.AuthAccounts().Id(ruleId).Get(&authAccount)
	if err != nil {
		t.Assertf(false, err.Error())
	}

	t.AssertEqual(fmt.Sprint(authAccount.Name), v.Get("authAccount.Name"))

	t.AssertEqual(fmt.Sprint(authAccount.Password), v.Get("authAccount.Password"))

	t.AssertEqual(fmt.Sprint(authAccount.Email), v.Get("authAccount.Email"))

	t.AssertEqual(fmt.Sprint(authAccount.Sex), v.Get("authAccount.Sex"))

	t.AssertEqual(fmt.Sprint(authAccount.Description), v.Get("authAccount.Description"))

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
