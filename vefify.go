package validation

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Validation struct {
	ginContext *gin.Context
	err        string
}

func (v *Validation) Error() string {
	return v.err
}

func (v *Validation) KeyString(key string, rules ...string) string {
	value := ""

	valueAny, exists := v.ginContext.Get(key)
	if exists {
		value = valueAny.(string)
	}

	v.handleValue(key, value, "Key", "String", rules...)
	return value
}

func (v *Validation) KeyInt64(key string, rules ...string) int64 {
	var value int64

	valueAny, exists := v.ginContext.Get(key)
	if exists {
		value = str2Int64(valueAny.(string))
	}

	v.handleValue(key, value, "Key", "Int64", rules...)
	return value
}

func (v *Validation) PostString(key string, rules ...string) string {
	value := v.ginContext.PostForm(key)
	v.handleValue(key, value, "Post", "String", rules...)
	return value
}

func (v *Validation) PostInt64(key string, rules ...string) int64 {
	valueStr := v.ginContext.PostForm(key)
	value := str2Int64(valueStr)
	v.handleValue(key, value, "Post", "Int64", rules...)
	return value
}

func (v *Validation) handleValue(key string, value interface{}, valueFrom string, valueType string, rules ...string) {
	if len(rules) > 0 {
		for i := 0; i < len(rules); i++ {
			ruleInfo := strings.Split(rules[i], ":")

			if ruleInfo[0] == Required {
				if value.(string) == "" {
					v.err = key + "不能为空"
				}
			}

			if ruleInfo[0] == Positive {
				if value.(int64) <= 0 {
					v.err = key + "必须大于0"
				}
			}

			if ruleInfo[0] == "Same" {
				if value.(string) != v.getOtherValue(valueFrom, ruleInfo[1]) {
					v.err = key + "的值必须等于" + ruleInfo[1] + "的值"
				}
			}

			if ruleInfo[0] == "Different" {
				if value.(string) == v.getOtherValue(valueFrom, ruleInfo[1]) {
					v.err = key + "的值必须不等于" + ruleInfo[1] + "的值"
				}
			}

			if ruleInfo[0] == "Min" {
				if value.(int64) < str2Int64(ruleInfo[1]) {
					v.err = key + "必须大于等于" + ruleInfo[1]
				}
			}

			if ruleInfo[0] == "Max" {
				if value.(int64) > str2Int64(ruleInfo[1]) {
					v.err = key + "必须小于等于" + ruleInfo[1]
				}
			}

			if ruleInfo[0] == "Between" {
				ruleValArr := strings.Split(ruleInfo[1], ",")
				thisValue := getThisValue(value, valueType)

				if thisValue < str2Int64(ruleValArr[0]) || thisValue > str2Int64(ruleValArr[1]) {
					v.err = key + "必须在" + ruleValArr[0] + "与" + ruleValArr[1] + "之间"
				}
			}

			if ruleInfo[0] == "In" {
				ruleValArr := strings.Split(ruleInfo[1], ",")
				thisValue := getThisValue(value, valueType)

				if !isInArr(thisValue, ruleValArr) {
					v.err = key + "必须是" + ruleInfo[1] + "其中之一"
				}
			}
		}
	}
}

func getThisValue(value any, valueType string) int64 {
	var thisValue int64

	if valueType == "String" {
		thisValue = int64(len(value.(string)))
	}

	if valueType == "Int64" {
		thisValue = value.(int64)
	}

	return thisValue
}

func (v *Validation) getOtherValue(valueFrom string, otherKey string) string {
	otherValue := ""

	if valueFrom == "Key" {
		otherValue = v.KeyString(otherKey)
	}

	if valueFrom == "Post" {
		otherValue = v.PostString(otherKey)
	}

	return otherValue
}

func isInArr(thisValue int64, ruleVal []string) bool {
	inFlag := false

	for j := 0; j < len(ruleVal); j++ {
		if thisValue == str2Int64(ruleVal[j]) {
			inFlag = true
			break
		}
	}

	return inFlag
}

func str2Int64(str string) int64 {
	dataNew, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return dataNew
}
