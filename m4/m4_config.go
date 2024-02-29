package m4

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Bg3Path       string
	ModFolderPath string
	Mods          []Mod
}

type Mod struct {
	Path   string
	Name   string
	Author string
}

func (config Config) SaveConfig() {
	// Allows user to customize config path by providing an environment variable
	configPath := *GetConfigFolderPath()
	configString, marshalError := json.MarshalIndent(config, "", "    ")
	if marshalError != nil {
		fmt.Println(marshalError)
		panic("JSON Marshal failed!")
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		writeError := os.Mkdir(configPath, 0777)
		if writeError != nil {
			fmt.Print(writeError)
			panic("Directory write failed!")
		}
	}

	writeError := os.WriteFile(configPath+"/config.json", configString, 0777)
	if writeError != nil {
		fmt.Print(writeError)
		panic("Config file write failed!")
	}
}

func GetConfigFolderPath() *string {
	configPath := os.Getenv("M4_CONFIG_PATH")
	if configPath == "" {
		home, _ := os.UserHomeDir()
		configPath = home + "/.m4/"
	}
	return &configPath
}
