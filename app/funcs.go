package app

import (
	"cn/com/hengwei/commons"
	"cn/com/hengwei/commons/httputils"
	"fmt"
	"strings"

	"github.com/revel/revel"
	"github.com/three-plus-three/forms"
)

func ResourcesURLFor(s string) string {
	return httputils.JoinURLPath(revel.AppRoot, "/internal/custom_resources/", s)
}

func initTemplateFuncs(env *commons.Environment) {
	revel.TemplateFuncs["assets"] = func(value string) string {
		return httputils.JoinURLPath(revel.AppRoot, "assets", value)
	}
	revel.TemplateFuncs["custom_resources_assets"] = func(value string) string {
		return ResourcesURLFor(value)
	}
	revel.TemplateFuncs["mc_assets"] = func(url string) string {
		return httputils.JoinURLPath(revel.AppRoot, url)
	}
	revel.TemplateFuncs["tpt_assets"] = func(value string) string {
		return httputils.JoinURLPath(revel.AppRoot, "tpt_assets", value)
	}
	revel.TemplateFuncs["default"] = func(value, defvalue interface{}) interface{} {
		if nil == value {
			return defvalue
		}
		if s, ok := value.(string); ok && "" == s {
			return defvalue
		}
		return value
	}

	revel.TemplateFuncs["tostring"] = func(value interface{}) interface{} {
		if nil == value {
			return ""
		}
		if s, ok := value.(string); ok {
			return s
		}
		return fmt.Sprint(value)
	}

	revel.TemplateFuncs["args"] = func() map[string]interface{} {
		return map[string]interface{}{}
	}

	revel.TemplateFuncs["arg"] = func(n string, v interface{}, args map[string]interface{}) map[string]interface{} {
		args[n] = v

		return args
	}

	revel.TemplateFuncs["list"] = func() []interface{} {
		return []interface{}{}
	}

	revel.TemplateFuncs["startsWith"] = func(s, sep string) bool {
		return strings.Index(s, sep) == 0
	}

	revel.TemplateFuncs["tabItem"] = func(id, label string, active bool, items []interface{}) []interface{} {
		return append(items, map[string]interface{}{"id": id, "label": label, "active": active})
	}

	revel.TemplateFuncs["menuItem"] = func(id string, class string, label string, items []interface{}) []interface{} {
		return append(items, map[string]interface{}{id: id, "class": class, "label": label})
	}

	revel.TemplateFuncs["urlPrefix"] = func() string {
		return revel.AppRoot
	}

	funcs := HtmlFuncMap()
	for k, v := range funcs {
		if _, ok := revel.TemplateFuncs[k]; !ok {
			revel.TemplateFuncs[k] = v
		}
	}
	for _, alias := range [][2]string{{"sum", "add"}} {
		if _, ok := revel.TemplateFuncs[alias[0]]; !ok {
			revel.TemplateFuncs[alias[0]] = funcs[alias[1]]
		}
	}

	forms.Init(revel.DevMode, revel.SourcePath, revel.TemplateFuncs)
	for k, v := range forms.FieldFuncs {
		revel.TemplateFuncs[k] = v
	}

	revel.TemplateFuncs["current_user_has_permission"] = CurrentUserHasPermission
	revel.TemplateFuncs["current_user_has_new_permission"] = func(ctx map[string]interface{}, permission string) bool {
		return CurrentUserHasPermission(ctx, permission+".new")
	}
	revel.TemplateFuncs["current_user_has_del_permission"] = func(ctx map[string]interface{}, permission string) bool {
		return CurrentUserHasPermission(ctx, permission+".del")
	}
	revel.TemplateFuncs["current_user_has_edit_permission"] = func(ctx map[string]interface{}, permission string) bool {
		return CurrentUserHasPermission(ctx, permission+".edit")
	}
	revel.TemplateFuncs["current_user_has_write_permission"] = func(ctx map[string]interface{}, permission string) bool {
		for _, tag := range []string{"edit", "new", "del"} {
			if CurrentUserHasPermission(ctx, permission+"."+tag) {
				return true
			}
		}
		return false
	}
	revel.TemplateFuncs["user_has_permission"] = UserHasPermission
}

func UserHasPermission(ctx map[string]interface{}, user, permission string) bool {
	return true
}

func CurrentUserHasPermission(ctx map[string]interface{}, permission string) bool {
	return true
}
