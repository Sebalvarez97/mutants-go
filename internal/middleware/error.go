package middleware

import (
	"github.com/Sebalvarez97/mutants-go/errors"
	"github.com/Sebalvarez97/mutants-go/tools/web"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// HandledError represents an error returned by an ErrorHandler.
type HandledError struct {
	StatusCode int
	Error      interface{}
	Notify     bool
}

//
// Middleware Error Handler in server package
//
func GetCustomErrorHandler() gin.HandlerFunc {
	return customErrorHandler(gin.ErrorTypeAny)
}

func customErrorHandler(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		log.Println("Handle APP error")
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var handledError *HandledError
			switch err.(type) {
			case *errors.ValidationError:
				handledError = &HandledError{
					StatusCode: http.StatusBadRequest,
					Error:      web.NewError(http.StatusBadRequest, err.(errors.CommunicationError).Message),
				}
			case *errors.CommunicationError:
				handledError = &HandledError{
					StatusCode: http.StatusInternalServerError,
					Error:      web.NewError(http.StatusInternalServerError, err.(errors.CommunicationError).Message),
				}
			case *errors.NotFoundError:
				handledError = &HandledError{
					StatusCode: http.StatusNotFound,
					Error:      web.NewError(http.StatusNotFound, err.(errors.CommunicationError).Message),
				}
			case errors.NotFoundErrorWithoutId:
				handledError = &HandledError{
					StatusCode: http.StatusNotFound,
					Error:      web.NewError(http.StatusNotFound, err.(errors.CommunicationError).Message),
				}
			case errors.AuthError:
				handledError = &HandledError{
					StatusCode: http.StatusUnauthorized,
					Error:      web.NewError(http.StatusUnauthorized, err.(errors.CommunicationError).Message),
				}
			default:
				handledError = &HandledError{
					StatusCode: http.StatusInternalServerError,
					Error:      web.NewError(http.StatusInternalServerError, err.(errors.CommunicationError).Message),
				}
			}
			// Put the error into response
			c.IndentedJSON(handledError.StatusCode, handledError)
			c.Abort()
			// or c.AbortWithStatusJSON(handledError.Code, handledError)
			return
		}
	}
}
