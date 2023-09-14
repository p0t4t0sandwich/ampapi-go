package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"ampapi/ampapi"
	"encoding/json"
)

// struct Core
type Core struct {
    *ampapi.AMPAPI
}

// Function NewCore
// Creates a new Core object
func NewCore(api *ampapi.AMPAPI) *Core {
	return &Core{api}
}

/* AcknowledgeAMPUpdate - 
 * Name Description Optional
 * return any
 */
func (a *Core) AcknowledgeAMPUpdate() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/AcknowledgeAMPUpdate", args), &res)
    return res
}

/* AddEventTrigger - 
 * Name Description Optional
 * param triggerId ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) AddEventTrigger(triggerId ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["triggerId"] = triggerId
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/AddEventTrigger", args), &res)
    return res
}

/* AddIntervalTrigger - 
 * Name Description Optional
 * param months []int32  False
 * param days []int32  False
 * param hours []int32  False
 * param minutes []int32  False
 * param daysOfMonth []int32  False
 * param description string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) AddIntervalTrigger(months []int32, days []int32, hours []int32, minutes []int32, daysOfMonth []int32, description string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["months"] = months
    args["days"] = days
    args["hours"] = hours
    args["minutes"] = minutes
    args["daysOfMonth"] = daysOfMonth
    args["description"] = description
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/AddIntervalTrigger", args), &res)
    return res
}

/* AddTask - 
 * Name Description Optional
 * param TriggerID ampapi.UUID  False
 * param MethodID string  False
 * param ParameterMapping map[string]string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) AddTask(TriggerID ampapi.UUID, MethodID string, ParameterMapping map[string]string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TriggerID"] = TriggerID
    args["MethodID"] = MethodID
    args["ParameterMapping"] = ParameterMapping
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/AddTask", args), &res)
    return res
}

/* AsyncTest - 
     * DEV: Async test method
 * Name Description Optional
 * return ampapi.Task[string]
 */
func (a *Core) AsyncTest() ampapi.Task[string] {
    var args = make(map[string]any)
    var res ampapi.Task[string]
    json.Unmarshal(a.ApiCall("Core/AsyncTest", args), &res)
    return res
}

/* CancelTask - 
 * Name Description Optional
 * param TaskId ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) CancelTask(TaskId ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TaskId"] = TaskId
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/CancelTask", args), &res)
    return res
}

/* ChangeTaskOrder - 
 * Name Description Optional
 * param TriggerID ampapi.UUID  False
 * param TaskID ampapi.UUID  False
 * param NewOrder int32  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) ChangeTaskOrder(TriggerID ampapi.UUID, TaskID ampapi.UUID, NewOrder int32) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TriggerID"] = TriggerID
    args["TaskID"] = TaskID
    args["NewOrder"] = NewOrder
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/ChangeTaskOrder", args), &res)
    return res
}

/* ChangeUserPassword - 
 * Name Description Optional
 * param Username string  False
 * param OldPassword string  False
 * param NewPassword string  False
 * param TwoFactorPIN string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) ChangeUserPassword(Username string, OldPassword string, NewPassword string, TwoFactorPIN string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Username"] = Username
    args["OldPassword"] = OldPassword
    args["NewPassword"] = NewPassword
    args["TwoFactorPIN"] = TwoFactorPIN
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/ChangeUserPassword", args), &res)
    return res
}

/* ConfirmTwoFactorSetup - 
 * Name Description Optional
 * param Username string  False
 * param TwoFactorCode string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) ConfirmTwoFactorSetup(Username string, TwoFactorCode string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Username"] = Username
    args["TwoFactorCode"] = TwoFactorCode
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/ConfirmTwoFactorSetup", args), &res)
    return res
}

/* CreateRole - 
 * Name Description Optional
 * param Name string  False
 * param AsCommonRole bool  True
 * return ampapi.Task[ampapi.ActionResult[ampapi.UUID]]
 */
func (a *Core) CreateRole(Name string, AsCommonRole bool) ampapi.Task[ampapi.ActionResult[ampapi.UUID]] {
    var args = make(map[string]any)
    args["Name"] = Name
    args["AsCommonRole"] = AsCommonRole
    var res ampapi.Task[ampapi.ActionResult[ampapi.UUID]]
    json.Unmarshal(a.ApiCall("Core/CreateRole", args), &res)
    return res
}

