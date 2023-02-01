package lenIn

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"strings"
)

func LenIn(val ...any) string {
	var valArr []string
	for i := 0; i < len(val); i++ {
		valArr = append(valArr, fmt.Sprintf("%v", val[i]))
	}

	return "LenIn:" + strings.Join(valArr, ",")
}

func VerifyLenIn(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	ruleValArr := strings.Split(ruleValue, ",")
	thisValue := helper.GetThisInt64ValueLen(fieldValue, fieldValueType)

	if !helper.IsInArr(helper.Int64ToStr(thisValue), fieldValueType, ruleValArr) {
		return fieldName + "的长度必须是" + ruleValue + "其中之一"
	}
	return ""
}
