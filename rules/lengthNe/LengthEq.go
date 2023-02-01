package lengthNe

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/message"
	"strings"
)

func LengthNe(length int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", length))

	return "LengthNe:" + strings.Join(valArr, ",")
}

func VerifyLengthNe(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	ruleLength := helper.Str2Int64(ruleValue)
	valueLength := helper.GetThisInt64ValueLength(fieldValue, fieldValueType)
	if !(valueLength != ruleLength) {
		return fmt.Sprintf(message.LengthNe, fieldName, ruleLength)
	}
	return ""
}
