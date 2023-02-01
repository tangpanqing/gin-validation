package in

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"strings"
)

func In(val ...any) string {
	var valArr []string
	for i := 0; i < len(val); i++ {
		valArr = append(valArr, fmt.Sprintf("%v", val[i]))
	}

	return "In:" + strings.Join(valArr, ",")
}

func VerifyIn(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleValArr := strings.Split(ruleValue, ",")
	if !helper.IsInArr(fieldValue, fieldValueType, ruleValArr) {
		return fieldName + "必须是" + ruleValue + "其中之一"
	}

	return ""
}
