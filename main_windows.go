package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func init() {
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, "twitch", registry.WRITE)
	if err != nil {
		fmt.Println(err)
	}

	defer key.Close()

	err = key.SetStringValue("", "URL: Twitch Handler")
	if err != nil {
		fmt.Println(err)
	}

	err = key.SetStringValue("URL Protocol", "")
	if err != nil {
		fmt.Println(err)
	}

	iconKey, _, err := registry.CreateKey(key, "DefaultIcon", registry.WRITE)
	if err != nil {
		fmt.Println(err)
	}

	defer iconKey.Close()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	err = iconKey.SetStringValue("", dir+"\\MultiMC.exe")
	if err != nil {
		fmt.Println(err)
	}

	shellKey, _, err := registry.CreateKey(key, "shell", registry.WRITE)
	if err != nil {
		fmt.Println(err)
	}

	defer shellKey.Close()

	openKey, _, err := registry.CreateKey(shellKey, "open", registry.WRITE)
	if err != nil {
		fmt.Println(err)
	}

	defer openKey.Close()

	commandKey, _, err := registry.CreateKey(openKey, "command", registry.WRITE)
	if err != nil {
		fmt.Println(err)
	}

	defer commandKey.Close()

	executable, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	err = commandKey.SetStringValue("", "\""+executable+"\" \"%1\"")
	if err != nil {
		fmt.Println(err)
	}
}
