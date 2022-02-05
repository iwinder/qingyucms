package qy_errors

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	unknownCoder defaultCoder = defaultCoder{1, http.StatusInternalServerError, "An internal server error occurred", ""}
)

type Coder interface {
	HTTPStatus() int
	String() string
	Reference() string
	Code() int
}

type defaultCoder struct {
	C    int
	HTTP int
	Ext  string
	Ref  string
}

func (coder defaultCoder) Code() int {
	return coder.C
}

func (coder defaultCoder) String() string {
	return coder.Ext
}

func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}
	return coder.HTTP
}

func (coder defaultCoder) Reference() string {
	return coder.Ref
}

var codes = map[int]Coder{}
var codeMux = &sync.Mutex{}

func ParseCoder(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codes[v.code]; ok {
			return coder
		}
	}

	return unknownCoder
}

func Register(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by `gitee.com/windcoder/errors` as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()
	codes[coder.Code()] = coder
}

func MustRegister(coder Coder) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by `gitee.com/windcoder/errors` as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}
	codes[coder.Code()] = coder
}
