package go_http_sdk

type Err struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

var (
	ErrOk         = Err{Code: 200, Msg: "ok"}
	ErrNotAuth    = Err{Code: 401, Msg: "not auth"}
	ErrRequestBad = Err{Code: 400, Msg: "request bad"}
	ErrInnerErr   = Err{Code: 500, Msg: "inner err"}
)