/* CreateTestTask - 
     * DEV: Creates a non-ending task with 50% progress for testing purposes
 * Name Description Optional
 * return any
 */
func (a *Core) CreateTestTask() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/CreateTestTask", args), &res)
    return res
}

/* CreateUser - 
 * Name Description Optional
 * param Username string  False
 * return ampapi.Task[ampapi.ActionResult[ampapi.UUID]]
 */
func (a *Core) CreateUser(Username string) ampapi.Task[ampapi.ActionResult[ampapi.UUID]] {
    var args = make(map[string]any)
    args["Username"] = Username
    var res ampapi.Task[ampapi.ActionResult[ampapi.UUID]]
    json.Unmarshal(a.ApiCall("Core/CreateUser", args), &res)
    return res
}

/* CurrentSessionHasPermission - 
 * Name Description Optional
 * param PermissionNode string  False
 * return bool
 */
func (a *Core) CurrentSessionHasPermission(PermissionNode string) bool {
    var args = make(map[string]any)
    args["PermissionNode"] = PermissionNode
    var res bool
    json.Unmarshal(a.ApiCall("Core/CurrentSessionHasPermission", args), &res)
    return res
}

/* DeleteInstanceUsers - 
 * Name Description Optional
 * param InstanceId ampapi.UUID  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) DeleteInstanceUsers(InstanceId ampapi.UUID) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["InstanceId"] = InstanceId
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/DeleteInstanceUsers", args), &res)
    return res
}

/* DeleteRole - 
 * Name Description Optional
 * param RoleId ampapi.UUID  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) DeleteRole(RoleId ampapi.UUID) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["RoleId"] = RoleId
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/DeleteRole", args), &res)
    return res
}

/* DeleteTask - 
 * Name Description Optional
 * param TriggerID ampapi.UUID  False
 * param TaskID ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) DeleteTask(TriggerID ampapi.UUID, TaskID ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TriggerID"] = TriggerID
    args["TaskID"] = TaskID
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/DeleteTask", args), &res)
    return res
}

/* DeleteTrigger - 
 * Name Description Optional
 * param TriggerID ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) DeleteTrigger(TriggerID ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TriggerID"] = TriggerID
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/DeleteTrigger", args), &res)
    return res
}

/* DeleteUser - 
 * Name Description Optional
 * param Username string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) DeleteUser(Username string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Username"] = Username
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/DeleteUser", args), &res)
    return res
}

/* DisableTwoFactor - 
 * Name Description Optional
 * param Password string  False
 * param TwoFactorCode string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) DisableTwoFactor(Password string, TwoFactorCode string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Password"] = Password
    args["TwoFactorCode"] = TwoFactorCode
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/DisableTwoFactor", args), &res)
    return res
}

/* DismissAllTasks - 
 * Name Description Optional
 * return ampapi.ActionResult[any]
 */
func (a *Core) DismissAllTasks() ampapi.ActionResult[any] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/DismissAllTasks", args), &res)
    return res
}

/* DismissTask - 
 * Name Description Optional
 * param TaskId ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) DismissTask(TaskId ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TaskId"] = TaskId
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/DismissTask", args), &res)
    return res
}

/* EditIntervalTrigger - 
 * Name Description Optional
 * param Id ampapi.UUID  False
 * param months []int32  False
 * param days []int32  False
 * param hours []int32  False
 * param minutes []int32  False
 * param daysOfMonth []int32  False
 * param description string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) EditIntervalTrigger(Id ampapi.UUID, months []int32, days []int32, hours []int32, minutes []int32, daysOfMonth []int32, description string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["Id"] = Id
    args["months"] = months
    args["days"] = days
    args["hours"] = hours
    args["minutes"] = minutes
    args["daysOfMonth"] = daysOfMonth
    args["description"] = description
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/EditIntervalTrigger", args), &res)
    return res
}

/* EditTask - 
 * Name Description Optional
 * param TriggerID ampapi.UUID  False
 * param TaskID ampapi.UUID  False
 * param ParameterMapping map[string]string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) EditTask(TriggerID ampapi.UUID, TaskID ampapi.UUID, ParameterMapping map[string]string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["TriggerID"] = TriggerID
    args["TaskID"] = TaskID
    args["ParameterMapping"] = ParameterMapping
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/EditTask", args), &res)
    return res
}

/* EnableTwoFactor - 
 * Name Description Optional
 * param Username string  False
 * param Password string  False
 * return ampapi.Task[any]
 */
