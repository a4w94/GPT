package prompt

const DefaultLanguage = "English"

var languageMaps = map[string]string{
	"en":    DefaultLanguage,
	"zh-tw": "Traditional Chinese",
	"zh-cn": "Simplified Chinese",
	"ja":    "Japanese",
}

// 取得語系代號對應詳細語系名稱
func GetLanguage(langCode string) string {
	if language, ok := languageMaps[langCode]; ok {
		return language
	}
	return DefaultLanguage
}
