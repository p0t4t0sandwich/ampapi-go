package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"ampapi-go"
	"encoding/json"
)

// struct Steamcmdplugin
type Steamcmdplugin struct {
	*ampapi.AMPAPI
}

// Function NewSteamcmdplugin
// Creates a new Steamcmdplugin object
func NewSteamcmdplugin(api *ampapi.AMPAPI) *Steamcmdplugin {
	return &Steamcmdplugin{api}
}

/* CancelSteamGuard -
 * Name Description Optional
 * return any
 */
func (a *Steamcmdplugin) CancelSteamGuard() any {
	var args = make(map[string]any)
	var res any
	json.Unmarshal(a.ApiCall("Steamcmdplugin/CancelSteamGuard", args), &res)
	return res
}

/* SteamGuardCode -
 * Name Description Optional
 * param code string  False
 * return any
 */
func (a *Steamcmdplugin) SteamGuardCode(code string) any {
	var args = make(map[string]any)
	args["code"] = code
	var res any
	json.Unmarshal(a.ApiCall("Steamcmdplugin/SteamGuardCode", args), &res)
	return res
}

/* SteamUsernamePassword -
 * Name Description Optional
 * param username string  False
 * param password string  False
 * return any
 */
func (a *Steamcmdplugin) SteamUsernamePassword(username string, password string) any {
	var args = make(map[string]any)
	args["username"] = username
	args["password"] = password
	var res any
	json.Unmarshal(a.ApiCall("Steamcmdplugin/SteamUsernamePassword", args), &res)
	return res
}
