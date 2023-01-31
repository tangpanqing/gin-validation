package validation

import (
	"fmt"
	"strings"
)

//Required 某字段必须存在,不能为空字符串
var Required = "Required"

//Positive 某字段必须是正数
var Positive = "Positive"

//Same 某字段必须等于另一个字段
func Same(other string) string {
	return "Same:" + other
}

//Different 某字段必须不等于另一个字段
func Different(other string) string {
	return "Different:" + other
}

//Between 某字段必须在两个值之间
func Between(min int64, max int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", min))
	valArr = append(valArr, fmt.Sprintf("%v", max))

	return "Between:" + strings.Join(valArr, ",")
}

//In 某字段必须是其中某一个
func In(val ...any) string {
	var valArr []string
	for i := 0; i < len(val); i++ {
		valArr = append(valArr, fmt.Sprintf("%v", val[i]))
	}

	return "In:" + strings.Join(valArr, ",")
}

//Min 某字段必须大于等于某值
func Min(min int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", min))

	return "Min:" + strings.Join(valArr, ",")
}

//Max 某字段必须小于等于某值
func Max(max int64) string {
	var valArr []string
	valArr = append(valArr, fmt.Sprintf("%v", max))

	return "Max:" + strings.Join(valArr, ",")
}
