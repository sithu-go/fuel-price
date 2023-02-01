package utils

import (
	"fmt"
	"fuel-price/pkg/dto"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func msgForTag(fe validator.FieldError) string {
	field := CapitalToUnderScore(fe.Field())
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%v field is required.", field)
	case "required_with":
		fmt.Println(fe.Param(), "ADSSA")
		return fmt.Sprintf("%v field is required.", field)
	case "oneof":
		return fmt.Sprintf("%v field must be one of %v", field, fe.Param())
	case "email":
		return "Invalid email."
	case "gte", "lte":
		return "invalid length"
	default:
		return "invalid payload" // default error
	}
}

func GenerateValidationErrorMessage(err error) string {
	if vErr, ok := err.(validator.ValidationErrors); ok {
		errMsg := ""
		for _, fieldErr := range vErr {
			errMsg += msgForTag(fieldErr)
		}
		return errMsg
	}
	return err.Error()
}
func GenerateGormErrorResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrMsg = err.Error()
	if IsErrNotFound(err) {
		res.ErrCode = 400
		res.HttpStatusCode = http.StatusBadRequest
		return res
	}

	if IsDuplicate(err) {
		fields := strings.Split(err.Error(), ".")
		field := fields[len(fields)-1]
		res.ErrCode = 400
		msg := "existed"
		res.ErrMsg = fmt.Sprintf("%v %s", strings.Trim(field, "'"), msg)
		res.HttpStatusCode = http.StatusBadRequest
		return res
	}
	res.ErrCode = 500
	res.HttpStatusCode = http.StatusInternalServerError
	return res
}

func GenerateValidationErrorResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrMsg = err.Error()
	if IsValidationError(err) {
		res.ErrCode = 422
		res.ErrMsg = GenerateValidationErrorMessage(err)
		res.HttpStatusCode = http.StatusUnprocessableEntity
		return res
	}
	res.ErrCode = 500
	res.HttpStatusCode = http.StatusInternalServerError
	return res
}

func GenerateAuthErrorResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 401
	res.ErrMsg = "permissin denied"
	res.HttpStatusCode = http.StatusUnauthorized
	return res
}

func GenerateTokenExpireErrorResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 403
	res.ErrMsg = "permission denied"
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateWrongOTPResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 401
	res.ErrMsg = "invalid otp responsed"
	res.HttpStatusCode = http.StatusUnauthorized
	return res
}

func GenerateOTPResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 401
	res.ErrMsg = "wrong password"
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateBadRequestResponseWithErrMsg(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 400
	res.ErrMsg = err.Error()
	res.HttpStatusCode = http.StatusBadRequest
	return res
}

func GenerateDisUserResponseWithErrMsg(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 400
	res.ErrMsg = "disabled user"
	res.HttpStatusCode = http.StatusBadRequest
	return res
}

func GenerateBadRequestResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 400
	res.ErrMsg = "invalid request"
	res.HttpStatusCode = http.StatusBadRequest
	return res
}

func GeneratePermissionDeniedResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 401
	res.ErrMsg = "permission denied"
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateServerError(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 500
	res.ErrMsg = err.Error()
	res.HttpStatusCode = http.StatusInternalServerError
	return res
}

func GenerateSuccessResponse(data any) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 0
	res.ErrMsg = "success"
	res.Data = data
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateExpireResponse(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 200
	res.ErrMsg = err.Error()
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateRoleError(data any) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 401
	res.ErrMsg = "role not found"
	res.Data = data
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateServerBusy(err error) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 600
	res.ErrMsg = err.Error()
	res.HttpStatusCode = http.StatusOK
	return res
}

func GenerateAddressKeyInvalid(lang, address string) *dto.Response {
	res := &dto.Response{}
	res.ErrCode = 400
	msg := "invalid address"
	res.ErrMsg = fmt.Sprintf("%s: %v", msg, address)
	res.HttpStatusCode = http.StatusOK
	return res
}
