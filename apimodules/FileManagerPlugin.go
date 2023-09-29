package apimodules

// A Go library for the AMP API
// Author: p0t4t0sandich

import (
	"encoding/json"

	"github.com/p0t4t0sandwich/ampapi-go"
)

// struct FileManagerPlugin
type FileManagerPlugin struct {
	*ampapi.AMPAPI
}

// Function NewFileManagerPlugin
// Creates a new FileManagerPlugin object
func NewFileManagerPlugin(api *ampapi.AMPAPI) *FileManagerPlugin {
	return &FileManagerPlugin{api}
}

/* AppendFileChunk -
 * Name Description Optional
 * param Filename string  False
 * param Data string  False
 * param Delete bool  False
 * return any
 */
func (a *FileManagerPlugin) AppendFileChunk(Filename string, Data string, Delete bool) any {
	var args = make(map[string]any)
	args["Filename"] = Filename
	args["Data"] = Data
	args["Delete"] = Delete
	var res any
	json.Unmarshal(a.ApiCall("FileManagerPlugin/AppendFileChunk", args), &res)
	return res
}

/* CalculateFileMD5Sum -
 * Name Description Optional
 * param FilePath string  False
 * return ampapi.ActionResult[string]
 */
func (a *FileManagerPlugin) CalculateFileMD5Sum(FilePath string) ampapi.ActionResult[string] {
	var args = make(map[string]any)
	args["FilePath"] = FilePath
	var res ampapi.ActionResult[string]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/CalculateFileMD5Sum", args), &res)
	return res
}

/* ChangeExclusion -
 * Name Description Optional
 * param ModifyPath string  False
 * param AsDirectory bool  False
 * param Exclude bool  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) ChangeExclusion(ModifyPath string, AsDirectory bool, Exclude bool) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["ModifyPath"] = ModifyPath
	args["AsDirectory"] = AsDirectory
	args["Exclude"] = Exclude
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/ChangeExclusion", args), &res)
	return res
}

/* CopyFile -
 * Name Description Optional
 * param Origin string  False
 * param TargetDirectory string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) CopyFile(Origin string, TargetDirectory string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Origin"] = Origin
	args["TargetDirectory"] = TargetDirectory
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/CopyFile", args), &res)
	return res
}

/* CreateArchive -
 * Name Description Optional
 * param PathToArchive string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) CreateArchive(PathToArchive string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["PathToArchive"] = PathToArchive
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/CreateArchive", args), &res)
	return res
}

/* CreateDirectory -
 * Name Description Optional
 * param NewPath string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) CreateDirectory(NewPath string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["NewPath"] = NewPath
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/CreateDirectory", args), &res)
	return res
}

/* DownloadFileFromURL -
 * Name Description Optional
 * param Source ampapi.URL  False
 * param TargetDirectory string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) DownloadFileFromURL(Source ampapi.URL, TargetDirectory string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Source"] = Source
	args["TargetDirectory"] = TargetDirectory
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/DownloadFileFromURL", args), &res)
	return res
}

/* Dummy -
 * Name Description Optional
 * return any
 */
func (a *FileManagerPlugin) Dummy() any {
	var args = make(map[string]any)
	var res any
	json.Unmarshal(a.ApiCall("FileManagerPlugin/Dummy", args), &res)
	return res
}

/* EmptyTrash -
 * Name Description Optional
 * param TrashDirectoryName string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) EmptyTrash(TrashDirectoryName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["TrashDirectoryName"] = TrashDirectoryName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/EmptyTrash", args), &res)
	return res
}

/* ExtractArchive -
 * Name Description Optional
 * param ArchivePath string  False
 * param DestinationPath string  True
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) ExtractArchive(ArchivePath string, DestinationPath string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["ArchivePath"] = ArchivePath
	args["DestinationPath"] = DestinationPath
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/ExtractArchive", args), &res)
	return res
}

/* GetDirectoryListing -
 * Name Description Optional
 * param Dir string  False
 * return ampapi.Result[map[string]any]
 */
func (a *FileManagerPlugin) GetDirectoryListing(Dir string) ampapi.Result[map[string]any] {
	var args = make(map[string]any)
	args["Dir"] = Dir
	var res ampapi.Result[map[string]any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/GetDirectoryListing", args), &res)
	return res
}

/* GetFileChunk -
 * Name Description Optional
 * param Filename string  False
 * param Position int64  False
 * param Length int32  False
 * return any
 */
func (a *FileManagerPlugin) GetFileChunk(Filename string, Position int64, Length int32) any {
	var args = make(map[string]any)
	args["Filename"] = Filename
	args["Position"] = Position
	args["Length"] = Length
	var res any
	json.Unmarshal(a.ApiCall("FileManagerPlugin/GetFileChunk", args), &res)
	return res
}

/* ReadFileChunk -
 * Name Description Optional
 * param Filename string  False
 * param Offset int64  False
 * param ChunkSize int64  True
 * return ampapi.ActionResult[string]
 */
func (a *FileManagerPlugin) ReadFileChunk(Filename string, Offset int64, ChunkSize int64) ampapi.ActionResult[string] {
	var args = make(map[string]any)
	args["Filename"] = Filename
	args["Offset"] = Offset
	args["ChunkSize"] = ChunkSize
	var res ampapi.ActionResult[string]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/ReadFileChunk", args), &res)
	return res
}

/* RenameDirectory - The name component of the new directory (not the full path)
 * Name Description Optional
 * param oldDirectory string The full path to the old directory False
 * param NewDirectoryName string The name component of the new directory (not the full path) False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) RenameDirectory(oldDirectory string, NewDirectoryName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["oldDirectory"] = oldDirectory
	args["NewDirectoryName"] = NewDirectoryName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/RenameDirectory", args), &res)
	return res
}

/* RenameFile -
 * Name Description Optional
 * param Filename string  False
 * param NewFilename string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) RenameFile(Filename string, NewFilename string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Filename"] = Filename
	args["NewFilename"] = NewFilename
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/RenameFile", args), &res)
	return res
}

/* TrashDirectory -
 * Name Description Optional
 * param DirectoryName string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) TrashDirectory(DirectoryName string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["DirectoryName"] = DirectoryName
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/TrashDirectory", args), &res)
	return res
}

/* TrashFile -
 * Name Description Optional
 * param Filename string  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) TrashFile(Filename string) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Filename"] = Filename
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/TrashFile", args), &res)
	return res
}

/* WriteFileChunk -
 * Name Description Optional
 * param Filename string  False
 * param Data string  False
 * param Offset int64  False
 * param FinalChunk bool  False
 * return ampapi.ActionResult[any]
 */
func (a *FileManagerPlugin) WriteFileChunk(Filename string, Data string, Offset int64, FinalChunk bool) ampapi.ActionResult[any] {
	var args = make(map[string]any)
	args["Filename"] = Filename
	args["Data"] = Data
	args["Offset"] = Offset
	args["FinalChunk"] = FinalChunk
	var res ampapi.ActionResult[any]
	json.Unmarshal(a.ApiCall("FileManagerPlugin/WriteFileChunk", args), &res)
	return res
}
