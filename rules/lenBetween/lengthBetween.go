package lenBetween

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LenBetween(min int64, max int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", min))
	valArr = append(valArr, fmt.Sprintf("%v", max))

	return "LenBetween:" + strings.Join(valArr, ",")
}

func VerifyLenBetween(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	ruleValArr := strings.Split(ruleValue, ",")
	thisValue := helper.GetThisInt64ValueLen(fieldValue, fieldValueType)

	if thisValue < helper.Str2Int64(ruleValArr[0]) || thisValue > helper.Str2Int64(ruleValArr[1]) {
		return fmt.Sprintf(message.LenBetween, fieldName, helper.Str2Int64(ruleValArr[0]), helper.Str2Int64(ruleValArr[1]))
	}

	return ""
}
