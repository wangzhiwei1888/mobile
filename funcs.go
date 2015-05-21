package mobiler

import (
	"regexp"
)

const (
	MobileRegexpStrict string = "(130|131|132|133|134|135|136|137|138|139|145|147|150|151|152|153|155|156|157|158|159|170|176|177|178|180|181|182|183|184|185|186|187|188|189)[\\d]{8}"
	MobileRegexp       string = "(130|131|132|133|134|135|136|137|138|139|145|147|150|151|152|153|155|156|157|158|159|170|176|177|178|180|181|182|183|184|185|186|187|188|189)[*\\d]{4}[\\d]{4}"
)

func getRegexp(strict bool) (reg *regexp.Regexp) {
	var pattern string
	if strict {
		pattern = MobileRegexpStrict
	} else {
		pattern = MobileRegexp
	}
	reg = regexp.MustCompile(pattern)
	return
}

func validate(text string, strict bool) bool {
	reg := getRegexp(strict)
	return reg.MatchString(text)
}

func find(text string, strict bool) []string {
	reg := getRegexp(strict)
	return reg.FindAllString(text, -1)
}

func findOne(text string, strict bool) string {
	reg := getRegexp(strict)
	return reg.FindString(text)
}

// IsMobile 函数用于验证一个字符串是否为手机号码
func IsMobile(text string, strict bool) bool {
	if len(text) != 11 {
		return false
	}
	return validate(text, strict)
}

// HasMobile 函数用于验证一个字符串中是否存在手机号码
func HasMobile(text string, strict bool) bool {
	if len(text) < 11 {
		return false
	}
	mobiles := find(text, strict)
	return len(mobiles) > 0
}

// CountMobile 函数用于计算一个字符串中包含的手机号码数量
func CountMobile(text string, strict bool) int {
	return len(find(text, strict))
}

// GetMobiles 函数用于获取一个字符串中存在的所有手机号码
func GetMobiles(text string, strict bool) (mobiles []string) {
	if len(text) < 11 {
		return []string{}
	}
	mobiles = find(text, strict)
	return
}

// GetMobile 函数用于根据一个字符串获取其中的第一个手机号码
func GetMobile(text string, strict bool) (mobile string) {
	if len(text) < 11 {
		return ""
	}
	mobile = findOne(text, strict)
	return
}
