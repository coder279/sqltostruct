package word

import (
	"strings"
	"unicode"
)

// ToUpper 将传递的单词转为大写
func ToUpper(s string) string{
	return strings.ToUpper(s)
}

// ToLower 将传递的单词转为小写
func ToLower(s string) string{
	return strings.ToLower(s)
}

// UnderScoreToUpperCamelCase 下划线转为大写驼峰单词
func UnderScoreToUpperCamelCase(s string) string{
	s = strings.Replace(s,"_"," ",-1)
	s = strings.Title(s)
	return strings.Replace(s," ","",-1)
}

// UnderScoreToLowerCamelCase 下划线转为大写驼峰单词
func UnderScoreToLowerCamelCase(s string) string{
	s = UnderScoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderScore 驼峰单词转为下划线
func CamelCaseToUnderScore(s string) string {
	var output []rune
	for i,r := range s {
		if i == 0 {
			output = append(output,unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r){
			output = append(output,'_')
		}
		output = append(output,unicode.ToLower(r))
	}
	return string(output)
}