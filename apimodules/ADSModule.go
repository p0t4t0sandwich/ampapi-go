package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"ampapi-go"
	"encoding/json"
)

// struct ADSModule
type ADSModule struct {
	*ampapi.AMPAPI
}

// Function NewADSModule
// Creates a new ADSModule object
func NewADSModule(api *ampapi.AMPAPI) *ADSModule {
	return &ADSModule{api}
}

/* AddDatastore -
 * Name Description Optional
 * param newDatastore ampapi.InstanceDatastore  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) AddDatastore(newDatastore ampapi.InstanceDatastore) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["newDatastore"] = newDatastore
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/AddDatastore", args), &res)
	return res
}

/* ApplyInstanceConfiguration -
 * Name Description Optional
 * param InstanceID ampapi.UUID  False
 * param Args map[string]string  False
 * param RebuildConfiguration bool  True
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) ApplyInstanceConfiguration(InstanceID ampapi.UUID, Args map[string]string, RebuildConfiguration bool) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceID"] = InstanceID
	args["Args"] = Args
	args["RebuildConfiguration"] = RebuildConfiguration
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/ApplyInstanceConfiguration", args), &res)
	return res
}

/* ApplyTemplate -
 * Name Description Optional
 * param InstanceID ampapi.UUID  False
 * param TemplateID int32  False
 * param NewFriendlyName string  True
 * param Secret string  True
 * param RestartIfPreviouslyRunning bool  True
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) ApplyTemplate(InstanceID ampapi.UUID, TemplateID int32, NewFriendlyName string, Secret string, RestartIfPreviouslyRunning bool) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["InstanceID"] = InstanceID
	args["TemplateID"] = TemplateID
	args["NewFriendlyName"] = NewFriendlyName
	args["Secret"] = Secret
	args["RestartIfPreviouslyRunning"] = RestartIfPreviouslyRunning
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/ApplyTemplate", args), &res)
	return res
}

/* AttachADS -
 * Name Description Optional
 * param Friendly string  False
 * param IsHTTPS bool  False
 * param Host string  False
 * param Port int32  False
 * param InstanceID ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) AttachADS(Friendly string, IsHTTPS bool, Host string, Port int32, InstanceID ampapi.UUID) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Friendly"] = Friendly
	args["IsHTTPS"] = IsHTTPS
	args["Host"] = Host
	args["Port"] = Port
	args["InstanceID"] = InstanceID
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/AttachADS", args), &res)
	return res
}

/* CloneTemplate -
 * Name Description Optional
 * param Id int32  False
 * param NewName string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) CloneTemplate(Id int32, NewName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Id"] = Id
	args["NewName"] = NewName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/CloneTemplate", args), &res)
	return res
}

/* ConvertToManaged -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) ConvertToManaged(InstanceName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/ConvertToManaged", args), &res)
	return res
}

/* CreateDeploymentTemplate -
 * Name Description Optional
 * param Name string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) CreateDeploymentTemplate(Name string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Name"] = Name
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/CreateDeploymentTemplate", args), &res)
	return res
}

/* CreateInstance -
 * Name Description Optional
 * param TargetADSInstance ampapi.UUID  False
 * param NewInstanceId ampapi.UUID  False
 * param Module string  False
 * param InstanceName string  False
 * param FriendlyName string  False
 * param IPBinding string  False
 * param PortNumber int32  False
 * param AdminUsername string  False
 * param AdminPassword string  False
 * param ProvisionSettings map[string]string  False
 * param AutoConfigure bool When enabled, all settings other than the Module, Target and FriendlyName are ignored and replaced with automatically generated values. True
 * param PostCreate any  True
 * param StartOnBoot bool  True
 * param DisplayImageSource string  True
 * param TargetDatastore int32  True
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) CreateInstance(TargetADSInstance ampapi.UUID, NewInstanceId ampapi.UUID, Module string, InstanceName string, FriendlyName string, IPBinding string, PortNumber int32, AdminUsername string, AdminPassword string, ProvisionSettings map[string]string, AutoConfigure bool, PostCreate any, StartOnBoot bool, DisplayImageSource string, TargetDatastore int32) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["TargetADSInstance"] = TargetADSInstance
	args["NewInstanceId"] = NewInstanceId
	args["Module"] = Module
	args["InstanceName"] = InstanceName
	args["FriendlyName"] = FriendlyName
	args["IPBinding"] = IPBinding
	args["PortNumber"] = PortNumber
	args["AdminUsername"] = AdminUsername
	args["AdminPassword"] = AdminPassword
	args["ProvisionSettings"] = ProvisionSettings
	args["AutoConfigure"] = AutoConfigure
	args["PostCreate"] = PostCreate
	args["StartOnBoot"] = StartOnBoot
	args["DisplayImageSource"] = DisplayImageSource
	args["TargetDatastore"] = TargetDatastore
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/CreateInstance", args), &res)
	return res
}

/* CreateLocalInstance -
 * Name Description Optional
 * param Instance any  False
 * param PostCreate any  True
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) CreateLocalInstance(Instance any, PostCreate any) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Instance"] = Instance
	args["PostCreate"] = PostCreate
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/CreateLocalInstance", args), &res)
	return res
}

/* DeleteDatastore -
 * Name Description Optional
 * param id int32  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) DeleteDatastore(id int32) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["id"] = id
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/DeleteDatastore", args), &res)
	return res
}

/* DeleteDeploymentTemplate -
 * Name Description Optional
 * param Id int32  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) DeleteDeploymentTemplate(Id int32) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Id"] = Id
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/DeleteDeploymentTemplate", args), &res)
	return res
}

/* DeleteInstance -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *ADSModule) DeleteInstance(InstanceName string) ampapi.Result[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.Result[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/DeleteInstance", args), &res)
	return res
}

/* DeleteInstanceUsers -
 * Name Description Optional
 * param InstanceId ampapi.UUID  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) DeleteInstanceUsers(InstanceId ampapi.UUID) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/DeleteInstanceUsers", args), &res)
	return res
}

/* DeployTemplate - A dictionary of setting nodes and values to create the new instance with. Identical in function to the provisioning arguments in the template itself.
 * Name Description Optional
 * param TemplateID int32 The ID of the template to be deployed, as per the Template Management UI in AMP itself. False
 * param NewUsername string If specified, AMP will create a new user with this name for this instance. Must be unique. If this user already exists, this will be ignored but the new instance will be assigned to this user. True
 * param NewPassword string If 'NewUsername' is specified and the user doesn't already exist, the password that will be assigned to this user. True
 * param NewEmail string If 'NewUsername' is specified and the user doesn't already exist, the email address that will be assigned to this user. True
 * param RequiredTags []string If specified, AMP will only deploy this template to targets that have every single 'tag' specified in their target configuration. You can adjust this via the controller by clicking 'Edit' on the target settings. True
 * param Tag string Unrelated to RequiredTags. This is to uniquely identify this instance to your own systems. It may be something like an order ID or service ID so you can find the associated instance again at a later time. If 'UseTagAsInstanceName' is enabled, then this will also be used as the instance name for the created instance - but it must be unique. True
 * param FriendlyName string A friendly name for this instance. If left blank, AMP will generate one for you. True
 * param Secret string Must be a non-empty strong in order to get a callback on deployment state change. This secret will be passed back to you in the callback so you can verify the request. True
 * param PostCreate any 0: Do nothing, 1: Start instance only, 2: Start instance and update application, 3: Full application startup. True
 * param ExtraProvisionSettings map[string]string A dictionary of setting nodes and values to create the new instance with. Identical in function to the provisioning arguments in the template itself. True
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *ADSModule) DeployTemplate(TemplateID int32, NewUsername string, NewPassword string, NewEmail string, RequiredTags []string, Tag string, FriendlyName string, Secret string, PostCreate any, ExtraProvisionSettings map[string]string) ampapi.Result[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["TemplateID"] = TemplateID
	args["NewUsername"] = NewUsername
	args["NewPassword"] = NewPassword
	args["NewEmail"] = NewEmail
	args["RequiredTags"] = RequiredTags
	args["Tag"] = Tag
	args["FriendlyName"] = FriendlyName
	args["Secret"] = Secret
	args["PostCreate"] = PostCreate
	args["ExtraProvisionSettings"] = ExtraProvisionSettings
	var res ampapi.Result[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/DeployTemplate", args), &res)
	return res
}

/* DetatchTarget -
 * Name Description Optional
 * param Id ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) DetatchTarget(Id ampapi.UUID) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Id"] = Id
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/DetatchTarget", args), &res)
	return res
}

/* ExtractEverywhere -
 * Name Description Optional
 * param SourceArchive string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) ExtractEverywhere(SourceArchive string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["SourceArchive"] = SourceArchive
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/ExtractEverywhere", args), &res)
	return res
}

/* GetApplicationEndpoints -
 * Name Description Optional
 * param instanceId ampapi.UUID  False
 * return ampapi.Result[[]ampapi.EndpointInfo]
 */
