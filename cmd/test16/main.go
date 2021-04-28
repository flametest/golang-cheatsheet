package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Sig string `json:"sig"`
}


func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Result{
			Code:    234,
			Message: "xx",
		})
	}, middle())

	e.Logger.Fatal(e.Start(":8080"))
}

type CustomWriter struct {
	W io.Writer
	H http.Header
	StatusCode int
}

func (c CustomWriter) Write(i []byte) (int, error) {
	t := &Result{}
	err := json.Unmarshal(i, t)
	if err != nil {
		return 0, err
	}
	t.Sig = "Xxx"
	i1, err := json.Marshal(t)
	if err != nil {
		return 0, err
	}
	return c.W.Write(i1)
}

func (c CustomWriter) Header() http.Header {
	return c.H
}

func (c CustomWriter) WriteHeader(statusCode int) {
	c.StatusCode = statusCode
}

func middle() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var log bytes.Buffer
			newWriter := io.MultiWriter(c.Response().Writer, &log)
			newRespWriter := CustomWriter{W: newWriter, H: c.Response().Header()}
			c.Response().Writer = newRespWriter
			//newResponse := echo.NewResponse(
			//	CustomWriter{
			//		W:          newWriter,
			//		H:          c.Response().Header(),
			//		StatusCode: c.Response().Status,
			//	}, c.Echo(),
			//	)
			//c.Reset(c.Request(), newResponse)
			fmt.Println(c.Response())
			f := handlerFunc(c)
			fmt.Println( c.Response())
			//t := &Result{}
			//err1 := json.Unmarshal(log.Bytes(), t)
			//if err1 != nil {
			//	return err1
			//}
			//t.Sig = "xxx"
			//s, err := json.Marshal(t)
			//if err != nil {
			//	return err
			//}

			//enc := json.NewEncoder(echo.NewResponse(CustomWriter{
			//	W:          newWriter1,
			//	H:          nil,
			//	StatusCode: 0,
			//},c.Echo()))
			//err2 := enc.Encode(t)
			//if err2 != nil {
			//	return err2
			//}
			//newWriter1 := bytes.NewBufferString(string(s))
			//fmt.Println(newWriter1)
			//c.Response().Before(func() {
			//	c.SetResponse(echo.NewResponse(
			//		CustomWriter{
			//			W:          newWriter1,
			//			H:          c.Response().Header(),
			//			StatusCode: c.Response().Status,
			//		}, c.Echo(),
			//		))
			//})
			return f
		}
	}
}