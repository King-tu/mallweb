package utils

import (
	"fmt"
	"regexp"

	"github.com/king-tu/mallweb/common"
)

func ValidateIdCardFormat(idCardType, idCardNo, idCardTypeColumn, idCardNoColumn string) []string {
	if idCardTypeColumn == "" {
		idCardTypeColumn = "证照类型"
	}
	if idCardNoColumn == "" {
		idCardNoColumn = "证照号码"
	}

	results := []string{}

	if len(idCardNo) == 0 {
		results = append(results, fmt.Sprintf("%s列为空", idCardNoColumn))
	}

	if len(idCardType) == 0 {
		results = append(results, fmt.Sprintf("%s列为空", idCardTypeColumn))
	}

	if len(idCardType) > 0 && !Contains(idCardType, common.IDCARD_TYPES) {
		results = append(results, fmt.Sprintf("%s输入不合法", idCardTypeColumn))
		return results
	}

	switch idCardType {
	case "居民身份证":
		if !ValidateIdCard(idCardNo) {
			results = append(results, fmt.Sprintf("%s列居民身份证号码输入不合法", idCardNoColumn))
		}
	case "中国护照":
		pat := "^[0-9a-zA-Z]{9}$"
		matched, _ := regexp.MatchString(pat, idCardNo)
		if !matched {
			results = append(results, fmt.Sprintf("%s列中国护照只能为9位数字和字母", idCardNoColumn))
		}
	case "港澳居民居住证", "台湾居民居住证":
		pat := "^\\d{18}$"
		matched, _ := regexp.MatchString(pat, idCardNo)
		if !matched {
			results = append(results, fmt.Sprintf("%s列%s只能为18位数字", idCardNoColumn, idCardType))
		}
	case "外国人永久居留身份证":
		pat := "^[0-9a-zA-Z]{15}$"
		matched, _ := regexp.MatchString(pat, idCardNo)
		if !matched {
			results = append(results, fmt.Sprintf("%s列外国人永久居留身份证只能为15位数字和字母", idCardNoColumn))
		}
	default:
		if len(idCardNo) > 20 {
			results = append(results, fmt.Sprintf("%s列不能超过20个字符", idCardNoColumn))
		}
	}

	return results
}