func (a *ADSModule) GetApplicationEndpoints(instanceId ampapi.UUID) ampapi.Result[[]ampapi.EndpointInfo] {
	var args = make(map[string]any)
	args["instanceId"] = instanceId
	var res ampapi.Result[[]ampapi.EndpointInfo]
	json.Unmarshal(a.ApiCall("ADSModule/GetApplicationEndpoints", args), &res)
	return res
}

/* GetDatastore -
 * Name Description Optional
 * param id int32  False
 * return ampapi.InstanceDatastore
 */
func (a *ADSModule) GetDatastore(id int32) ampapi.InstanceDatastore {
	var args = make(map[string]any)
	args["id"] = id
	var res ampapi.InstanceDatastore
	json.Unmarshal(a.ApiCall("ADSModule/GetDatastore", args), &res)
	return res
}

/* GetDatastoreInstances -
 * Name Description Optional
 * param datastoreId int32  False
 * return ampapi.Result[map[string]any]
 */
func (a *ADSModule) GetDatastoreInstances(datastoreId int32) ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	args["datastoreId"] = datastoreId
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetDatastoreInstances", args), &res)
	return res
}

/* GetDatastores -
 * Name Description Optional
 * return ampapi.Result[[]ampapi.InstanceDatastore]
 */
func (a *ADSModule) GetDatastores() ampapi.Result[[]ampapi.InstanceDatastore] {
	var args = make(map[string]any)
	var res ampapi.Result[[]ampapi.InstanceDatastore]
	json.Unmarshal(a.ApiCall("ADSModule/GetDatastores", args), &res)
	return res
}

