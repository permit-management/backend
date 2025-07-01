package app

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/permit-management/backend/pkg/errcode"
	"github.com/permit-management/backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Validation is used to validate the request parameters and return the error response if the validation fails.
func Validation(c *gin.Context, param any, response *Response, withDetail bool) error {
	if err := c.ShouldBind(param); err != nil {
		logger.WithTrace(c).Errorf("Data Validation Error: %v", err)
		if withDetail {
			msg := []string{}
			if castedObject, ok := err.(validator.ValidationErrors); ok {
				for _, err := range castedObject {
					switch err.Tag() {
					case "required":
						msg = append(msg, fmt.Sprintf("%s is required", err.Field()))
					case "email":
						msg = append(msg, fmt.Sprintf("%s is not valid email", err.Field()))
					case "gte":
						msg = append(msg, fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param()))
					case "lte":
						msg = append(msg, fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param()))
					}
				}
			}
			if len(msg) == 0 {
				msg = append(msg, err.Error())
			}
			response.ToErrorResponse(errcode.InvalidRequest.WithDetails(msg...))
		} else {
			response.ToErrorResponse(errcode.InvalidRequest)
		}
		return err
	}
	return nil
}
