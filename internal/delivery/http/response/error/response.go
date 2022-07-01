package error

type ServerErrorResponse struct {
	Success   bool            `json:"success"`
	ErrorCode ServerErrorCode `json:"code"`
	Error     string          `json:"error"`
}
