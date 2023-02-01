package between

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"strings"
)

func Between(min any, max any) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", min))
	valArr = append(valArr, fmt.Sprintf("%v", max))

	return "Between:" + strings.Join(valArr, ",")
}

func VerifyBetween(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	ruleValueArr := strings.Split(ruleValue, ",")

	if fieldValueType == "Int64" {
		thisValue := fieldValue.(int64)
		if thisValue < helper.Str2Int64(ruleValueArr[0]) || thisValue > helper.Str2Int64(ruleValueArr[1]) {
			return fieldName + "必须在" + ruleValueArr[0] + "与" + ruleValueArr[1] + "之间"
		}
	}

	if fieldValueType == "Float64" {
		thisValue := fieldValue.(float64)
		if thisValue < helper.Str2Float64(ruleValueArr[0]) || thisValue > helper.Str2Float64(ruleValueArr[1]) {
			return fieldName + "必须在" + ruleValueArr[0] + "与" + ruleValueArr[1] + "之间"
		}
	}

	return ""
}
