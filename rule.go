package validation

import (
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
)

//Required 某字段必须存在,不能为空字符串
var Required = required.Required

//Positive 某字段必须是正数
var Positive = positive.Positive

//Same 某字段必须等于另一个字段
func Same(other string) string {
	return "Same:" + other
}

//Different 某字段必须不等于另一个字段
func Different(other string) string {
	return "Different:" + other
}

//Eq 某字段必须等于某值
func Eq(val any) string {
	return eq.Eq(val)
}

//Ne 某字段必须不等于某值
func Ne(val any) string {
	return ne.Ne(val)
}

//Gt 某字段必须大于某值
func Gt(val any) string {
	return gt.Gt(val)
}

//Ge 某字段必须大于等于某值
func Ge(val any) string {
	return ge.Ge(val)
}

//Lt 某字段必须小于某值
func Lt(val any) string {
	return lt.Lt(val)
}

//Le 某字段必须小于等于某值
func Le(val any) string {
	return le.Le(val)
}

//Between 某字段的值必须在两个值之间
func Between(min any, max any) string {
	return between.Between(min, max)
}

//In 某字段的值必须是其中某一个
func In(val ...any) string {
	return in.In(val...)
}

//LengthEq 某字段的长度必须等于某值
func LengthEq(val int64) string {
	return lengthEq.LengthEq(val)
}

//LengthNe 某字段的长度必须等于某值
func LengthNe(val int64) string {
	return lengthNe.LengthNe(val)
}

//LengthGt 某字段的长度必须大于某值
func LengthGt(val int64) string {
	return lengthGt.LengthGt(val)
}

//LengthGe 某字段的长度必须大于等于某值
func LengthGe(val int64) string {
	return lengthGe.LengthGe(val)
}

//LengthLt 某字段的长度必须小于某值
func LengthLt(val int64) string {
	return lengthLt.LengthLt(val)
}

//LengthLe 某字段的长度必须小于等于某值
func LengthLe(val int64) string {
	return lengthLe.LengthLe(val)
}

//LengthBetween 某字段的长度必须在两个值之间
func LengthBetween(min int64, max int64) string {
	return lengthBetween.LengthBetween(min, max)
}

//LengthIn 某字段的长度必须是其中某一个
func LengthIn(val ...any) string {
	return lengthIn.LengthIn(val...)
}
