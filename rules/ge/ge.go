package ge

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
)

func Ge(val any) string {
	return "Ge:" + fmt.Sprintf("%v", val)
}

func VerifyGe(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	flag := (fieldValueType == "Int64" && !(fieldValue.(int64) >= helper.Str2Int64(ruleValue))) ||
		(fieldValueType == "Float64" && !(fieldValue.(float64) >= helper.Str2Float64(ruleValue)))

	if flag {
		return fieldName + "必须大于等于" + ruleValue
	}

	return ""
}
