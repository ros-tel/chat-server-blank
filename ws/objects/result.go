package objects

const (
	MSG_TYPE_ERROR MessageType = "error"

	MSG_RESULT_CODE_OK ResultCode = "OK"
)

type (
	ResultCode string

	MessageResult struct {
		Code  ResultCode `json:"code,omitempty"`
		Error string     `json:"error,omitempty"`
	}
)
