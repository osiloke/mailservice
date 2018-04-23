package mailer

import (
	"encoding/json"
	"errors"
	// "github.com/mgutz/logxi/v1"
)

// var prefixMap = map[string]string{
// 	"create": "New",
// 	"delete": "Removed",
// 	"update": "Updated",
// }

var defaultServer = &SMTPServer{
	Host: "localhost",
	Port: "1025",
}

// SendMail send an email
func SendMail(config, params, data, trace string) error {
	c := Config{}
	err := json.Unmarshal([]byte(config), &c)
	if err != nil {
		return err
	}
	p := Params{}
	err = json.Unmarshal([]byte(params), &p)
	if err != nil {
		return err
	}
	d := map[string]interface{}{}
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		return err
	}

	m := &Mail{}
	m.Sender = p.Sender
	m.To = p.Receiver.receivers(d)
	if m.To == nil {
		return errors.New("no receivers")
	}
	m.Subject = "New " + d["StoreTitle"].(string)
	if p.hasSubject() {
		if subject, err := p.parseSubjectTemplate(d); err == nil {
			m.Subject = subject
		}
	}
	body, err := p.parseTemplate(d)
	if err != nil {
		return err
	}
	m.Body = body
	if c.Mailgun != nil {
		return c.Mailgun.sendMail(m)
	}
	if c.Smtp != nil {
		return c.Smtp.sendMail(m)
	}
	return defaultServer.sendMail(m)
}
