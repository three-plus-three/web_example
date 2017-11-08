package app

import (
	"github.com/revel/revel"
	"github.com/three-plus-three/forms"
	"github.com/three-plus-three/modules/environment"
	"github.com/three-plus-three/web_example/app/libs"
	"github.com/three-plus-three/web_example/app/models"
	"github.com/three-plus-three/web_example/app/routes"

	_ "github.com/three-plus-three/modules/bind"
	"github.com/three-plus-three/modules/toolbox"
	"github.com/three-plus-three/modules/web_ext"
)

var Lifecycle *libs.Lifecycle

func init() {
	web_ext.Init(environment.ENV_WSERVER_PROXY_ID, "例子",
		func(data *web_ext.Lifecycle) error {
			//if err := models.DropTables(data.ModelEngine); err != nil {
			//	return err
			//}
			if err := models.InitTables(data.ModelEngine); err != nil {
				return err
			}

			data.Variables["userLevel"] = []forms.InputChoice{{Value: "1", Label: "high"},
				{Value: "2", Label: "modium"},
				{Value: "3", Label: "low"}}
			revel.TemplateFuncs["userLevel_format"] = func(level string) string {
				switch level {
				case "1":
					return "high"
				case "2":
					return "modium"
				case "3":
					return "low"
				}
				return level
			}

			Lifecycle = &libs.Lifecycle{
				Lifecycle: data,
				DB:        models.DB{Engine: data.ModelEngine},
				DataDB:    models.DB{Engine: data.DataEngine},
			}
			return nil
		},
		func(data *web_ext.Lifecycle) ([]toolbox.Menu, error) {
			return []toolbox.Menu{
				{Title: "主页", UID: "Home", URL: routes.Home.Index()},
				{Title: "用户", URL: "#", Children: []toolbox.Menu{
					{Title: "用户账号", UID: "AuthAccounts", URL: routes.AuthAccounts.Index()},
					{Title: "在线用户", UID: "OnlineUsers", URL: routes.OnlineUsers.Index()},
				}}}, nil
		})

}
