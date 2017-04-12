package app

import (
	"cn/com/hengwei/commons"
	"crypto/sha1"
	"log"
	"os"
	"path/filepath"

	"github.com/three-plus-three/web_example/app/libs"
	"github.com/three-plus-three/web_example/app/models"

	"github.com/revel/revel"
	"github.com/three-plus-three/sessions"
	sso "github.com/three-plus-three/sso/client"
)

var Lifecycle *libs.Lifecycle

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.HTTPMethodOverride,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		SessionFilter,                 // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		GlobalVariablesFilter,         // set global variables
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(func() {
		env, err := commons.NewEnvironment(commons.Options{Name: "nsm",
			ConfDir: filepath.Join(os.Getenv("hw_root_dir"), "conf")})
		if nil != err {
			log.Println(err)
			os.Exit(-1)
			return
		}
		if revel.RunMode == "test" {
			env.Db.Models.Schema = env.Db.Models.Schema + "_models_test"
			env.Db.Data.Schema = env.Db.Data.Schema + "_test"
		}

		lifecycle, err := libs.NewLifecycle(env, nil)
		if nil != err {
			log.Println(err)
			os.Exit(-1)
			return
		}
		Lifecycle = lifecycle
		Lifecycle.URLPrefix = env.DaemonUrlPath
		lifecycle.Variables = readVariables(env)
		Lifecycle.CheckUser = initSSO(env)

		if revel.DevMode {
			Lifecycle.DB.Engine.ShowSQL()
		}
		if err := models.InitTables(Lifecycle.DB.Engine); err != nil {
			log.Println(err)
			os.Exit(-1)
			return
		}

		revel.AppRoot = "/" + env.Config.StringWithDefault("daemon.urlpath", "hengwei") + "/aaa"
		//revel.Config.SetOption("app.secret", Env.Config.StringWithDefault("app.secret", ""))
		revel.Config.SetOption("cookie.prefix", "PLAY")
		revel.Config.SetOption("cookie.path", env.RawDaemonUrlPath)
		revel.CookiePrefix = "PLAY"

		var secretKey []byte
		if secretStr := env.Config.StringWithDefault("app.secret", ""); secretStr != "" {
			secretKey = []byte(secretStr)
		}
		GlobalSessionFilter = sessions.SessionFilter(sso.DefaultSessionKey,
			env.RawDaemonUrlPath, sha1.New, secretKey)

		initTemplateFuncs(env)
	}, 0)

	revel.OnAppStart(func() {
		initMenuList(Lifecycle)
	}, 2)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var GlobalSessionFilter revel.Filter

func SessionFilter(c *revel.Controller, filterChain []revel.Filter) {
	//if GlobalSessionFilter != nil {
	GlobalSessionFilter(c, filterChain)
	//}
}

// GlobalVariablesFilter will set global variables
func GlobalVariablesFilter(c *revel.Controller, fc []revel.Filter) {
	// Make global vars available in templates as {{.global.xyz}}
	c.ViewArgs["global"] = Lifecycle.Variables

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
