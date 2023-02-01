package ne

import (
	"fmt"
)

func Ne(val any) string {
	return "Ne:" + fmt.Sprintf("%v", val)
}

func VerifyNe(fieldName string, fieldValue any, fieldValueType string, ruleValue string) string {
	if !(fmt.Sprintf("%v", fieldValue) != ruleValue) {
		return fieldName + "必须不等于" + ruleValue
	}
	return ""
}
