package word

import (
	"strings"
	"unicode"
)


// 全转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 全转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}


// 下划线单词转大写驼峰单词
func UnderscoreToUpperCameLCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}


// 下划线 单词转小写驼峰单词
func UnderscoreToLowerCameLCase(s string) string {
	s = UnderscoreToUpperCameLCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}


//驼峰单词转下划线单词
func CameCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

