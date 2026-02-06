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
	Data any    `json:"data"`
}

func SendResponse(c *web.Context, data any, errs error, msgS ...string) {
	if errs == nil {
		message := "success"
		if len(msgS) > 0 {
			message = msgS[0]
		}

		err := c.Json(http.StatusOK, ResponseData{
			Code: 0,
			Msg:  message,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		return
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
