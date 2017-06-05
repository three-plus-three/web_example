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

	v.Set("authAccount.ManagerID", "2891689786100605429")

	v.Set("authAccount.LeaderID", "2277565402848467490")

	v.Set("authAccount.Name", "zcj")

	v.Set("authAccount.Password", "uz5w7k7pu")

	v.Set("authAccount.Email", "Voluptatem libero iste nam praesentium.")

	v.Set("authAccount.Sex", "Sunt doloribus eveniet sit omnis.")

	v.Set("authAccount.Level", "Cumque nam sunt voluptates cupiditate earum.")

	v.Set("authAccount.Description", "Quaerat quae non ea voluptatem sint qui.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "1997-06-07T13:15:21+08:00")

	v.Set("authAccount.UpdatedAt", "1978-02-25T02:41:05+08:00")

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

	v.Set("authAccount.ManagerID", "5970798000744751003")

	v.Set("authAccount.LeaderID", "229706376718258173")

	v.Set("authAccount.Name", "ylc")

	v.Set("authAccount.Password", "ch0dxhd6a")

	v.Set("authAccount.Email", "Nemo doloremque molestiae dolorem beatae enim velit.")

	v.Set("authAccount.Sex", "Id eius voluptatem fugiat omnis est neque temporibus.")

	v.Set("authAccount.Level", "Recusandae laborum non voluptatum sapiente eos atque.")

	v.Set("authAccount.Description", "Molestias sunt quae et iure.")

	v.Set("authAccount.Birthday", "abc")

	v.Set("authAccount.CreatedAt", "1995-06-11T22:40:22+08:00")

	v.Set("authAccount.UpdatedAt", "2001-12-12T14:07:04+08:00")

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
