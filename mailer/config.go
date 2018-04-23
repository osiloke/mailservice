package mailer

// Config server config
type Config struct {
	Mailgun *Mailgun    `json:"mailgun"`
	Smtp    *SMTPServer `jsom:"smtp"`
}
