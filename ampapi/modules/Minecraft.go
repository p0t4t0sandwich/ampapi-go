package modules

import (
	"ampapi/ampapi"
	"ampapi/ampapi/apimodules"
)

// Struct Minecraft
type Minecraft struct {
	CommonAPI
	MinecraftModule *apimodules.MinecraftModule
}

// Function NewMinecraft
// Creates a new Minecraft object
func NewMinecraft(baseUri string, optional ...string) *Minecraft {
	var minecraft = new(Minecraft)

	minecraft.CommonAPI = *NewCommonAPI(baseUri, optional...)

	minecraft.MinecraftModule = apimodules.NewMinecraftModule(&minecraft.AMPAPI)

	if minecraft.AMPAPI.Username != "" && (minecraft.AMPAPI.Password != "" || minecraft.AMPAPI.RememberMeToken != "") {
		minecraft.Login()
	}

	return minecraft
}

// Simplified login function
func (minecraft *Minecraft) Login() ampapi.LoginResult {
	var loginResult ampapi.LoginResult = minecraft.CommonAPI.Login()

	if loginResult.Success {
		minecraft.AMPAPI.SessionId = loginResult.SessionId
		minecraft.AMPAPI.RememberMeToken = loginResult.RememberMeToken

		// Update the session ID and remember me token of submodules
		minecraft.MinecraftModule.SessionId = loginResult.SessionId
		minecraft.MinecraftModule.RememberMeToken = loginResult.RememberMeToken
	}

	return loginResult
}
