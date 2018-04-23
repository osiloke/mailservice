package mailer

import (
	"strings"

	"github.com/mgutz/logxi/v1"
	"github.com/stretchr/stew/objects"
)

func arrayFaceToArrayString(a []interface{}) []string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.(string)
	}
	return s
}

// Receiver receiver config
type Receiver struct {
	List     string        `json:"list"`
	FromData []interface{} `json:"fromData"`
	Email    string        `json:"email"`
}

func (r *Receiver) receiversFromData(tplData map[string]interface{}) []string {
	objectData := objects.Map(tplData)
	receivers := []string{}
	for _, path := range r.FromData {
		v := objectData.Get(path.(string))
		if _v, ok := v.(string); ok {
			receivers = append(receivers, _v)
		}
	}
	return receivers
}
func (r *Receiver) receiversFromList(tplData map[string]interface{}) []string {
	objectData := objects.Map(tplData)
	log.Debug("getting receiver list", "receiver", r, "data", tplData)
	v := objectData.Get(r.List)
	if _v, ok := v.([]interface{}); ok {
		emails := []string{}
		for _, e := range _v {
			if element, ok := e.(map[string]interface{}); ok {
				if email, ok := element["email"].(string); ok {
					emails = append(emails, strings.TrimSpace(email))
				}
			}
		}
		log.Debug("retrieved emails", "emails", emails)
		return emails
	}
	return nil
}

func (r *Receiver) receivers(tplData map[string]interface{}) []string {
	rvr := []string{}
	if len(r.List) != 0 {
		for _, v := range r.receiversFromList(tplData) {
			rvr = append(rvr, v)
		}
	}

	if len(r.FromData) != 0 {
		for _, v := range r.receiversFromData(tplData) {
			rvr = append(rvr, v)
		}
	}

	if len(r.Email) != 0 {
		rvr = append(rvr, r.Email)
	}
	return rvr
}
