package app

import (
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/web_example/app/libs"
	"github.com/three-plus-three/web_example/app/routes"
)

func initMenuList(lifecycle *libs.Lifecycle) {
	lifecycle.MenuList = []toolbox.Menu{
		{Title: "主页", Name: "Home", URL: routes.Home.Index()},
		{Title: "用户", URL: "#", Children: []toolbox.Menu{
			{Title: "用户账号", Name: "AuthAccounts", URL: routes.AuthAccounts.Index()},
			{Title: "在线用户", Name: "OnlineUsers", URL: routes.OnlineUsers.Index()},
		}},
	}
}
