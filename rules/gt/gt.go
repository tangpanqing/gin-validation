package gt

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
)

func Gt(val any) string {
	return "Gt:" + fmt.Sprintf("%v", val)
}

func VerifyGt(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	flag := (fieldValueType == "Int64" && !(fieldValue.(int64) > helper.Str2Int64(ruleValue))) ||
		(fieldValueType == "Float64" && !(fieldValue.(float64) > helper.Str2Float64(ruleValue)))

	if flag {
		return fieldName + "必须大于" + ruleValue
	}
	return ""
}
