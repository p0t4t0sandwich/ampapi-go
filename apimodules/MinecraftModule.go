package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"encoding/json"

	"github.com/p0t4t0sandwich/ampapi-go"
)

// struct MinecraftModule
type MinecraftModule struct {
	*ampapi.AMPAPI
}

// Function NewMinecraftModule
// Creates a new MinecraftModule object
func NewMinecraftModule(api *ampapi.AMPAPI) *MinecraftModule {
	return &MinecraftModule{api}
}

/* AcceptEULA -
 * Name Description Optional
 * return bool
 */
func (a *MinecraftModule) AcceptEULA() bool {
	var args = make(map[string]any)
	var res bool
	json.Unmarshal(a.ApiCall("MinecraftModule/AcceptEULA", args), &res)
	return res
}

/* AddOPEntry -
 * Name Description Optional
 * param UserOrUUID string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *MinecraftModule) AddOPEntry(UserOrUUID string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["UserOrUUID"] = UserOrUUID
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("MinecraftModule/AddOPEntry", args), &res)
	return res
}

/* AddToWhitelist -
 * Name Description Optional
 * param UserOrUUID string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *MinecraftModule) AddToWhitelist(UserOrUUID string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["UserOrUUID"] = UserOrUUID
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("MinecraftModule/AddToWhitelist", args), &res)
	return res
}

/* BanUserByID -
 * Name Description Optional
 * param ID string  False
 * return any
 */
func (a *MinecraftModule) BanUserByID(ID string) any {
	var args = make(map[string]any)
	args["ID"] = ID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/BanUserByID", args), &res)
	return res
}

/* BukGetCategories -
 * Name Description Optional
 * return map[string]any
 */
func (a *MinecraftModule) BukGetCategories() map[string]any {
	var args = make(map[string]any)
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetCategories", args), &res)
	return res
}

/* BukGetInstallUpdatePlugin -
 * Name Description Optional
 * param pluginId int32  False
 * return ampapi.Task[ampapi.RunningTask]
 */
