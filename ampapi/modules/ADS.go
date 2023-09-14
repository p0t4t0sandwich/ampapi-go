package modules

import (
	"ampapi/ampapi"
	"ampapi/ampapi/apimodules"
	"encoding/json"
)

// Struct ADS
type ADS struct {
	CommonAPI
	ADSModule *apimodules.ADSModule
}

// Function NewADS
// Creates a new ADS object
func NewADS(baseUri string, optional ...string) *ADS {
	var ads = new(ADS)

	ads.CommonAPI = *NewCommonAPI(baseUri, optional...)

	ads.ADSModule = apimodules.NewADSModule(&ads.AMPAPI)

	if ads.AMPAPI.Username != "" && (ads.AMPAPI.Password != "" || ads.AMPAPI.RememberMeToken != "") {
		ads.Login()
	}

	return ads
}

// Simplified login function
func (ads *ADS) Login() ampapi.LoginResult {
	var loginResult ampapi.LoginResult = ads.CommonAPI.Login()

	if loginResult.Success {
		ads.AMPAPI.SessionId = loginResult.SessionId
		ads.AMPAPI.RememberMeToken = loginResult.RememberMeToken

		// Update the session ID and remember me token of submodules
		ads.ADSModule.SessionId = loginResult.SessionId
		ads.ADSModule.RememberMeToken = loginResult.RememberMeToken
	}

	return loginResult
}

func (ads *ADS) InstanceLogin(instanceId ampapi.UUID, module string) interface{} {
	var args = make(map[string]any)
	args["username"] = ads.Username
	args["password"] = ads.Password
	args["token"] = ""
	args["rememberMe"] = true

	var loginResult ampapi.LoginResult
	json.Unmarshal(ads.ApiCall("ADSModule/Servers/"+instanceId.String()+"/API/Core/Login", args), &loginResult)

	if loginResult.Success {
		// Prepare the parameters for the instance
		newBaseUri := ads.CommonAPI.BaseURI + "API/ADSModule/Servers/" + instanceId.String()
		rememberMeToken := loginResult.RememberMeToken
		sessionId := loginResult.SessionId

		// Return the correct module
		switch module {
		case "ADS":
			return NewADS(newBaseUri, ads.CommonAPI.Username, "", rememberMeToken, sessionId)
		case "GenericModule":
			return NewGenericModule(newBaseUri, ads.CommonAPI.Username, "", rememberMeToken, sessionId)
		case "Minecraft":
			return NewMinecraft(newBaseUri, ads.CommonAPI.Username, "", rememberMeToken, sessionId)
		default:
			return NewCommonAPI(newBaseUri, ads.CommonAPI.Username, "", rememberMeToken, sessionId)
		}
	} else {
		return nil
	}
}
