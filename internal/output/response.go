package output

import (
	"reflect"

	apperror "github.com/Scrummyy/scrummyy-api/internal/app_error"
	"github.com/Scrummyy/scrummyy-api/internal/datatypes"

	"github.com/gin-gonic/gin"
)

// Error returns a formatted response with error metadata added to it.
func Error(c *gin.Context, err apperror.AppErrorInterface) {
	r := datatypes.ApiResponseErrorV3{}
	r.Errors = []*datatypes.Error{
		{
			Message: err.Error(),
		},
	}
	c.JSON(err.GetHttpCode(), r)
}

// ValidationError outputs all validation errors passed to it and sets the appropriate http code.
func ValidationError(c *gin.Context, code int, validationErrors []error) {
	r := datatypes.ApiResponseErrorV3{}

	for _, ve := range validationErrors {
		r.Errors = append(r.Errors, &datatypes.Error{
			Message: ve.Error(),
		})
	}

	c.JSON(code, r)
}

// ErrorWithCode returns a formatted response with error metadata added to it.
func ErrorWithCode(c *gin.Context, err apperror.AppErrorInterface, code string) {
	r := datatypes.ApiResponseErrorV3{}
	r.Errors = []*datatypes.Error{
		{
			Code:    code,
			Message: err.Error(),
		},
	}
	c.JSON(err.GetHttpCode(), r)
}

// SuccessMultiple returns a formatted response for multiple items with pagination data.
func SuccessMultiple(c *gin.Context, code int, resp interface{}, pagination *datatypes.Pagination) {
	r := datatypes.ApiResponseSuccessV3{}
	// set data to blank slice by default
	r.Data = []string{}
	if resp != nil {
		// get type
		rt := reflect.TypeOf(resp)
		//nolint //no need to add cases for all the primitive types
		switch rt.Kind() {
		case reflect.Slice:
			rV := reflect.ValueOf(resp)
			if rV.Len() != 0 {
				r.Data = resp
			}
		default:
			r.Data = resp
		}
	}

	if pagination != nil {
		r.TotalCount = pagination.TotalRecords
		r.Total = pagination.TotalRecords
		r.Pagination = pagination
	}

	c.JSON(code, r)
}

// SuccessSingle returns a formatted response for single item.
func SuccessSingle(c *gin.Context, code int, resp interface{}) {
	c.JSON(code, resp)
}

// Status returns a only with status code.
func Status(c *gin.Context, code int) {
	c.Status(code)
}
