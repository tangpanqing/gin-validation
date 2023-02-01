package lenLt

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LenLt(length int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", length))

	return "LenLt:" + strings.Join(valArr, ",")
}

func VerifyLenLt(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	ruleLen := helper.Str2Int64(ruleValue)
	valueLen := helper.GetThisInt64ValueLen(fieldValue, fieldValueType)
	if !(valueLen < ruleLen) {
		return fmt.Sprintf(message.LenLt, fieldName, ruleLen)
	}
	return ""
}
