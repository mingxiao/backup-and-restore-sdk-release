package azure

import (
	"encoding/json"
	"io/ioutil"
)

type Artifact struct {
	path string
}

func NewArtifact(path string) Artifact {
	return Artifact{path: path}
}

func (a Artifact) Read() (map[string]ContainerBackup, error) {
	filesContents, err := ioutil.ReadFile(a.path)
	if err != nil {
		return nil, err
	}

	var backups = make(map[string]ContainerBackup)
	err = json.Unmarshal(filesContents, &backups)
	if err != nil {
		return nil, err
	}

	return backups, nil
}

func (a Artifact) Write(backups map[string]ContainerBackup) error {
	filesContents, err := json.Marshal(backups)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(a.path, filesContents, 0644)
}
