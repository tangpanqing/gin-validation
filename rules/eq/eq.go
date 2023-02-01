package eq

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/message"
)

func Eq(val any) string {
	return "Eq:" + fmt.Sprintf("%v", val)
}

func VerifyEq(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	if !(fmt.Sprintf("%v", fieldValue) == ruleValue) {
		return fmt.Sprintf(message.Eq, fieldName, ruleValue)
	}
	return ""
}
