#!/bin/python3
from __future__ import annotations

import sys

import requests
import json


type_dict = {
    "InstanceDatastore": "ampapi.InstanceDatastore",
    "ActionResult": "ampapi.ActionResult[any]",
    "Int32": "int32",
    "IEnumerable<InstanceDatastore>": "ampapi.Result[[]ampapi.InstanceDatastore]",
    "RunningTask": "ampapi.Result[ampapi.RunningTask]",
    "Task<RunningTask>": "ampapi.Task[ampapi.RunningTask]",
    "IEnumerable<JObject>": "ampapi.Result[map[string]any]",
    "Guid": "ampapi.UUID",
    "IEnumerable<DeploymentTemplate>": "ampapi.Result[[]any]",
    "String": "string",
    "DeploymentTemplate": "any",
    "Boolean": "bool",
    "List<String>": "[]string",
    "PostCreateActions": "any",
    "Dictionary<String, String>": "map[string]string",
    "RemoteTargetInfo": "ampapi.RemoteTargetInfo",
    "IEnumerable<ApplicationSpec>": "ampapi.Result[[]any]",
    "Void": "any",
    "IEnumerable<EndpointInfo>": "ampapi.Result[[]ampapi.EndpointInfo]",
    "IEnumerable<IADSInstance>": "ampapi.Result[[]ampapi.IADSInstance]",
    "JObject": "map[string]any",
    "PortProtocol": "any",
    "ActionResult<String>": "ampapi.ActionResult[string]",
    "IADSInstance": "ampapi.Result[ampapi.IADSInstance]",
    "Uri": "ampapi.URL",
    "IEnumerable<PortUsage>": "ampapi.Result[[]any]",
    "Dictionary<String, Int32>": "map[string]int",
    "LocalAMPInstance": "any",
    "ContainerMemoryPolicy": "any",
    "Single": "any",
    "Int64": "int64",
    "FileChunkData": "any",
    "IEnumerable<BackupManifest>": "ampapi.Result[[]any]",
    "Nullable<DateTime>": "any", # Optional?
    "IEnumerable<IAuditLogEntry>": "ampapi.Result[[]any]",
    "Dictionary<String, IEnumerable<JObject>>": "ampapi.Result[map[string][]map[string]any]", 
    "IDictionary<String, String>": "map[string]string",
    "List<JObject>": "[]map[string]any",
    "String[]": "[]string",
    "Nullable<Boolean>": "bool", # Optional?
    "ScheduleInfo": "any",
    "Int32[]": "[]int32",
    "TimeIntervalTrigger": "any",
    "IEnumerable<WebSessionSummary>": "ampapi.Result[[]any]",
    "IList<IPermissionsTreeNode>": "[]any",
    "WebauthnLoginInfo": "any",
    "IEnumerable<WebauthnCredentialSummary>": "ampapi.Result[[]any]",
    "IEnumerable<RunningTask>": "ampapi.Result[[]ampapi.RunningTask]",
    "ModuleInfo": "ampapi.Result[ampapi.ModuleInfo]",
    "Dictionary<String, Dictionary<String, MethodInfoSummary>>": "map[string]map[string]any",
    "Object": "any",
    "UpdateInfo": "ampapi.Result[ampapi.UpdateInfo]",
    "IEnumerable<ListeningPortSummary>": "ampapi.Result[[]any]",
    "Task<JObject>": "ampapi.Task[map[string]any]",
    "Task<ActionResult<TwoFactorSetupInfo>>": "ampapi.Task[any]",
    "Task<IEnumerable<String>>": "ampapi.Task[[]string]",
    "Task<UserInfo>": "ampapi.Task[ampapi.UserInfo]",
    "Task<IEnumerable<UserInfoSummary>>": "ampapi.Task[[]any]",
    "Task<IEnumerable<UserInfo>>": "ampapi.Task[[]ampapi.UserInfo]",
    "Task<String>": "ampapi.Task[string]",
    "Task<AuthRoleSummary>": "ampapi.Task[any]",
    "Task<IEnumerable<AuthRoleSummary>>": "ampapi.Task[[]any]",
    "Task<IDictionary<Guid, String>>": "ampapi.Task[map[ampapi.UUID]any]",
    "Task<ActionResult>": "ampapi.Task[ampapi.ActionResult[any]]",
    "Task<ActionResult<Guid>>": "ampapi.Task[ampapi.ActionResult[ampapi.UUID]]",

    ## Custom types
    "ampapi.Result[ampapi.Instance]": "ampapi.Result[ampapi.Instance]",
    "ampapi.Result[ampapi.RemoteTargetInfo]": "ampapi.Result[ampapi.RemoteTargetInfo]",
    "ampapi.SettingsSpec": "ampapi.SettingsSpec",
    "ampapi.Status": "ampapi.Status",
    "ampapi.Updates": "ampapi.Updates",
    "ampapi.Result[map[string]string]": "ampapi.Result[map[string]string]",
    "ampapi.LoginResult": "ampapi.LoginResult"
}

