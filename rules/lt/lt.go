package lt

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
)

func Lt(val any) string {
	return "Lt:" + fmt.Sprintf("%v", val)
}

func VerifyLt(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	flag := (fieldValueType == "Int64" && !(fieldValue.(int64) < helper.Str2Int64(ruleValue))) ||
		(fieldValueType == "Float64" && !(fieldValue.(float64) < helper.Str2Float64(ruleValue)))

	if flag {
		return fieldName + "必须小于" + ruleValue
	}
	return ""
}
