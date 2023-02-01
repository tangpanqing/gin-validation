package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v "github.com/tangpanqing/gin-validation"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAll(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = &http.Request{
		PostForm: map[string][]string{
			"name":  {"tang"},
			"desc":  {""},
			"age":   {"18"},
			"score": {"90.5"},

			"password":      {"123456"},
			"newPassword":   {"1234567"},
			"reNewPassword": {"1234567"},
		},
	}

	va := v.New(c)
	va.PostString("desc", v.Required)
	va.PostInt64("age", v.Positive)

	va.PostString("name", v.Eq("tang"))
	va.PostInt64("age", v.Eq(18))
	va.PostFloat64("score", v.Eq(90.5))

	va.PostString("name", v.Ne("tang1"))
	va.PostInt64("age", v.Ne(17))
	va.PostFloat64("score", v.Ne(90.4))

	va.PostInt64("age", v.Gt(17))
	va.PostFloat64("score", v.Gt(90.4))

	va.PostInt64("age", v.Ge(18))
	va.PostFloat64("score", v.Ge(90.5))

	va.PostInt64("age", v.Lt(19))
	va.PostFloat64("score", v.Lt(90.6))

	va.PostInt64("age", v.Le(18))
	va.PostFloat64("score", v.Le(90.5))

	va.PostInt64("age", v.Between(17, 19))
	va.PostFloat64("score", v.Between(90.4, 90.6))

	va.PostInt64("name", v.In("tang"))
	va.PostInt64("age", v.In(18))
	va.PostFloat64("score", v.In(90.5))

	va.PostString("name", v.LenEq(4))
	va.PostString("name", v.LenNe(5))
	va.PostString("name", v.LenGt(3))
	va.PostString("name", v.LenGe(4))
	va.PostString("name", v.LenLt(5))
	va.PostString("name", v.LenLe(4))
	va.PostString("name", v.LenBetween(3, 5))
	va.PostString("name", v.LenIn(4))

	va.PostString("newPassword", v.Different("password"))
	va.PostString("newPassword", v.Same("reNewPassword"))

	fmt.Println(va.Error(), "==")
}