/* GetDeploymentTemplates -
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *ADSModule) GetDeploymentTemplates() ampapi.Result[[]any] {
	var args = make(map[string]any)
	var res ampapi.Result[[]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetDeploymentTemplates", args), &res)
	return res
}

/* GetGroup -
 * Name Description Optional
 * param GroupId ampapi.UUID  False
 * return ampapi.Result[ampapi.IADSInstance]
 */
func (a *ADSModule) GetGroup(GroupId ampapi.UUID) ampapi.Result[ampapi.IADSInstance] {
	var args = make(map[string]any)
	args["GroupId"] = GroupId
	var res ampapi.Result[ampapi.IADSInstance]
	json.Unmarshal(a.ApiCall("ADSModule/GetGroup", args), &res)
	return res
}

/* GetInstance -
 * Name Description Optional
 * param InstanceId ampapi.UUID  False
 * return ampapi.Result[ampapi.Instance]
 */
func (a *ADSModule) GetInstance(InstanceId ampapi.UUID) ampapi.Result[ampapi.Instance] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	var res ampapi.Result[ampapi.Instance]
	json.Unmarshal(a.ApiCall("ADSModule/GetInstance", args), &res)
	return res
}

/* GetInstanceNetworkInfo -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.Result[[]any]
 */
func (a *ADSModule) GetInstanceNetworkInfo(InstanceName string) ampapi.Result[[]any] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.Result[[]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetInstanceNetworkInfo", args), &res)
	return res
}

/* GetInstanceStatuses -
 * Name Description Optional
 * return ampapi.Result[map[string]any]
 */
func (a *ADSModule) GetInstanceStatuses() ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetInstanceStatuses", args), &res)
	return res
}

/* GetInstances -
 * Name Description Optional
 * return ampapi.Result[[]ampapi.IADSInstance]
 */
