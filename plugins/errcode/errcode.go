package errcode

type StatusCodeError struct {
	ErrCode int
	ErrMsg  map[string]string
}

func NewStatusCodeError(errCode int, errMsg map[string]string) StatusCodeError {
	return StatusCodeError{ErrCode: errCode, ErrMsg: errMsg}
}

func (e StatusCodeError) Message(lang string) string {
	if v, ok := e.ErrMsg[lang]; ok {
		return v
	}
	return e.ErrMsg["en"]
}