func (a *Core) EnableTwoFactor(Username string, Password string) ampapi.Task[any] {
    var args = make(map[string]any)
    args["Username"] = Username
    args["Password"] = Password
    var res ampapi.Task[any]
    json.Unmarshal(a.ApiCall("Core/EnableTwoFactor", args), &res)
    return res
}

/* EndUserSession - 
 * Name Description Optional
 * param Id ampapi.UUID  False
 * return any
 */
func (a *Core) EndUserSession(Id ampapi.UUID) any {
    var args = make(map[string]any)
    args["Id"] = Id
    var res any
    json.Unmarshal(a.ApiCall("Core/EndUserSession", args), &res)
    return res
}

/* GetAMPRolePermissions - 
 * Name Description Optional
 * param RoleId ampapi.UUID  False
 * return ampapi.Task[[]string]
 */
func (a *Core) GetAMPRolePermissions(RoleId ampapi.UUID) ampapi.Task[[]string] {
    var args = make(map[string]any)
    args["RoleId"] = RoleId
    var res ampapi.Task[[]string]
    json.Unmarshal(a.ApiCall("Core/GetAMPRolePermissions", args), &res)
    return res
}

/* GetAMPUserInfo - 
 * Name Description Optional
 * param Username string  False
 * return ampapi.Task[ampapi.UserInfo]
 */
func (a *Core) GetAMPUserInfo(Username string) ampapi.Task[ampapi.UserInfo] {
    var args = make(map[string]any)
    args["Username"] = Username
    var res ampapi.Task[ampapi.UserInfo]
    json.Unmarshal(a.ApiCall("Core/GetAMPUserInfo", args), &res)
    return res
}

/* GetAMPUsersSummary - 
 * Name Description Optional
 * return ampapi.Task[[]any]
 */
func (a *Core) GetAMPUsersSummary() ampapi.Task[[]any] {
    var args = make(map[string]any)
    var res ampapi.Task[[]any]
    json.Unmarshal(a.ApiCall("Core/GetAMPUsersSummary", args), &res)
    return res
}

/* GetAPISpec - 
 * Name Description Optional
 * return map[string]map[string]any
 */
func (a *Core) GetAPISpec() map[string]map[string]any {
    var args = make(map[string]any)
    var res map[string]map[string]any
    json.Unmarshal(a.ApiCall("Core/GetAPISpec", args), &res)
    return res
}

/* GetActiveAMPSessions - 
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *Core) GetActiveAMPSessions() ampapi.Result[[]any] {
    var args = make(map[string]any)
    var res ampapi.Result[[]any]
    json.Unmarshal(a.ApiCall("Core/GetActiveAMPSessions", args), &res)
    return res
}

/* GetAllAMPUserInfo - 
 * Name Description Optional
 * return ampapi.Task[[]ampapi.UserInfo]
 */
func (a *Core) GetAllAMPUserInfo() ampapi.Task[[]ampapi.UserInfo] {
    var args = make(map[string]any)
    var res ampapi.Task[[]ampapi.UserInfo]
    json.Unmarshal(a.ApiCall("Core/GetAllAMPUserInfo", args), &res)
    return res
}

/* GetAuditLogEntries - 
 * Name Description Optional
 * param Before any  False
 * param Count int32  False
 * return ampapi.Result[[]any]
 */
func (a *Core) GetAuditLogEntries(Before any, Count int32) ampapi.Result[[]any] {
    var args = make(map[string]any)
    args["Before"] = Before
    args["Count"] = Count
    var res ampapi.Result[[]any]
    json.Unmarshal(a.ApiCall("Core/GetAuditLogEntries", args), &res)
    return res
}

/* GetConfig - 
 * Name Description Optional
 * param node string  False
 * return map[string]any
 */
func (a *Core) GetConfig(node string) map[string]any {
    var args = make(map[string]any)
    args["node"] = node
    var res map[string]any
    json.Unmarshal(a.ApiCall("Core/GetConfig", args), &res)
    return res
}

/* GetConfigs - 
 * Name Description Optional
 * param nodes []string  False
 * return ampapi.Result[map[string]any]
 */
