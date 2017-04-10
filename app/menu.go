package app

import (
	"github.com/three-plus-three/web_example/app/libs"
	"github.com/three-plus-three/web_example/app/routes"
	web_app "github.com/three-plus-three/web_templates/app"
)

func initMenuList(lifecycle *libs.Lifecycle) {
	lifecycle.MenuList = []web_app.Menu{
		{Title: "主页", Name: "Home", URL: routes.Home.Index()},
		{Title: "用户", URL: "#", Children: []web_app.Menu{
			{Title: "用户账号", Name: "AuthAccounts", URL: routes.AuthAccounts.Index(0, 0)},
			{Title: "在线用户", Name: "OnlineUsers", URL: routes.OnlineUsers.Index(0, 0)},
		}},
	}
}
