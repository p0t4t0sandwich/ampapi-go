package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"ampapi/ampapi"
	"encoding/json"
)

// struct LocalFileBackupPlugin
type LocalFileBackupPlugin struct {
    *ampapi.AMPAPI
}

// Function NewLocalFileBackupPlugin
// Creates a new LocalFileBackupPlugin object
func NewLocalFileBackupPlugin(api *ampapi.AMPAPI) *LocalFileBackupPlugin {
	return &LocalFileBackupPlugin{api}
}

/* DeleteFromS3 - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * return ampapi.Task[ampapi.ActionResult[any]]
 */
func (a *LocalFileBackupPlugin) DeleteFromS3(BackupId ampapi.UUID) ampapi.Task[ampapi.ActionResult[any]] {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    var res ampapi.Task[ampapi.ActionResult[any]]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/DeleteFromS3", args), &res)
    return res
}

/* DeleteLocalBackup - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * return any
 */
func (a *LocalFileBackupPlugin) DeleteLocalBackup(BackupId ampapi.UUID) any {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    var res any
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/DeleteLocalBackup", args), &res)
    return res
}

/* DownloadFromS3 - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *LocalFileBackupPlugin) DownloadFromS3(BackupId ampapi.UUID) ampapi.Result[ampapi.RunningTask] {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    var res ampapi.Result[ampapi.RunningTask]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/DownloadFromS3", args), &res)
    return res
}

/* GetBackups - 
 * Name Description Optional
 * return ampapi.Result[[]any]
 */
func (a *LocalFileBackupPlugin) GetBackups() ampapi.Result[[]any] {
    var args = make(map[string]any)
    var res ampapi.Result[[]any]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/GetBackups", args), &res)
    return res
}

/* RestoreBackup - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * param DeleteExistingData bool  True
 * return ampapi.ActionResult[any]
 */
func (a *LocalFileBackupPlugin) RestoreBackup(BackupId ampapi.UUID, DeleteExistingData bool) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    args["DeleteExistingData"] = DeleteExistingData
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/RestoreBackup", args), &res)
    return res
}

/* SetBackupSticky - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * param Sticky bool  False
 * return any
 */
func (a *LocalFileBackupPlugin) SetBackupSticky(BackupId ampapi.UUID, Sticky bool) any {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    args["Sticky"] = Sticky
    var res any
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/SetBackupSticky", args), &res)
    return res
}

/* TakeBackup - 
 * Name Description Optional
 * param Title string  False
 * param Description string  False
 * param Sticky bool  False
 * return ampapi.ActionResult[any]
 */
func (a *LocalFileBackupPlugin) TakeBackup(Title string, Description string, Sticky bool) ampapi.ActionResult[any] {
    var args = make(map[string]any)
    args["Title"] = Title
    args["Description"] = Description
    args["Sticky"] = Sticky
    var res ampapi.ActionResult[any]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/TakeBackup", args), &res)
    return res
}

/* UploadToS3 - 
 * Name Description Optional
 * param BackupId ampapi.UUID  False
 * return ampapi.Result[ampapi.RunningTask]
 */
func (a *LocalFileBackupPlugin) UploadToS3(BackupId ampapi.UUID) ampapi.Result[ampapi.RunningTask] {
    var args = make(map[string]any)
    args["BackupId"] = BackupId
    var res ampapi.Result[ampapi.RunningTask]
    json.Unmarshal(a.ApiCall("LocalFileBackupPlugin/UploadToS3", args), &res)
    return res
}

