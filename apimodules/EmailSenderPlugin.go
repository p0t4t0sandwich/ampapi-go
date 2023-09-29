package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"ampapi-go"
	"encoding/json"
)

// struct EmailSenderPlugin
type EmailSenderPlugin struct {
	*ampapi.AMPAPI
}

// Function NewEmailSenderPlugin
// Creates a new EmailSenderPlugin object
func NewEmailSenderPlugin(api *ampapi.AMPAPI) *EmailSenderPlugin {
	return &EmailSenderPlugin{api}
}

/* TestSMTPSettings -
 * Name Description Optional
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *EmailSenderPlugin) TestSMTPSettings() ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("EmailSenderPlugin/TestSMTPSettings", args), &res)
	return res
}
