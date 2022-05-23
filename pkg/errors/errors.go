package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	pkgErrors "github.com/pkg/errors"

	"google.golang.org/grpc/codes"
)

type ServerError struct {
	Status   int        `json:"-"`
	Code     string     `json:"code"`
	GRPCCode codes.Code `json:"grpccode"`
	Message  string     `json:"message"`
}

func (e ServerError) Error() string {
	return e.Message
}

func (e ServerError) Map() map[string]interface{} {
	return map[string]interface{}{
		"status":   e.Status,
		"code":     e.Code,
		"grpccode": e.GRPCCode,
		"message":  e.Message,
	}
}

var (
	// 400
	ErrInvalidInput = ServerError{Code: "400001", Message: "One of the request inputs is not valid.", Status: http.StatusBadRequest, GRPCCode: codes.InvalidArgument}
	// 401
	ErrUnauthorized              = ServerError{Code: "401001", Message: http.StatusText(http.StatusUnauthorized), Status: http.StatusUnauthorized, GRPCCode: codes.Unauthenticated}
	ErrInvalidAuthenticationInfo = ServerError{Code: "401001", Message: "The authentication information was not provided in the correct format. Verify the value of Authorization header.", Status: http.StatusUnauthorized, GRPCCode: codes.Unauthenticated}
	// 403
	ErrForbidden                   = ServerError{Code: "403000", Message: http.StatusText(http.StatusForbidden), Status: http.StatusForbidden}
	ErrAccountIsDisabled           = ServerError{Code: "403001", Message: "The specified account is disabled.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrAuthenticationFailed        = ServerError{Code: "403002", Message: "Server failed to authenticate the request. Make sure the value of the Authorization header is formed correctly including the signature.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrNotAllowed                  = ServerError{Code: "403003", Message: "The request is understood, but it has been refused or access is not allowed.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrOtpExpired                  = ServerError{Code: "403004", Message: "OTP is expired.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrOtpRequired                 = ServerError{Code: "403007", Message: "OTP Binding is required.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrOtpAuthorizationRequired    = ServerError{Code: "403008", Message: "Two-factor authorization is required.", Status: http.StatusForbidden, GRPCCode: codes.PermissionDenied}
	ErrUsernameOrPasswordIncorrect = ServerError{Code: "403006", Message: "Username or Password is incorrect.", Status: http.StatusUnauthorized, GRPCCode: codes.Unauthenticated}
	// 404
	ErrResourceNotFound = ServerError{Code: "404001", Message: "The specified resource does not exist.", Status: http.StatusNotFound, GRPCCode: codes.NotFound}
	ErrPageNotFound     = ServerError{Code: "404003", Message: "Page Not Fount.", Status: http.StatusNotFound, GRPCCode: codes.NotFound}
	// 409
	ErrResourceAlreadyExists = ServerError{Code: "409004", Message: "The specified resource already exists.", Status: http.StatusConflict, GRPCCode: codes.AlreadyExists}
	// 500
	ErrInternalError = ServerError{Code: "500001", Message: "The server encountered an internal error. Please retry the request.", Status: http.StatusInternalServerError, GRPCCode: codes.Internal}
)

// Export origin func of errors
var (
	Unwrap = errors.Unwrap
	Is = errors.Is
	As   = errors.As
	Wrap = pkgErrors.Wrap
	New  = pkgErrors.New
	Cause        = pkgErrors.Cause
	WithMessage  = pkgErrors.WithMessage
	WithMessagef = pkgErrors.WithMessagef
	WithStack    = pkgErrors.WithStack
)

// StackTrace returns stack frames
func StackTrace(e error) []string {
	stacktrace := fmt.Sprintf("%+v\n", e)
	output := strings.Split(stacktrace, "\n")
	return output[:len(output)-1]
}
