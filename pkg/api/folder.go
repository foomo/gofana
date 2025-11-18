package api

type Folder struct {
	UID     string
	Name    string
	Folders []Folder
}
