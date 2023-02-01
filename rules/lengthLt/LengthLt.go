package lengthLt

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LengthLt(length int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", length))

	return "LengthLt:" + strings.Join(valArr, ",")
}

func VerifyLengthLt(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleLength := helper.Str2Int64(ruleValue)
	valueLength := helper.GetThisInt64ValueLength(fieldValue, fieldValueType)
	if !(valueLength < ruleLength) {
		return fmt.Sprintf(message.LengthLt, fieldName, ruleLength)
	}
	return ""
}