func (a *Core) GetConfigs(nodes []string) ampapi.Result[map[string]any] {
    var args = make(map[string]any)
    args["nodes"] = nodes
    var res ampapi.Result[map[string]any]
    json.Unmarshal(a.ApiCall("Core/GetConfigs", args), &res)
    return res
}

/* GetDiagnosticsInfo - 
 * Name Description Optional
 * return map[string]string
 */
func (a *Core) GetDiagnosticsInfo() map[string]string {
    var args = make(map[string]any)
    var res map[string]string
    json.Unmarshal(a.ApiCall("Core/GetDiagnosticsInfo", args), &res)
    return res
}

/* GetModuleInfo - 
 * Name Description Optional
 * return ampapi.Result[ampapi.ModuleInfo]
 */
func (a *Core) GetModuleInfo() ampapi.Result[ampapi.ModuleInfo] {
    var args = make(map[string]any)
    var res ampapi.Result[ampapi.ModuleInfo]
    json.Unmarshal(a.ApiCall("Core/GetModuleInfo", args), &res)
    return res
}

/* GetNewGuid - 
 * Name Description Optional
 * return ampapi.UUID
 */
func (a *Core) GetNewGuid() ampapi.UUID {
    var args = make(map[string]any)
    var res ampapi.UUID
    json.Unmarshal(a.ApiCall("Core/GetNewGuid", args), &res)
    return res
}

/* GetPermissionsSpec - 
 * Name Description Optional
 * return []any
 */
func (a *Core) GetPermissionsSpec() []any {
    var args = make(map[string]any)
    var res []any
    json.Unmarshal(a.ApiCall("Core/GetPermissionsSpec", args), &res)
    return res
}

/* GetPortSummaries - 
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *Core) GetPortSummaries() ampapi.Result[[]any] {
    var args = make(map[string]any)
    var res ampapi.Result[[]any]
    json.Unmarshal(a.ApiCall("Core/GetPortSummaries", args), &res)
    return res
}

/* GetProvisionSpec - 
 * Name Description Optional
 * return []map[string]any
 */
func (a *Core) GetProvisionSpec() []map[string]any {
    var args = make(map[string]any)
    var res []map[string]any
    json.Unmarshal(a.ApiCall("Core/GetProvisionSpec", args), &res)
    return res
}

/* GetRemoteLoginToken - 
 * Name Description Optional
 * param Description string  True
 * param IsTemporary bool  True
 * return ampapi.Task[string]
 */
func (a *Core) GetRemoteLoginToken(Description string, IsTemporary bool) ampapi.Task[string] {
    var args = make(map[string]any)
    args["Description"] = Description
    args["IsTemporary"] = IsTemporary
    var res ampapi.Task[string]
    json.Unmarshal(a.ApiCall("Core/GetRemoteLoginToken", args), &res)
    return res
}

/* GetRole - 
 * Name Description Optional
 * param RoleId ampapi.UUID  False
 * return ampapi.Task[any]
 */
func (a *Core) GetRole(RoleId ampapi.UUID) ampapi.Task[any] {
    var args = make(map[string]any)
    args["RoleId"] = RoleId
    var res ampapi.Task[any]
    json.Unmarshal(a.ApiCall("Core/GetRole", args), &res)
    return res
}

/* GetRoleData - 
 * Name Description Optional
 * return ampapi.Task[[]any]
 */
func (a *Core) GetRoleData() ampapi.Task[[]any] {
    var args = make(map[string]any)
    var res ampapi.Task[[]any]
    json.Unmarshal(a.ApiCall("Core/GetRoleData", args), &res)
    return res
}

/* GetRoleIds - 
 * Name Description Optional
 * return ampapi.Task[map[ampapi.UUID]any]
 */
func (a *Core) GetRoleIds() ampapi.Task[map[ampapi.UUID]any] {
    var args = make(map[string]any)
    var res ampapi.Task[map[ampapi.UUID]any]
    json.Unmarshal(a.ApiCall("Core/GetRoleIds", args), &res)
    return res
}

/* GetScheduleData - 
 * Name Description Optional
 * return any
 */
func (a *Core) GetScheduleData() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/GetScheduleData", args), &res)
    return res
}

/* GetSettingValues - 
 * Name Description Optional
 * param SettingNode string  False
 * param WithRefresh bool  True
 * return map[string]string
 */
