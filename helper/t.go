package helper

import (
	"strconv"
)

func Str2Int64(str string) int64 {
	dataNew, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return dataNew
}

func Str2Float64(str string) float64 {
	dataNew, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return dataNew
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func GetThisStringValue(value any, valueType string) string {
	var thisValue string

	if valueType == "String" {
		thisValue = value.(string)
	}

	if valueType == "Int64" {
		thisValue = Int64ToStr(value.(int64))
	}

	if valueType == "Float64" {
		thisValue = Float64ToStr(value.(float64))
	}

	return thisValue
}

func GetThisInt64ValueLen(value any, valueType string) int64 {
	var thisValue int64

	if valueType == "String" {
		thisValue = int64(len(value.(string)))
	}

	return thisValue
}

func IsInArr(fieldValue any, fieldValueType string, ruleVal []string) bool {

	inFlag := false

	for j := 0; j < len(ruleVal); j++ {
		var flag bool

		if fieldValueType == "String" {
			flag = fieldValue.(string) == ruleVal[j]
		}

		if fieldValueType == "Int64" {
			flag = fieldValue.(int64) == Str2Int64(ruleVal[j])
		}

		if fieldValueType == "Float64" {
			flag = fieldValue.(float64) == Str2Float64(ruleVal[j])
		}

		if flag {
			inFlag = true
			break
		}
	}

	return inFlag
}

func getThisStringValue(value any, valueType string) string {
	var thisValue string

	if valueType == "String" {
		thisValue = value.(string)
	}

	if valueType == "Int64" {
		thisValue = Int64ToStr(value.(int64))
	}

	return thisValue
}

func GetThisInt64Value(value any, valueType string) int64 {
	var thisValue int64

	if valueType == "String" {
		thisValue = Str2Int64(value.(string))
	}

	if valueType == "Int64" {
		thisValue = value.(int64)
	}

	return thisValue
}
