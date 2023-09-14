package ampapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// AMPAPI struct
type AMPAPI struct {
	BaseURI         string
	DataSource      string
	Username        string
	Password        string
	RememberMeToken string
	SessionId       string
}

// NewAMPAPI creates a new AMPAPI object
func NewAMPAPI(baseUri string, optional ...string) *AMPAPI {
	ampapi := &AMPAPI{}

	ampapi.BaseURI = baseUri
	if string(baseUri[len(baseUri)-1]) == "/" {
		ampapi.DataSource = baseUri + "API"
	} else {
		ampapi.DataSource = baseUri + "/API"
	}
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

	if ampapi.Username != "" && (ampapi.Password != "" || ampapi.RememberMeToken != "") {
		ampapi.Login()
	}
	return ampapi
}

// Var for the request method
var RequestMethod = "POST"

// apiCall is a function that takes an endpoint and returns the response
func (ampapi *AMPAPI) ApiCall(endpoint string, args map[string]any) []byte {
	args["SESSIONID"] = ampapi.SessionId

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(args)

	client := &http.Client{}
	req, err := http.NewRequest(RequestMethod, ampapi.DataSource+"/"+endpoint, body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "ampapi-go/1.0.0")
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

// Login is an easy wrapper function for the Core/Login endpoint
func (ampapi *AMPAPI) Login() LoginResult {
	var args = make(map[string]any)
	args["username"] = ampapi.Username
	args["password"] = ampapi.Password
	args["token"] = ampapi.RememberMeToken
	args["rememberMe"] = true

	var loginResult LoginResult
	json.Unmarshal(ampapi.ApiCall("Core/Login", args), &loginResult)

	if loginResult.Success {
		ampapi.SessionId = loginResult.SessionId
		ampapi.RememberMeToken = loginResult.RememberMeToken
	}

	return loginResult
}
