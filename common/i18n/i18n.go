package i18n

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"io/ioutil"
	"strings"
)

var DefaultLang string = "zh_CN"

var SupportedLanguages []string = []string{"en", "en-us", "en-US", "en_US", "zh_CN", "zh_cn", "zh-CN", "zh-cn"}

func T(language string) i18n.TranslateFunc {
	switch language {
	case "en", "en-us", "en-US", "en_US":
		language = "en-US"
	case "zh_CN", "zh_cn", "zh-CN", "zh-cn":
		language = "zh-CN"
	default:
		language = "zh-CN"
	}

	T, _ := i18n.Tfunc(language)
	return T
}

func ParseAcceptLanguage(language string) string {
	if language == "" || strings.Trim(language, " ") == "" {
		return DefaultLang
	}

	langs := strings.Split(language, ";")[0]
	if len(langs) == 0 {
		return DefaultLang
	}

	langList := strings.Split(langs, ",")
	if len(langList) == 0 {
		return DefaultLang
	}

	switch langList[0] {
	case "en", "en-us", "en-US", "en_US":
		return "en-US"
	case "zh_CN", "zh_cn", "zh-CN", "zh-cn", "zh-Hans":
		return "zh-CN"
	}

	return DefaultLang
}

func LoadLocaleMessageFiles(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := i18n.LoadTranslationFile(dir + "/" + file.Name()); err != nil {
			return err
		}
	}

	return nil
}