func (a *Core) GetSettingValues(SettingNode string, WithRefresh bool) map[string]string {
    var args = make(map[string]any)
    args["SettingNode"] = SettingNode
    args["WithRefresh"] = WithRefresh
    var res map[string]string
    json.Unmarshal(a.ApiCall("Core/GetSettingValues", args), &res)
    return res
}

/* GetSettingsSpec - 
 * Name Description Optional
 * return ampapi.SettingsSpec
 */
func (a *Core) GetSettingsSpec() ampapi.SettingsSpec {
    var args = make(map[string]any)
    var res ampapi.SettingsSpec
    json.Unmarshal(a.ApiCall("Core/GetSettingsSpec", args), &res)
    return res
}

/* GetStatus - 
 * Name Description Optional
 * return ampapi.Status
 */
func (a *Core) GetStatus() ampapi.Status {
    var args = make(map[string]any)
    var res ampapi.Status
    json.Unmarshal(a.ApiCall("Core/GetStatus", args), &res)
    return res
}

/* GetTasks - 
 * Name Description Optional
 * return ampapi.Result[[]ampapi.RunningTask]
 */
func (a *Core) GetTasks() ampapi.Result[[]ampapi.RunningTask] {
    var args = make(map[string]any)
    var res ampapi.Result[[]ampapi.RunningTask]
    json.Unmarshal(a.ApiCall("Core/GetTasks", args), &res)
    return res
}

/* GetTimeIntervalTrigger - 
 * Name Description Optional
 * param Id ampapi.UUID  False
 * return any
 */
func (a *Core) GetTimeIntervalTrigger(Id ampapi.UUID) any {
    var args = make(map[string]any)
    args["Id"] = Id
    var res any
    json.Unmarshal(a.ApiCall("Core/GetTimeIntervalTrigger", args), &res)
    return res
}

/* GetUpdateInfo - 
 * Name Description Optional
 * return ampapi.Result[ampapi.UpdateInfo]
 */
func (a *Core) GetUpdateInfo() ampapi.Result[ampapi.UpdateInfo] {
    var args = make(map[string]any)
    var res ampapi.Result[ampapi.UpdateInfo]
    json.Unmarshal(a.ApiCall("Core/GetUpdateInfo", args), &res)
    return res
}

/* GetUpdates - 
     * Gets changes to the server status, in addition to any notifications or console output that have occured since the last time GetUpdates() was called by the current session.
 * Name Description Optional
 * return ampapi.Updates
 */
func (a *Core) GetUpdates() ampapi.Updates {
    var args = make(map[string]any)
    var res ampapi.Updates
    json.Unmarshal(a.ApiCall("Core/GetUpdates", args), &res)
    return res
}

/* GetUserActionsSpec - 
 * Name Description Optional
 * return any
 */
func (a *Core) GetUserActionsSpec() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/GetUserActionsSpec", args), &res)
    return res
}

/* GetUserInfo - 
 * Name Description Optional
 * param UID string  False
 * return map[string]any
 */
func (a *Core) GetUserInfo(UID string) map[string]any {
    var args = make(map[string]any)
    args["UID"] = UID
    var res map[string]any
    json.Unmarshal(a.ApiCall("Core/GetUserInfo", args), &res)
    return res
}

/* GetUserList - 
     * Returns a list of in-application users
 * Name Description Optional
 * return ampapi.Result[map[string]string]
 */
func (a *Core) GetUserList() ampapi.Result[map[string]string] {
    var args = make(map[string]any)
    var res ampapi.Result[map[string]string]
    json.Unmarshal(a.ApiCall("Core/GetUserList", args), &res)
    return res
}

/* GetWebauthnChallenge - 
 * Name Description Optional
 * return ampapi.ActionResult[string]
 */
func (a *Core) GetWebauthnChallenge() ampapi.ActionResult[string] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[string]
    json.Unmarshal(a.ApiCall("Core/GetWebauthnChallenge", args), &res)
    return res
}

/* GetWebauthnCredentialIDs - 
 * Name Description Optional
 * param username string  False
 * return any
 */
func (a *Core) GetWebauthnCredentialIDs(username string) any {
    var args = make(map[string]any)
    args["username"] = username
    var res any
    json.Unmarshal(a.ApiCall("Core/GetWebauthnCredentialIDs", args), &res)
    return res
}

