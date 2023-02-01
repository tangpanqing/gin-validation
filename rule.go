package validation

import (
	"github.com/tangpanqing/gin-validation/rules/between"
	"github.com/tangpanqing/gin-validation/rules/eq"
	"github.com/tangpanqing/gin-validation/rules/ge"
	"github.com/tangpanqing/gin-validation/rules/gt"
	"github.com/tangpanqing/gin-validation/rules/in"
	"github.com/tangpanqing/gin-validation/rules/le"
	"github.com/tangpanqing/gin-validation/rules/lenBetween"
	"github.com/tangpanqing/gin-validation/rules/lenEq"
	"github.com/tangpanqing/gin-validation/rules/lenGe"
	"github.com/tangpanqing/gin-validation/rules/lenGt"
	"github.com/tangpanqing/gin-validation/rules/lenIn"
	"github.com/tangpanqing/gin-validation/rules/lenLe"
	"github.com/tangpanqing/gin-validation/rules/lenLt"
	"github.com/tangpanqing/gin-validation/rules/lenNe"
	"github.com/tangpanqing/gin-validation/rules/lt"
	"github.com/tangpanqing/gin-validation/rules/ne"
	"github.com/tangpanqing/gin-validation/rules/positive"
	"github.com/tangpanqing/gin-validation/rules/required"
	"github.com/tangpanqing/gin-validation/rules/same"
)

//Required 某字段必须存在,不能为空字符串
var Required = required.Required

//Positive 某字段必须是正数
var Positive = positive.Positive

//Same 某字段必须等于另一个字段
func Same(otherField string) string {
	return same.Same(otherField)
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

//LenEq 某字段的长度必须等于某值
func LenEq(val int64) string {
	return lenEq.LenEq(val)
}

//LenNe 某字段的长度必须等于某值
func LenNe(val int64) string {
	return lenNe.LenNe(val)
}

//LenGt 某字段的长度必须大于某值
func LenGt(val int64) string {
	return lenGt.LenGt(val)
}

//LenGe 某字段的长度必须大于等于某值
func LenGe(val int64) string {
	return lenGe.LenGe(val)
}

//LenLt 某字段的长度必须小于某值
func LenLt(val int64) string {
	return lenLt.LenLt(val)
}

//LenLe 某字段的长度必须小于等于某值
func LenLe(val int64) string {
	return lenLe.LenLe(val)
}

//LenBetween 某字段的长度必须在两个值之间
func LenBetween(min int64, max int64) string {
	return lenBetween.LenBetween(min, max)
}

//LenIn 某字段的长度必须是其中某一个
func LenIn(val ...any) string {
	return lenIn.LenIn(val...)
}
