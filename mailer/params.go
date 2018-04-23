package mailer

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"html/template"
	textTpl "text/template"
	"time"

	"github.com/Masterminds/sprig"
	humanize "github.com/dustin/go-humanize"
	"github.com/jinzhu/now"
	"github.com/mgutz/logxi/v1"
)

var fmap map[string]interface{}

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

// Params params for mail
type Params struct {
	SubjectTemplate string   `json:"subjectTemplate"`
	Template        string   `json:"template"`
	Sender          string   `json:"sender"`
	Receiver        Receiver `json:"receiver"`
}

func (p *Params) hasSubject() bool {
	return p.SubjectTemplate != ""
}
func (p *Params) parseSubjectTemplate(tplData map[string]interface{}) (string, error) {
	t, err := textTpl.New("SubjectTemplate").Funcs(fmap).Delims("%{", "}%").Parse(p.SubjectTemplate)
	if err != nil {
		log.Warn("cannot parse template", "err", err.Error())
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, tplData)
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func (p *Params) parseTemplate(tplData map[string]interface{}) (string, error) {
	sDec, err := b64.StdEncoding.DecodeString(p.Template)
	if err != nil {
		return "", errors.New("template could not be decoded")
	}
	tplString := string(sDec)
	t, err := template.New("EmailTemplate").Funcs(fmap).Delims("%{", "}%").Parse(tplString)
	if err != nil {
		log.Warn("cannot parse template", "err", err.Error())
		return "", err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, tplData)
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}