func (a *ADSModule) GetInstances() ampapi.Result[[]ampapi.IADSInstance] {
	var args = make(map[string]any)
	var res ampapi.Result[[]ampapi.IADSInstance]
	json.Unmarshal(a.ApiCall("ADSModule/GetInstances", args), &res)
	return res
}

/* GetLocalInstances -
 * Name Description Optional
 * return ampapi.Result[map[string]any]
 */
func (a *ADSModule) GetLocalInstances() ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetLocalInstances", args), &res)
	return res
}

/* GetProvisionArguments -
 * Name Description Optional
 * param ModuleName string  False
 * return ampapi.Result[map[string]any]
 */
func (a *ADSModule) GetProvisionArguments(ModuleName string) ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	args["ModuleName"] = ModuleName
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetProvisionArguments", args), &res)
	return res
}

/* GetProvisionFitness -
 * Name Description Optional
 * return map[string]any
 */
func (a *ADSModule) GetProvisionFitness() map[string]any {
	var args = make(map[string]any)
	var res map[string]any
	json.Unmarshal(a.ApiCall("ADSModule/GetProvisionFitness", args), &res)
	return res
}

/* GetSupportedApplications -
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *ADSModule) GetSupportedApplications() ampapi.Result[[]any] {
	var args = make(map[string]any)
	var res ampapi.Result[[]any]
	json.Unmarshal(a.ApiCall("ADSModule/GetSupportedApplications", args), &res)
	return res
}

/* GetTargetInfo -
 * Name Description Optional
 * return ampapi.Result[ampapi.RemoteTargetInfo]
 */
func (a *ADSModule) GetTargetInfo() ampapi.Result[ampapi.RemoteTargetInfo] {
	var args = make(map[string]any)
	var res ampapi.Result[ampapi.RemoteTargetInfo]
	json.Unmarshal(a.ApiCall("ADSModule/GetTargetInfo", args), &res)
	return res
}

/* HandoutInstanceConfigs -
 * Name Description Optional
 * param ForModule string  False
 * param SettingNode string  False
 * param Values []string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) HandoutInstanceConfigs(ForModule string, SettingNode string, Values []string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["ForModule"] = ForModule
	args["SettingNode"] = SettingNode
	args["Values"] = Values
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/HandoutInstanceConfigs", args), &res)
	return res
}

/* ManageInstance -
 * Name Description Optional
 * param InstanceId ampapi.UUID  False
 * return ampapi.ActionResult[string]
 */
func (a *ADSModule) ManageInstance(InstanceId ampapi.UUID) ampapi.ActionResult[string] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	var res ampapi.ActionResult[string]
	json.Unmarshal(a.ApiCall("ADSModule/ManageInstance", args), &res)
	return res
}

/* ModifyCustomFirewallRule -
 * Name Description Optional
 * param instanceId ampapi.UUID  False
 * param PortNumber int32  False
 * param Range int32  False
 * param Protocol any  False
 * param Description string  False
 * param Open bool  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) ModifyCustomFirewallRule(instanceId ampapi.UUID, PortNumber int32, Range int32, Protocol any, Description string, Open bool) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["instanceId"] = instanceId
	args["PortNumber"] = PortNumber
	args["Range"] = Range
	args["Protocol"] = Protocol
	args["Description"] = Description
	args["Open"] = Open
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/ModifyCustomFirewallRule", args), &res)
	return res
}

/* MoveInstanceDatastore -
 * Name Description Optional
 * param instanceId ampapi.UUID  False
 * param datastoreId int32  False
 * return ampapi.Task[ampapi.RunningTask]
 */
