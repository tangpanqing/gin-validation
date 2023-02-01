package different

//Different 某字段必须不等于另一个字段
func Different(otherField string) string {
	return "Different:" + otherField
}

func VerifyDifferent(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	if fieldValue.(string) == ruleValue {
		return fieldName + "的值必须不等于" + ruleFieldName + "的值"
	}
	return ""
}
