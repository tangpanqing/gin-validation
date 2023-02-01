package lengthEq

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LengthEq(length int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", length))

	return "LengthEq:" + strings.Join(valArr, ",")
}

func VerifyLengthEq(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleLength := helper.Str2Int64(ruleValue)
	valueLength := helper.GetThisInt64ValueLength(fieldValue, fieldValueType)
	if valueLength != ruleLength {
		return fmt.Sprintf(message.LengthEq, fieldName, ruleLength)
	}
	return ""
}
