package mailer

// https://hackernoon.com/golang-sendmail-sending-mail-through-net-smtp-package-5cadbe2670e0
import (
	"crypto/tls"
	// "github.com/jaytaylor/html2text"
	"github.com/mgutz/logxi/v1"
	"net/smtp"
)

// Mail a mail object
type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

// BuildMessage make messsage
func (mail *Mail) BuildMessage() string {
	header := ""
	header += "\r\n" + mail.Body

	return header
}

// SMTPServer smtp server settings
type SMTPServer struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Port      string `json:"port"`
	TLSConfig *tls.Config
}

// ServerName server name
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (s *SMTPServer) sendMail(mm *Mail) error {
	log.Info("sending smtp mail")
	body := mm.BuildMessage()
	subject := mm.Subject
	// mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	// msg := []byte(subject + mime + "\n" + body)
	// text, err := html2text.FromString(body, html2text.Options{PrettyTables: true})
	// if err != nil {
	// 	log.Info("unable to generate text")
	// 	return err
	// }
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	for _, recipient := range mm.To {

		msg := "From: " + mm.Sender + "\r\n" +
			"To: " + recipient + "\r\n" +
			"MIME-Version: 1.0" + "\r\n" +
			"Content-type: text/html" + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body + "\r\n"
		return smtp.SendMail(
			s.ServerName(),
			auth,
			mm.Sender,
			[]string{recipient},
			[]byte(msg),
		)
	}
	return nil
}

// NewSMTPServer create a new SMTPServer
func NewSMTPServer(host, username, password, port string) *SMTPServer {
	return &SMTPServer{host, username, password, port, &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}}
}
