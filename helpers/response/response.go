package response

import (
	echo "github.com/labstack/echo/v4"
)

type SuccessResp struct {
	SuccessCode int `json:"successCode"`
	Data        interface{} `json:"data"`
}

type ErrorResp struct {
	ErrorCode int `json:"errorCode"`
	Messages  interface{} `json:"data"`
}

func SuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	resp := &SuccessResp{
		SuccessCode: statusCode,
		Data:        data,
	}
	c.Response().WriteHeader(statusCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(statusCode, resp, "  ")
}

func ErrorResponse(c echo.Context, errorCode int, messages interface{}) error {
	resp := &ErrorResp{
		ErrorCode: errorCode,
		Messages:  messages,
	}
	c.Response().WriteHeader(errorCode)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.JSONPretty(errorCode, resp, "  ")
}
