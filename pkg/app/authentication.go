package app

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	USER_ID   = "ID"
	USERNAME  = "USERNAME"
	FULLNAME  = "FULLNAME"
	EMAIL     = "EMAIL"
	BRANCH_ID = "BRANCH_ID"
	ROLE_ID   = "ROLE_ID"
	ROLE_NAME = "ROLE_NAME"
)

// Validation is used to validate the request parameters and return the error response if the validation fails.
func AuthenticationString(c *gin.Context) string {
	return strings.TrimSpace(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
}
