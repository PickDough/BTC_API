package middleware

import (
	"bytes"
	"io"
	"net/http"
)

type ResponseWriter struct {
	W    http.ResponseWriter
	Buff bytes.Buffer
	Code int
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.W.Header()
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.Code = statusCode
}

func (rw *ResponseWriter) Write(data []byte) (int, error) {
	return rw.Buff.Write(data)
}

func (rw *ResponseWriter) Done() (int64, error) {
	if rw.Code > 0 {
		rw.W.WriteHeader(rw.Code)
	}
	return io.Copy(rw.W, &rw.Buff)
}
