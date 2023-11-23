package ampapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Package metadata
const (
	Author  = "Dylan Sperrer"
	Email   = "dylan@neuralnexus.dev"
	Version = "1.0.9"
)

// AMPAPI struct
type AMPAPI struct {
	BaseURI         string
	DataSource      string
	Username        string
	Password        string
	RememberMeToken string
	SessionId       string
	RequestMethod   string
	lastAPICall     int64
	RelogInterval   int64
}

// NewAMPAPI creates a new AMPAPI object
func NewAMPAPI(baseUri string, optional ...string) *AMPAPI {
	var ampapi *AMPAPI = new(AMPAPI)

	ampapi.BaseURI = baseUri
	if string(baseUri[len(baseUri)-1]) != "/" {
		ampapi.BaseURI += "/"
	}
	ampapi.DataSource = ampapi.BaseURI + "API/"

	if len(optional) > 0 {
		ampapi.Username = optional[0]
	} else {
		ampapi.Username = ""
	}
	if len(optional) > 1 {
		ampapi.Password = optional[1]
	} else {
		ampapi.Password = ""
	}
	if len(optional) > 2 {
		ampapi.RememberMeToken = optional[2]
	} else {
		ampapi.RememberMeToken = ""
	}
	if len(optional) > 3 {
		ampapi.SessionId = optional[3]
	} else {
		ampapi.SessionId = ""
	}

	ampapi.RequestMethod = "POST"
	ampapi.lastAPICall = time.Now().UnixMilli()
	ampapi.RelogInterval = 1000 * 60 * 5

	return ampapi
}

// apiCall is a function that takes an endpoint and returns the response
func (ampapi *AMPAPI) ApiCall(endpoint string, args map[string]any) []byte {
	// Check the last API call time, and if it's been more than the relog interval, relog.
	var now int64 = time.Now().UnixMilli()
	if now-ampapi.lastAPICall > ampapi.RelogInterval {
		ampapi.lastAPICall = now
		ampapi.Login()
	} else {
		ampapi.lastAPICall = now
	}
	args["SESSIONID"] = ampapi.SessionId

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(args)

	client := &http.Client{}
	req, err := http.NewRequest(ampapi.RequestMethod, ampapi.DataSource+endpoint, body)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer req.Body.Close()

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "ampapi-go/1.0.8")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return result
}

// Simplified login function
func (ampapi *AMPAPI) Login() LoginResult {
	var args = make(map[string]any)
	args["username"] = ampapi.Username
	args["password"] = ampapi.Password
	args["token"] = ampapi.RememberMeToken
	args["rememberMe"] = true

	// If remember me token is empty, use the password.
	if ampapi.RememberMeToken != "" {
		args["password"] = ampapi.Password
	}

	var loginResult LoginResult
	json.Unmarshal(ampapi.ApiCall("Core/Login", args), &loginResult)

	if loginResult.Success {
		ampapi.SessionId = loginResult.SessionId
		ampapi.RememberMeToken = loginResult.RememberMeToken
	}

	return loginResult
}
