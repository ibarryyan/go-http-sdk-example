package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	HeaderName = "barry yan"
)

var (
	data map[string]string

	ErrOk         = Err{Code: 200, Msg: "ok"}
	ErrNotAuth    = Err{Code: 401, Msg: "not auth"}
	ErrRequestBad = Err{Code: 400, Msg: "request bad"}
)

func init() {
	data = make(map[string]string)
}

type T struct {
	Key string `json:"key,omitempty"`
	Val string `json:"val,omitempty"`
}

type Err struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func headerInterceptor(c *gin.Context) {
	header := c.Request.Header.Get("name")

	if header != HeaderName {
		c.JSON(http.StatusUnauthorized, ErrNotAuth)
		c.Abort()
		return
	}
	c.Next()
}

func create(c *gin.Context) {
	var t T
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, ErrRequestBad)
		return
	}
	data[t.Key] = t.Val
	c.JSON(http.StatusOK, ErrOk)
	return
}

func get(c *gin.Context) {
	key := c.Param("key")
	val := data[key]
	c.JSON(http.StatusOK, val)
	return
}

func main() {
	r := gin.Default()
	r.Use(headerInterceptor)

	r.POST("/create", create)
	r.GET("/get/:key", get)

	_ = r.Run(":9999")
}
