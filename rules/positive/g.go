package positive

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/message"
)

var Positive = "Positive"

func VerifyPositive(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	if fieldValue.(int64) <= 0 {
		return fmt.Sprintf(message.Positive, fieldName)
	}
	return ""
}
