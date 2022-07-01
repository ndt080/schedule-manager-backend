package smtp

type SmtpRequest struct {
	From    string
	To      []string
	Subject string
	Body    string
}
