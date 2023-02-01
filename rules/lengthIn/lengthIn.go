package lengthIn

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"strings"
)

func LengthIn(val ...any) string {
	var valArr []string
	for i := 0; i < len(val); i++ {
		valArr = append(valArr, fmt.Sprintf("%v", val[i]))
	}

	return "LengthIn:" + strings.Join(valArr, ",")
}

func VerifyLengthIn(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleValArr := strings.Split(ruleValue, ",")
	thisValue := helper.GetThisInt64ValueLength(fieldValue, fieldValueType)

	if !helper.IsInArr(helper.Int64ToStr(thisValue), fieldValueType, ruleValArr) {
		return fieldName + "的长度必须是" + ruleValue + "其中之一"
	}
	return ""
}
