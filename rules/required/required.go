package required

import (
	"fmt"
	"github.com/tangpanqing/gin-validation/message"
)

var Required = "Required"

func VerifyRequired(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	if fieldValue.(string) == "" {
		return fmt.Sprintf(message.Required, fieldName)
	}
	return ""
}