func (a *ADSModule) MoveInstanceDatastore(instanceId ampapi.UUID, datastoreId int32) ampapi.Task[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["instanceId"] = instanceId
	args["datastoreId"] = datastoreId
	var res ampapi.Task[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/MoveInstanceDatastore", args), &res)
	return res
}

/* ReactivateLocalInstances -
 * Name Description Optional
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *ADSModule) ReactivateLocalInstances() ampapi.Result[ampapi.RunningTask] {
	var args = make(map[string]any)
	var res ampapi.Result[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/ReactivateLocalInstances", args), &res)
	return res
}

/* RefreshAppCache -
 * Name Description Optional
 * return any
 */
func (a *ADSModule) RefreshAppCache() any {
	var args = make(map[string]any)
	var res any
	json.Unmarshal(a.ApiCall("ADSModule/RefreshAppCache", args), &res)
	return res
}

/* RefreshGroup -
 * Name Description Optional
 * param GroupId ampapi.UUID  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) RefreshGroup(GroupId ampapi.UUID) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["GroupId"] = GroupId
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/RefreshGroup", args), &res)
	return res
}

/* RefreshInstanceConfig -
 * Name Description Optional
 * param InstanceId string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) RefreshInstanceConfig(InstanceId string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/RefreshInstanceConfig", args), &res)
	return res
}

/* RefreshRemoteConfigStores -
 * Name Description Optional
 * return any
 */
func (a *ADSModule) RefreshRemoteConfigStores() any {
	var args = make(map[string]any)
	var res any
	json.Unmarshal(a.ApiCall("ADSModule/RefreshRemoteConfigStores", args), &res)
	return res
}

/* RegisterTarget -
 * Name Description Optional
 * param controllerUrl string  False
 * param myUrl string  False
 * param username string  False
 * param password string  False
 * param twoFactorToken string  False
 * param friendlyName string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) RegisterTarget(controllerUrl string, myUrl string, username string, password string, twoFactorToken string, friendlyName string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["controllerUrl"] = controllerUrl
	args["myUrl"] = myUrl
	args["username"] = username
	args["password"] = password
	args["twoFactorToken"] = twoFactorToken
	args["friendlyName"] = friendlyName
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/RegisterTarget", args), &res)
	return res
}

/* RepairDatastore -
 * Name Description Optional
 * param id int32  False
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *ADSModule) RepairDatastore(id int32) ampapi.Result[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["id"] = id
	var res ampapi.Result[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/RepairDatastore", args), &res)
	return res
}

/* RequestDatastoreSizeCalculation -
 * Name Description Optional
 * param datastoreId int32  False
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *ADSModule) RequestDatastoreSizeCalculation(datastoreId int32) ampapi.Result[ampapi.RunningTask] {
	var args = make(map[string]any)
	args["datastoreId"] = datastoreId
	var res ampapi.Result[ampapi.RunningTask]
	json.Unmarshal(a.ApiCall("ADSModule/RequestDatastoreSizeCalculation", args), &res)
	return res
}

/* RestartInstance -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) RestartInstance(InstanceName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/RestartInstance", args), &res)
	return res
}

/* Servers -
 * Name Description Optional
 * param id string  False
 * param REQ_RAWJSON string  False
 * return ampapi.Task[map[string]any]
 */
func (a *ADSModule) Servers(id string, REQ_RAWJSON string) ampapi.Task[map[string]any] {
	var args = make(map[string]any)
	args["id"] = id
	args["REQ_RAWJSON"] = REQ_RAWJSON
	var res ampapi.Task[map[string]any]
	json.Unmarshal(a.ApiCall("ADSModule/Servers", args), &res)
	return res
}

/* SetInstanceConfig -
 * Name Description Optional
 * param InstanceName string  False
 * param SettingNode string  False
 * param Value string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) SetInstanceConfig(InstanceName string, SettingNode string, Value string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	args["SettingNode"] = SettingNode
	args["Value"] = Value
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/SetInstanceConfig", args), &res)
	return res
}

/* SetInstanceNetworkInfo -
 * Name Description Optional
 * param InstanceId ampapi.UUID  False
 * param PortMappings map[string]int  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) SetInstanceNetworkInfo(InstanceId ampapi.UUID, PortMappings map[string]int) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	args["PortMappings"] = PortMappings
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/SetInstanceNetworkInfo", args), &res)
	return res
}

/* SetInstanceSuspended -
 * Name Description Optional
 * param InstanceName string  False
 * param Suspended bool  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) SetInstanceSuspended(InstanceName string, Suspended bool) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	args["Suspended"] = Suspended
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/SetInstanceSuspended", args), &res)
	return res
}

/* StartAllInstances -
 * Name Description Optional
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) StartAllInstances() ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/StartAllInstances", args), &res)
	return res
}

/* StartInstance -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) StartInstance(InstanceName string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/StartInstance", args), &res)
	return res
}

/* StopAllInstances -
 * Name Description Optional
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) StopAllInstances() ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/StopAllInstances", args), &res)
	return res
}

/* StopInstance -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) StopInstance(InstanceName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/StopInstance", args), &res)
	return res
}

/* TestADSLoginDetails -
 * Name Description Optional
 * param url string  False
 * param username string  False
 * param password string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) TestADSLoginDetails(url string, username string, password string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["url"] = url
	args["username"] = username
	args["password"] = password
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/TestADSLoginDetails", args), &res)
	return res
}

/* UpdateDatastore -
 * Name Description Optional
 * param updatedDatastore ampapi.InstanceDatastore  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) UpdateDatastore(updatedDatastore ampapi.InstanceDatastore) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["updatedDatastore"] = updatedDatastore
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/UpdateDatastore", args), &res)
	return res
}

/* UpdateDeploymentTemplate -
 * Name Description Optional
 * param templateToUpdate any  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) UpdateDeploymentTemplate(templateToUpdate any) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["templateToUpdate"] = templateToUpdate
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/UpdateDeploymentTemplate", args), &res)
	return res
}

/* UpdateInstanceInfo -
 * Name Description Optional
 * param InstanceId string  False
 * param FriendlyName string  False
 * param Description string  False
 * param StartOnBoot bool  False
 * param Suspended bool  False
 * param ExcludeFromFirewall bool  False
 * param RunInContainer bool  False
 * param ContainerMemory int32  False
 * param MemoryPolicy any  False
 * param ContainerMaxCPU any  False
 * param ContainerImage string  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) UpdateInstanceInfo(InstanceId string, FriendlyName string, Description string, StartOnBoot bool, Suspended bool, ExcludeFromFirewall bool, RunInContainer bool, ContainerMemory int32, MemoryPolicy any, ContainerMaxCPU any, ContainerImage string) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["InstanceId"] = InstanceId
	args["FriendlyName"] = FriendlyName
	args["Description"] = Description
	args["StartOnBoot"] = StartOnBoot
	args["Suspended"] = Suspended
	args["ExcludeFromFirewall"] = ExcludeFromFirewall
	args["RunInContainer"] = RunInContainer
	args["ContainerMemory"] = ContainerMemory
	args["MemoryPolicy"] = MemoryPolicy
	args["ContainerMaxCPU"] = ContainerMaxCPU
	args["ContainerImage"] = ContainerImage
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/UpdateInstanceInfo", args), &res)
	return res
}

/* UpdateTarget -
 * Name Description Optional
 * param TargetID ampapi.UUID  False
 * return any
 */
func (a *ADSModule) UpdateTarget(TargetID ampapi.UUID) any {
	var args = make(map[string]any)
	args["TargetID"] = TargetID
	var res any
	json.Unmarshal(a.ApiCall("ADSModule/UpdateTarget", args), &res)
	return res
}

/* UpdateTargetInfo -
 * Name Description Optional
 * param Id ampapi.UUID  False
 * param FriendlyName string  False
 * param Url ampapi.URL  False
 * param Description string  False
 * param Tags []string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) UpdateTargetInfo(Id ampapi.UUID, FriendlyName string, Url ampapi.URL, Description string, Tags []string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Id"] = Id
	args["FriendlyName"] = FriendlyName
	args["Url"] = Url
	args["Description"] = Description
	args["Tags"] = Tags
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/UpdateTargetInfo", args), &res)
	return res
}

/* UpgradeAllInstances -
 * Name Description Optional
 * param RestartRunning bool  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *ADSModule) UpgradeAllInstances(RestartRunning bool) ampapi.Task[ampapi.ActionResult[any]] {
	var args = make(map[string]any)
	args["RestartRunning"] = RestartRunning
	var res ampapi.Task[ampapi.ActionResult[any]]
	json.Unmarshal(a.ApiCall("ADSModule/UpgradeAllInstances", args), &res)
	return res
}

/* UpgradeInstance -
 * Name Description Optional
 * param InstanceName string  False
 * return ampapi.ActionResult[any]
 */
func (a *ADSModule) UpgradeInstance(InstanceName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["InstanceName"] = InstanceName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("ADSModule/UpgradeInstance", args), &res)
	return res
}