/* GetWebauthnCredentialSummaries - 
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *Core) GetWebauthnCredentialSummaries() ampapi.Result[[]any] {
    var args = make(map[string]any)
    var res ampapi.Result[[]any]
    json.Unmarshal(a.ApiCall("Core/GetWebauthnCredentialSummaries", args), &res)
    return res
}

/* GetWebserverMetrics - 
 * Name Description Optional
 * return any
 */
func (a *Core) GetWebserverMetrics() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/GetWebserverMetrics", args), &res)
    return res
}

/* Kill - 
 * Name Description Optional
 * return any
 */
func (a *Core) Kill() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/Kill", args), &res)
    return res
}

/* Login - 
 * Name Description Optional
 * param username string  False
 * param password string  False
 * param token string  False
 * param rememberMe bool  False
 * return ampapi.LoginResult
 */
func (a *Core) Login(username string, password string, token string, rememberMe bool) ampapi.LoginResult {
    var args = make(map[string]any)
    args["username"] = username
    args["password"] = password
    args["token"] = token
    args["rememberMe"] = rememberMe
    var res ampapi.LoginResult
    json.Unmarshal(a.ApiCall("Core/Login", args), &res)
    return res
}

/* Logout - 
 * Name Description Optional
 * return any
 */
func (a *Core) Logout() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/Logout", args), &res)
    return res
}

/* RefreshSettingValueList - 
 * Name Description Optional
 * param Node string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) RefreshSettingValueList(Node string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["Node"] = Node
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/RefreshSettingValueList", args), &res)
    return res
}

/* RefreshSettingsSourceCache - 
 * Name Description Optional
 * return any
 */
func (a *Core) RefreshSettingsSourceCache() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/RefreshSettingsSourceCache", args), &res)
    return res
}

/* RenameRole - 
 * Name Description Optional
 * param RoleId ampapi.UUID  False
 * param NewName string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) RenameRole(RoleId ampapi.UUID, NewName string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["RoleId"] = RoleId
    args["NewName"] = NewName
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/RenameRole", args), &res)
    return res
}

/* ResetUserPassword - 
 * Name Description Optional
 * param Username string  False
 * param NewPassword string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) ResetUserPassword(Username string, NewPassword string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Username"] = Username
    args["NewPassword"] = NewPassword
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/ResetUserPassword", args), &res)
    return res
}

/* Restart - 
 * Name Description Optional
 * return ampapi.ActionResult[any]
 */
func (a *Core) Restart() ampapi.ActionResult[any] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/Restart", args), &res)
    return res
}

/* RestartAMP - 
 * Name Description Optional
 * return any
 */
func (a *Core) RestartAMP() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/RestartAMP", args), &res)
    return res
}

/* Resume - 
     * Allows the service to be re-started after previously being suspended.
 * Name Description Optional
 * return any
 */
func (a *Core) Resume() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/Resume", args), &res)
    return res
}

/* RevokeWebauthnCredential - 
 * Name Description Optional
 * param ID int32  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) RevokeWebauthnCredential(ID int32) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["ID"] = ID
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/RevokeWebauthnCredential", args), &res)
    return res
}

/* RunEventTriggerImmediately - 
 * Name Description Optional
 * param triggerId ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) RunEventTriggerImmediately(triggerId ampapi.UUID) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["triggerId"] = triggerId
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/RunEventTriggerImmediately", args), &res)
    return res
}

/* SendConsoleMessage - 
 * Name Description Optional
 * param message string  False
 * return any
 */
func (a *Core) SendConsoleMessage(message string) any {
    var args = make(map[string]any)
    args["message"] = message
    var res any
    json.Unmarshal(a.ApiCall("Core/SendConsoleMessage", args), &res)
    return res
}

/* SetAMPRolePermission - 
 * Name Description Optional
 * param RoleId ampapi.UUID  False
 * param PermissionNode string  False
 * param Enabled bool  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) SetAMPRolePermission(RoleId ampapi.UUID, PermissionNode string, Enabled bool) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["RoleId"] = RoleId
    args["PermissionNode"] = PermissionNode
    args["Enabled"] = Enabled
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/SetAMPRolePermission", args), &res)
    return res
}

/* SetAMPUserRoleMembership - 
 * Name Description Optional
 * param UserId ampapi.UUID  False
 * param RoleId ampapi.UUID  False
 * param IsMember bool  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) SetAMPUserRoleMembership(UserId ampapi.UUID, RoleId ampapi.UUID, IsMember bool) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["UserId"] = UserId
    args["RoleId"] = RoleId
    args["IsMember"] = IsMember
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/SetAMPUserRoleMembership", args), &res)
    return res
}

/* SetConfig - 
 * Name Description Optional
 * param node string  False
 * param value string  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) SetConfig(node string, value string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["node"] = node
    args["value"] = value
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/SetConfig", args), &res)
    return res
}

/* SetConfigs - 
 * Name Description Optional
 * param data map[string]string  False
 * return bool
 */
