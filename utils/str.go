package files

import "strings"

/**
首字母转大写
 */
func StrFirstToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	strArr := []rune(str)
	if strArr[0] >= 97 && strArr[0] <= 122 {
		strArr[0] -= 32
	}
	return string(strArr)
}
/**
字符串转驼峰
 */
func StrToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	spi := strings.Split(str, "_")
	var tmp []string
	for _, s := range spi {
		tmp = append(tmp, strings.Title(s))
	}
	return strings.Join(tmp,"")
}