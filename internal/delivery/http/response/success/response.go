package success

type ServerSuccessResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
