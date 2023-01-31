package validation

import "github.com/gin-gonic/gin"

func New(c *gin.Context) *Validation {
	o := Validation{
		ginContext: c,
	}
	return &o
}
