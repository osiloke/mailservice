package mailer

import (
	"github.com/jaytaylor/html2text"
	mailgun "github.com/mailgun/mailgun-go"
	"github.com/mgutz/logxi/v1"
)

// Mailgun mailgun sender
type Mailgun struct {
	Domain       string `json:"domain"`
	APIKey       string `json:"key"`
	PublicAPIKey string `json:"public"`
}

func (m *Mailgun) sendMail(mm *Mail) error {
	log.Info("sending mailgun mail")
	gun := mailgun.NewMailgun(m.Domain, m.APIKey, m.PublicAPIKey)
	body := mm.BuildMessage()
	text, err := html2text.FromString(body, html2text.Options{PrettyTables: true})
	if err != nil {
		log.Info("unable to generate text")
		return err
	}
	for _, recipient := range mm.To {
		body := mm.BuildMessage()
		mmail := mailgun.NewMessage(
			mm.Sender,
			mm.Subject,
			text,
			recipient,
		)
		mmail.SetHtml(body)
		mmail.SetTracking(true)
		_, _, err = gun.Send(mmail)
		if err != nil {
			log.Warn(err.Error())
			//cach UnexpectedResponseError
			if er, ok := err.(*mailgun.UnexpectedResponseError); ok {
				if er.Actual != 401 {
					log.Warn(err.Error())
					return err
				}
			}
		}
	}
	return nil
}