custom_types = {
    # API.ADSModule.GetInstance
    "ADSModule.GetInstance": "ampapi.Result[ampapi.Instance]",
    # API.ADSModule.GetTargetInfo
    "ADSModule.GetTargetInfo": "ampapi.Result[ampapi.RemoteTargetInfo]",

    # API.Core.GetSettingsSpec
    "Core.GetSettingsSpec": "ampapi.SettingsSpec",
    # API.Core.GetStatus
    "Core.GetStatus": "ampapi.Status",
    # API.Core.GetUpdates
    "Core.GetUpdates": "ampapi.Updates",
    # API.Core.GetUserList
    "Core.GetUserList": "ampapi.Result[map[string]string]",
    # API.Core.Login
    "Core.Login": "ampapi.LoginResult",
}

def generate_apimodule_method(module: str, method: str, method_spec: dict):
    # Read the template file
    api_module_method_template = ""
    with open("templates/api_module_method.txt", "r") as tf:
        api_module_method_template = tf.read()
        tf.close()

    # Get the method description
    description = ""
    if "Description" in method_spec.keys():
        description = "\n     * " + method_spec["Description"]

    # Get the method parameters
    parameters_docs = ""
    methodParams = method_spec["Parameters"]
    if len(methodParams) > 0:
        parameters_docs += "\n"
    for i in range(len(methodParams)):
        api_module_method_parameter_doc_template = ""
        with open("templates/api_module_method_parameter_doc.txt", "r") as tf:
            api_module_method_parameter_doc_template = tf.read()
            tf.close()

        name = methodParams[i]["Name"]
        type_name = methodParams[i]["TypeName"]

        # Print out the type if it hasn't been added to the type_dict
        if not type_name in type_dict.keys(): print(type_name)

        description = methodParams[i]["Description"]
        optional = methodParams[i]["Optional"]
        if optional == "true": type_name += ", " + optional

        template = api_module_method_parameter_doc_template\
            .replace("%METHOD_PARAMETER_NAME%", name)\
            .replace("%METHOD_PARAMETER_TYPE%", type_dict[type_name])\
            .replace("%METHOD_PARAMETER_DESCRIPTION%", description)\
            .replace("%METHOD_PARAMETER_OPTIONAL%", str(optional))

        parameters_docs += template
    parameters_docs = parameters_docs[:-1]

    # Get the method return type
    return_type = method_spec["ReturnTypeName"]

    # Print out the type if it hasn't been added to the type_dict
    if not return_type in type_dict.keys(): print(return_type)
    return_type = type_dict[return_type]

    # Get the method parameters
    parameters = ""
    for i in range(len(methodParams)):
        name = methodParams[i]["Name"]
        type_name = methodParams[i]["TypeName"]

        # Print out the type if it hasn't been added to the type_dict
        if not type_name in type_dict.keys(): print(type_name)
        parameters += f"{name} {type_dict[type_name]}, "

    parameters = parameters[:-2]

    # Get the parameters for the data map
    map_string = ""
    if len(methodParams) > 0:
        map_string += "\n"
    for i in range(len(methodParams)):
        api_module_method_parameter_map_template = ""
        with open("templates/api_module_method_parameter_map.txt", "r") as tf:
            api_module_method_parameter_map_template = tf.read()
            tf.close()

        name = methodParams[i]["Name"]
        map_string += api_module_method_parameter_map_template.replace("%METHOD_PARAMETER_NAME%", name)
    map_string = map_string[:-1]

    # Replace placeholders
    template = api_module_method_template\
        .replace("%METHOD_DESCRIPTION%", description)\
        .replace("%METHOD_PARAMETER_DOC%", parameters_docs)\
        .replace("%MODULE_NAME%", module)\
        .replace("%METHOD_NAME%", method)\
        .replace("%METHOD_PARAMETERS%", parameters)\
        .replace("%METHOD_RETURN_TYPE%", return_type)\
        .replace("%METHOD_PARAMETER_MAP%", map_string)

    # End result will return a string
    return template

def generate_apimodule(module: str, methods: dict):
    # Read the template file
    api_module_template = ""
    with open("templates/api_module.txt", "r") as tf:
        api_module_template = tf.read()
        tf.close()

    # Create a new file called f{module}.java
    f = open(f"../ampapi/apimodules/{module}.go","w+")
    f.write(api_module_template.replace("%MODULE_NAME%", module))

    for method in methods.keys():
        f.write(generate_apimodule_method(module, method, methods[method]))

    f.close()

def generate_spec(spec: dict):
    for module in spec.keys():
        if module == "CommonCorePlugin": continue
        
        # Special case for Golang, make sure the first letter is capitalized
        methods = spec[module]
        if module[0].islower():
            module = module[0].upper() + module[1:]

        generate_apimodule(module, methods)

def load_custom_types(spec: dict):
    for type_index in custom_types.keys():
        type_module = type_index.split(".")[0]
        type_method = type_index.split(".")[1]

        # Update the return type
        spec[type_module][type_method]["ReturnTypeName"] = custom_types[type_index]

if __name__ == "__main__":
    # Load remote file
    res = requests.get("https://raw.githubusercontent.com/p0t4t0sandwich/ampapi-spec/main/APISpec.json")
    spec = json.loads(res.content)

    # Load custom types
    load_custom_types(spec)

    generate_spec(spec)