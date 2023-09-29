package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"encoding/json"

	"github.com/p0t4t0sandwich/ampapi-go"
)

// struct RCONPlugin
type RCONPlugin struct {
	*ampapi.AMPAPI
}

// Function NewRCONPlugin
// Creates a new RCONPlugin object
func NewRCONPlugin(api *ampapi.AMPAPI) *RCONPlugin {
	return &RCONPlugin{api}
}

/* Dummy -
 * Name Description Optional
 * return any
 */
func (a *RCONPlugin) Dummy() any {
	var args = make(map[string]any)
	var res any
	json.Unmarshal(a.ApiCall("RCONPlugin/Dummy", args), &res)
	return res
}
