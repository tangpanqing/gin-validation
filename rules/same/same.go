package same

//Same 某字段必须等于另一个字段
func Same(otherField string) string {
	return "Same:" + otherField
}

func VerifySame(fieldName string, fieldValue any, fieldValueType string, ruleValue string, ruleFieldName string) string {
	if fieldValue.(string) != ruleValue {
		return fieldName + "的值必须等于" + ruleFieldName + "的值"
	}
	return ""
}
