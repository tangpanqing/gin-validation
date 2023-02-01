package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/tangpanqing/gin-validation/helper"
	"github.com/tangpanqing/gin-validation/rules/between"
	"github.com/tangpanqing/gin-validation/rules/eq"
	"github.com/tangpanqing/gin-validation/rules/ge"
	"github.com/tangpanqing/gin-validation/rules/gt"
	"github.com/tangpanqing/gin-validation/rules/in"
	"github.com/tangpanqing/gin-validation/rules/le"
	"github.com/tangpanqing/gin-validation/rules/lengthBetween"
	"github.com/tangpanqing/gin-validation/rules/lengthEq"
	"github.com/tangpanqing/gin-validation/rules/lengthGe"
	"github.com/tangpanqing/gin-validation/rules/lengthGt"
	"github.com/tangpanqing/gin-validation/rules/lengthIn"
	"github.com/tangpanqing/gin-validation/rules/lengthLe"
	"github.com/tangpanqing/gin-validation/rules/lengthLt"
	"github.com/tangpanqing/gin-validation/rules/lengthNe"
	"github.com/tangpanqing/gin-validation/rules/lt"
	"github.com/tangpanqing/gin-validation/rules/ne"
	"github.com/tangpanqing/gin-validation/rules/positive"
	"github.com/tangpanqing/gin-validation/rules/required"
	"strings"
)

type Validation struct {
	ginContext *gin.Context
	err        string
}

type HandlerFunc func(key string, value any, valueType string, ruleValue string) string

var funcMap = map[string]HandlerFunc{
	"Required": required.VerifyRequired,
	"Positive": positive.VerifyPositive,

	"Eq":      eq.VerifyEq,
	"Ne":      ne.VerifyNe,
	"Gt":      gt.VerifyGt,
	"Ge":      ge.VerifyGe,
	"Lt":      lt.VerifyLt,
	"Le":      le.VerifyLe,
	"Between": between.VerifyBetween,
	"In":      in.VerifyIn,

	"LengthEq":      lengthEq.VerifyLengthEq,
	"LengthNe":      lengthNe.VerifyLengthNe,
	"LengthGt":      lengthGt.VerifyLengthGt,
	"LengthGe":      lengthGe.VerifyLengthGe,
	"LengthLt":      lengthLt.VerifyLengthLt,
	"LengthLe":      lengthLe.VerifyLengthLe,
	"LengthBetween": lengthBetween.VerifyLengthBetween,
	"LengthIn":      lengthIn.VerifyLengthIn,
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
		value = helper.Str2Int64(valueAny.(string))
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
	value := helper.Str2Int64(valueStr)
	v.handleValue(key, value, "Post", "Int64", rules...)
	return value
}

func (v *Validation) PostFloat64(key string, rules ...string) float64 {
	valueStr := v.ginContext.PostForm(key)
	value := helper.Str2Float64(valueStr)
	v.handleValue(key, value, "Post", "Float64", rules...)
	return value
}

func (v *Validation) handleValue(fieldName string, fieldValue any, valueFrom string, fieldValueType string, rules ...string) {
	if len(rules) > 0 {
		for i := 0; i < len(rules); i++ {
			ruleInfo := strings.Split(rules[i], ":")
			ruleName := ruleInfo[0]

			ruleValue := ""
			if len(ruleInfo) > 1 {
				ruleValue = ruleInfo[1]
			}

			funcName := funcMap[ruleName]
			v.err = funcName(fieldName, fieldValue, fieldValueType, ruleValue)

			//if ruleName == "Same" {
			//	if value.(string) != v.getOtherValue(valueFrom, ruleValue) {
			//		v.err = key + "的值必须等于" + ruleInfo[1] + "的值"
			//	}
			//}
			//
			//if ruleName == "Different" {
			//	if value.(string) == v.getOtherValue(valueFrom, ruleValue) {
			//		v.err = key + "的值必须不等于" + ruleValue + "的值"
			//	}
			//}
		}
	}
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
