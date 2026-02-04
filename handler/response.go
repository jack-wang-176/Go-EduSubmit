package handler

import (
	"errors"
	"homework_submit/pkg"
	"log"
	"net/http"

	"github.com/jack-wang-176/Maple/web"
)

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func SendResponse(c *web.Context, data any, errs error) {
	if errs != nil {
		err := c.Json(http.StatusOK, ResponseData{
			Code: 0,
			Msg:  "success",
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	var e *pkg.CollectError
	if errors.As(errs, &e) {
		if e.Raw != nil {
			log.Fatal(e.Raw)
		}

		err := c.Json(e.Status, ResponseData{
			Code: e.Code,
			Msg:  e.Msg,
			Data: nil,
		})
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err := c.Json(http.StatusInternalServerError, ResponseData{
		Code: 10000,
		Msg:  "Unknown Server Error",
		Data: nil,
	})
	if err != nil {
		log.Fatal(err)
	}
}
