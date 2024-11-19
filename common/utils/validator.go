package utils

import (
	"net"
	"regexp"
	"strings"
	"unicode"
)

func ValidatorEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func ValidatorMobile(mobile string) bool {
	regular := "^((1[2,3][0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|(19[0-9])|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

func ValidatorName(name string) bool {
	const pattern = `^[A-Za-z0-9-_]+$`
	match, _ := regexp.MatchString(pattern, name)
	if !match {
		return false
	}
	const patternSymbol = `^[-_]*$`
	matchSymbol, _ := regexp.MatchString(patternSymbol, name)
	if matchSymbol {
		return false
	}
	return true
}

func ValidatorCode(code string) bool {
	if len(code) < 4 || len(code) > 50 {
		return false
	}
	const pattern = `^[a-zA-Z0-9-_:]+$`
	match, _ := regexp.MatchString(pattern, code)
	if !match {
		return false
	}

	const patternSymbol = `^[-_:]*$`
	matchSymbol, _ := regexp.MatchString(patternSymbol, code)
	if matchSymbol {
		return false
	}

	return match
}

func ValidatorChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

// ValidatorPassword 密码中允许出现数字、大写字母、小写字母、特殊字符（.@$!%*#_~?&^）
// 包含其中2种字符，且长度在8-16之间
func ValidatorPassword(pwd string, minLen, maxLen int) bool {
	if strings.Contains(pwd, " ") {
		return false
	}
	if len(pwd) < minLen || len(pwd) > maxLen {
		return false
	}
	// 过滤掉这四类字符以外的密码串,直接判断不合法
	re, err := regexp.Compile(`^[a-zA-Z0-9.@$!%*#_~?&^]{8,16}$`)
	if err != nil {
		return false
	}
	match := re.MatchString(pwd)
	if !match {
		return false
	}

	var level = 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[.@$!%*#_~?&^]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	// 包含两种以上字符即可
	if level <= 2 {
		return false
	}

	return true
}

// ValidatorUserName min-max 4-32
func ValidatorUserName(name string, min, max int) bool {
	if name == "" {
		return false
	}
	if strings.Contains(name, " ") {
		return false
	}
	if len(name) <= min || len(name) > max {
		return false
	}

	pattern := `[.@$!%*#_~?&^]+`
	match, err := regexp.MatchString(pattern, name)
	if err != nil || match {
		return false
	}
	return true
}

func ValidatorSpecialSymbol(str string) bool {
	if str == "" {
		return false
	}
	if strings.Contains(str, " ") {
		return false
	}

	pattern := `[.@$!%*#_~?&^]+`
	match, err := regexp.MatchString(pattern, str)
	if err != nil || match {
		return false
	}
	return true
}

func ValidatorIp(str string) bool {
	address := net.ParseIP(str)
	if address == nil {
		return false
	}
	return true
}

func ValidatorPort(port int64) bool {
	if port < 1 || port > 65535 {
		return false
	}
	return true
}

func ContainsStr(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
