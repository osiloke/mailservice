package mailer

import (
	"bytes"
	"encoding/json"
	"html/template"
	"time"

	"github.com/Masterminds/sprig"
	humanize "github.com/dustin/go-humanize"
	"github.com/jinzhu/now"
	"github.com/osiloke/mail"
	"github.com/stretchr/stew/objects"
)

var fmap template.FuncMap

func init() {
	fmap = sprig.FuncMap()
	fmap["timeago"] = humanize.Time
	fmap["ftime"] = func(val string) time.Time {
		ret, err := now.Parse(val)
		if err != nil {
			return now.BeginningOfDay()
		}
		return ret
	}
	fmap["shortdate"] = func(t time.Time) string {
		return t.Format("2006-01-02")
	}
}

// Config server config
type Config struct {
	Server   string `json:"server"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Receiver receiver config
type Receiver struct {
	List string `json:"list"`
}

func (r *Receiver) receivers(tplData map[string]interface{}) []string {
	objectData := objects.Map(tplData)
	v := objectData.Get(r.List)
	if _v, ok := v.(map[string]interface{}); ok {
		emails := make([]string, len(_v))
		for _ = range _v {
			if email, ok := _v["email"].(string); ok {
				emails = append(emails, email)
			}
		}
		return emails
	}
	return nil
}

// Params params for mail
type Params struct {
	Template string   `json:"template"`
	Sender   string   `json:"sender"`
	Receiver Receiver `json:"receiver"`
}

func (p *Params) parseTemplate(tplData map[string]interface{}) (string, error) {
	t, _ := template.New("EmailTemplate").Funcs(fmap).Delims("{{", "}}").Parse(p.Template)
	buf := new(bytes.Buffer)
	err := t.Execute(buf, tplData)
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

var prefixMap = map[string]string{
	"create": "New",
	"delete": "Removed",
	"update": "Updated",
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
	server := mail.NewSMTPServer(
		c.Server,
		c.Username,
		c.Password,
		"465",
	)

	m := &mail.Mail{}
	m.Sender = p.Sender
	m.To = p.Receiver.receivers(d)
	// m.Cc = []string{"mnp@gmail.com"}
	// m.Bcc = []string{"a69@outlook.com"}
	m.Subject = "New " + d["StoreTitle"].(string)
	body, err := p.parseTemplate(d)
	if err != nil {
		return err
	}
	m.Body = body
	return mail.SendMail(m, server)
}
