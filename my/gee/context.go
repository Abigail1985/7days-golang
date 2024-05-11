package gee

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Path    string
	Method  string
	Params  map[string]string
	Status  int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Method:  r.Method,
		Path:    r.URL.Path,
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) SetStatus(status int) {
	c.Status = status
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) HTML(code int, html string) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	//这个函数的功能是把values标准化成一个符合format格式的string返回
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}