func (a *Core) SetConfigs(data map[string]string) bool {
    var args = make(map[string]any)
    args["data"] = data
    var res bool
    json.Unmarshal(a.ApiCall("Core/SetConfigs", args), &res)
    return res
}

/* SetTriggerEnabled - 
 * Name Description Optional
 * param Id ampapi.UUID  False
 * param Enabled bool  False
 * return ampapi.ActionResult[any]
 */
func (a *Core) SetTriggerEnabled(Id ampapi.UUID, Enabled bool) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["Id"] = Id
    args["Enabled"] = Enabled
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/SetTriggerEnabled", args), &res)
    return res
}

/* Sleep - 
 * Name Description Optional
 * return ampapi.ActionResult[any]
 */
func (a *Core) Sleep() ampapi.ActionResult[any] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/Sleep", args), &res)
    return res
}

/* Start - 
 * Name Description Optional
 * return ampapi.ActionResult[any]
 */
func (a *Core) Start() ampapi.ActionResult[any] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/Start", args), &res)
    return res
}

/* Stop - 
 * Name Description Optional
 * return any
 */
func (a *Core) Stop() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/Stop", args), &res)
    return res
}

/* Suspend - 
     * Prevents the current instance from being started, and stops it if it's currently running.
 * Name Description Optional
 * return any
 */
func (a *Core) Suspend() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/Suspend", args), &res)
    return res
}

/* UpdateAMPInstance - 
 * Name Description Optional
 * return any
 */
func (a *Core) UpdateAMPInstance() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/UpdateAMPInstance", args), &res)
    return res
}

/* UpdateAccountInfo - 
 * Name Description Optional
 * param EmailAddress string  False
 * param TwoFactorPIN string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) UpdateAccountInfo(EmailAddress string, TwoFactorPIN string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["EmailAddress"] = EmailAddress
    args["TwoFactorPIN"] = TwoFactorPIN
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/UpdateAccountInfo", args), &res)
    return res
}

/* UpdateApplication - 
 * Name Description Optional
 * return ampapi.ActionResult[any]
 */
func (a *Core) UpdateApplication() ampapi.ActionResult[any] {
    var args = make(map[string]any)
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/UpdateApplication", args), &res)
    return res
}

/* UpdateUserInfo - 
 * Name Description Optional
 * param Username string  False
 * param Disabled bool  False
 * param PasswordExpires bool  False
 * param CannotChangePassword bool  False
 * param MustChangePassword bool  False
 * param EmailAddress string  True
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *Core) UpdateUserInfo(Username string, Disabled bool, PasswordExpires bool, CannotChangePassword bool, MustChangePassword bool, EmailAddress string) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["Username"] = Username
    args["Disabled"] = Disabled
    args["PasswordExpires"] = PasswordExpires
    args["CannotChangePassword"] = CannotChangePassword
    args["MustChangePassword"] = MustChangePassword
    args["EmailAddress"] = EmailAddress
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("Core/UpdateUserInfo", args), &res)
    return res
}

/* UpgradeAMP - 
 * Name Description Optional
 * return any
 */
func (a *Core) UpgradeAMP() any {
    var args = make(map[string]any)
    var res any
    json.Unmarshal(a.ApiCall("Core/UpgradeAMP", args), &res)
    return res
}

/* WebauthnRegister - 
 * Name Description Optional
 * param attestationObject string  False
 * param clientDataJSON string  False
 * param description string  True
 * return ampapi.ActionResult[any]
 */
func (a *Core) WebauthnRegister(attestationObject string, clientDataJSON string, description string) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["attestationObject"] = attestationObject
    args["clientDataJSON"] = clientDataJSON
    args["description"] = description
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("Core/WebauthnRegister", args), &res)
    return res
}

