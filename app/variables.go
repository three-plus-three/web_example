package app

import (
	"bytes"
	"cn/com/hengwei/commons"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/revel/revel"
	"github.com/three-plus-three/forms"
)

func readVariables(env *commons.Environment) map[string]interface{} {
	variables := map[string]interface{}{
		"urlPrefix":           env.DaemonUrlPath,
		"application_context": env.DaemonUrlPath,
		"application_catalog": env.Config.StringWithDefault("application.catalog", "all"),
		"head_title_text": readFileWithDefault([]string{
			env.Fs.FromDataConfig("resources/profiles/header.txt"),
			env.Fs.FromData("resources/profiles/header.txt"),
			filepath.Join(os.Getenv("hw_root_dir"), "data/resources/profiles/header.txt")}, "IT综合运维管理平台"),
		"footer_title_text": readFileWithDefault([]string{
			env.Fs.FromDataConfig("resources/profiles/footer.txt"),
			env.Fs.FromData("resources/profiles/footer.txt"),
			filepath.Join(os.Getenv("hw_root_dir"), "data/resources/profiles/footer.txt")}, "© 2017 恒维信息技术(上海)有限公司, 保留所有版权。"),
	}

	variables["userLevel"] = []forms.InputChoice{{Value: "1", Label: "high"},
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
	return variables
}

func readFileWithDefault(files []string, defaultValue string) string {
	for _, s := range files {
		content, e := ioutil.ReadFile(s)
		if nil == e {
			if content = bytes.TrimSpace(content); len(content) > 0 {
				return string(content)
			}
		}
	}
	return defaultValue
}
