package middleware

import (
	ini "gopkg.in/ini.v1"
)

var (
	iniFileMap = make(map[string]*ini.File)
)

// InitConfig ...
func InitConfig() {
	data, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, "config/application.ini")

	if err != nil {
		panic(err)
	}
	iniFileMap["application"] = data
}

func GetConfig(key string) string {
	var (
		iniFile *ini.File
		ok      bool
	)

	if iniFile, ok = iniFileMap["application"]; !ok {
		return ""
	}

	if res := iniFile.Section("common").Key(key).String(); len(res) > 0 {
		return res
	}

	return ""
}
