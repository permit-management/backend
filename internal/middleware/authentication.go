package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/pkg/app"
	"github.com/permit-management/backend/pkg/errcode"
	"github.com/permit-management/backend/pkg/logger"
)

type AuthenticationPayload struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	RoleID   string `json:"role_id"`
	BranchID uint   `json:"branch_id"`
	AppName  string `json:"app_name"`
}

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := app.NewResponse(c)

		customId := c.Request.Header.Get("X-Consumer-Custom-ID")
		if len(customId) == 0 {
			customId = c.Request.Header.Get("X-Consumer-Custom-Id")
		}
		if len(customId) == 0 {
			logger.WithTrace(c).Info("Invalid Custom ID")
			response.ToErrorResponse(errcode.InvalidToken.WithDetails("Invalid Token"))
			c.Abort()
			return
		}

		var payload AuthenticationPayload
		err := json.Unmarshal([]byte(customId), &payload)

		if err != nil {
			logger.WithTrace(c).Info(err)
			response.ToErrorResponse(errcode.InvalidToken.WithDetails("Invalid Token"))
			c.Abort()
			return
		}

		c.Set(app.USER_ID, payload.UserID)
		c.Set(app.USERNAME, payload.Username)
		c.Set(app.FULLNAME, payload.Fullname)
		c.Set(app.EMAIL, payload.Email)
		c.Set(app.ROLE_NAME, payload.Role)
		c.Set(app.ROLE_ID, payload.RoleID)
		c.Set(app.BRANCH_ID, payload.BranchID)

		c.Next()
	}
}
