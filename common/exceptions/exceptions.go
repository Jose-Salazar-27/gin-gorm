package exceptions

type Exception struct {
	Code int
	Msg  string
}

func Throw(code int, msg string) *Exception {
	return &Exception{Code: code, Msg: msg}
}