func (a *MinecraftModule) BukGetInstallUpdatePlugin(pluginId int32) ampapi.Task[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["pluginId"] = pluginId
	var res ampapi.Task[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetInstallUpdatePlugin", args), &res)
	return res
}

/* BukGetInstalledPlugins -
 * Name Description Optional
 * return map[string]any
 */
func (a *MinecraftModule) BukGetInstalledPlugins() map[string]any {
	var args = make(map[string]any)
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetInstalledPlugins", args), &res)
	return res
}

/* BukGetPluginInfo -
 * Name Description Optional
 * param PluginId int32  False
 * return map[string]any
 */
func (a *MinecraftModule) BukGetPluginInfo(PluginId int32) map[string]any {
	var args = make(map[string]any)
	args["PluginId"] = PluginId
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetPluginInfo", args), &res)
	return res
}

/* BukGetPluginsForCategory -
 * Name Description Optional
 * param CategoryId string  False
 * param PageNumber int32  False
 * param PageSize int32  False
 * return map[string]any
 */
func (a *MinecraftModule) BukGetPluginsForCategory(CategoryId string, PageNumber int32, PageSize int32) map[string]any {
	var args = make(map[string]any)
	args["CategoryId"] = CategoryId
	args["PageNumber"] = PageNumber
	args["PageSize"] = PageSize
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetPluginsForCategory", args), &res)
	return res
}

/* BukGetPopularPlugins -
 * Name Description Optional
 * return map[string]any
 */
func (a *MinecraftModule) BukGetPopularPlugins() map[string]any {
	var args = make(map[string]any)
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetPopularPlugins", args), &res)
	return res
}

/* BukGetRemovePlugin -
 * Name Description Optional
 * param PluginId int32  False
 * return any
 */
func (a *MinecraftModule) BukGetRemovePlugin(PluginId int32) any {
	var args = make(map[string]any)
	args["PluginId"] = PluginId
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetRemovePlugin", args), &res)
	return res
}

/* BukGetSearch -
 * Name Description Optional
 * param Query string  False
 * param PageNumber int32  False
 * param PageSize int32  False
 * return map[string]any
 */
func (a *MinecraftModule) BukGetSearch(Query string, PageNumber int32, PageSize int32) map[string]any {
	var args = make(map[string]any)
	args["Query"] = Query
	args["PageNumber"] = PageNumber
	args["PageSize"] = PageSize
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/BukGetSearch", args), &res)
	return res
}

/* ClearInventoryByID -
 * Name Description Optional
 * param ID string  False
 * return any
 */
func (a *MinecraftModule) ClearInventoryByID(ID string) any {
	var args = make(map[string]any)
	args["ID"] = ID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/ClearInventoryByID", args), &res)
	return res
}

/* GetFailureReason -
 * Name Description Optional
 * return string
 */
func (a *MinecraftModule) GetFailureReason() string {
	var args = make(map[string]any)
	var res string
	json.Unmarshal(a.ApiCall("MinecraftModule/GetFailureReason", args), &res)
	return res
}

/* GetHeadByUUID -
 * Name Description Optional
 * param id string  False
 * return string
 */
func (a *MinecraftModule) GetHeadByUUID(id string) string {
	var args = make(map[string]any)
	args["id"] = id
	var res string
	json.Unmarshal(a.ApiCall("MinecraftModule/GetHeadByUUID", args), &res)
	return res
}

/* GetOPWhitelist -
 * Name Description Optional
 * return map[string]any
 */
func (a *MinecraftModule) GetOPWhitelist() map[string]any {
	var args = make(map[string]any)
	var res map[string]any
	json.Unmarshal(a.ApiCall("MinecraftModule/GetOPWhitelist", args), &res)
	return res
}

/* GetWhitelist -
 * Name Description Optional
 * return ampapi.Result[map[string]any]
 */
func (a *MinecraftModule) GetWhitelist() ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("MinecraftModule/GetWhitelist", args), &res)
	return res
}

/* KickUserByID -
 * Name Description Optional
 * param ID string  False
 * return any
 */
func (a *MinecraftModule) KickUserByID(ID string) any {
	var args = make(map[string]any)
	args["ID"] = ID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/KickUserByID", args), &res)
	return res
}

/* KillByID -
 * Name Description Optional
 * param ID string  False
 * return any
 */
func (a *MinecraftModule) KillByID(ID string) any {
	var args = make(map[string]any)
	args["ID"] = ID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/KillByID", args), &res)
	return res
}

/* LoadOPList -
 * Name Description Optional
 * return ampapi.Result[map[string]any]
 */
func (a *MinecraftModule) LoadOPList() ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("MinecraftModule/LoadOPList", args), &res)
	return res
}

/* RemoveOPEntry -
 * Name Description Optional
 * param UserOrUUID string  False
 * return any
 */
func (a *MinecraftModule) RemoveOPEntry(UserOrUUID string) any {
	var args = make(map[string]any)
	args["UserOrUUID"] = UserOrUUID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/RemoveOPEntry", args), &res)
	return res
}

/* RemoveWhitelistEntry -
 * Name Description Optional
 * param UserOrUUID string  False
 * return any
 */
func (a *MinecraftModule) RemoveWhitelistEntry(UserOrUUID string) any {
	var args = make(map[string]any)
	args["UserOrUUID"] = UserOrUUID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/RemoveWhitelistEntry", args), &res)
	return res
}

/* SmiteByID -
 * Name Description Optional
 * param ID string  False
 * return any
 */
func (a *MinecraftModule) SmiteByID(ID string) any {
	var args = make(map[string]any)
	args["ID"] = ID
	var res any
	json.Unmarshal(a.ApiCall("MinecraftModule/SmiteByID", args), &res)
	return res
}
