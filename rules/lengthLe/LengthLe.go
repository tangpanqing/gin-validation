package lengthLe

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LengthLe(length int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", length))

	return "LengthLe:" + strings.Join(valArr, ",")
}

func VerifyLengthLe(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleLength := helper.Str2Int64(ruleValue)
	valueLength := helper.GetThisInt64ValueLength(fieldValue, fieldValueType)
	if !(valueLength <= ruleLength) {
		return fmt.Sprintf(message.LengthLe, fieldName, ruleLength)
	}
	return ""
}
