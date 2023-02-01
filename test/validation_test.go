package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v "github.com/tangpanqing/gin-validation"
	"github.com/tangpanqing/gin-validation/message"
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

	va.PostString("name", v.LengthEq(4))
	va.PostString("name", v.LengthNe(5))
	va.PostString("name", v.LengthGt(3))
	va.PostString("name", v.LengthGe(4))
	va.PostString("name", v.LengthLt(5))
	va.PostString("name", v.LengthLe(4))
	va.PostString("name", v.LengthBetween(3, 5))
	va.PostString("name", v.LengthIn(4))

	fmt.Println(va.Error(), "==")

	//testRequired(c, "name1")
	//testLengthEq(c, "name", 3)
	//testLengthGt(c, "name", 4)
	//testLengthBetween(c, "name", 5, 6)
}

func testRequired(c *gin.Context, key string) {
	va := v.New(c)
	va.PostString(key, v.Required)

	if va.Error() != fmt.Sprintf(message.Required, key) {
		fmt.Println("testRequired found err with key = " + key)
	}
}

func testLengthEq(c *gin.Context, key string, length int64) {
	va := v.New(c)
	va.PostString(key, v.LengthEq(length))

	if va.Error() != fmt.Sprintf(message.LengthEq, key, length) {
		fmt.Println("testLengthEq found err with key = " + key)
	}
}

func testLengthGt(c *gin.Context, key string, length int64) {
	va := v.New(c)
	va.PostString(key, v.LengthGt(length))

	if va.Error() != fmt.Sprintf(message.LengthGt, key, length) {
		fmt.Println("testLengthGt found err with key = " + key)
	}
}

func testLengthBetween(c *gin.Context, key string, min int64, max int64) {
	va := v.New(c)
	va.PostString(key, v.LengthBetween(min, max))

	if va.Error() != fmt.Sprintf(message.LengthBetween, key, min, max) {
		fmt.Println("testLengthBetween found err with key = " + key)
	}
}
