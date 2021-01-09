package main

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	err = os.Chdir(exPath)
	if err != nil {
		panic(err)
	}

	hasAdmin, err := CheckAdmin()
	if err != nil {
		panic(err)
	}

	if hasAdmin {
		err = UpdateRegistry()
		if err != nil {
			panic(err)
		}
	}
}

func UpdateRegistry() error {
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, "curseforge", registry.WRITE)
	if err != nil {
		return err
	}

	defer key.Close()

	err = key.SetStringValue("", "URL:  CurseForge Protocol")
	if err != nil {
		return err
	}

	err = key.SetStringValue("URL Protocol", "")
	if err != nil {
		return err
	}

	iconKey, _, err := registry.CreateKey(key, "DefaultIcon", registry.WRITE)
	if err != nil {
		return err
	}

	defer iconKey.Close()

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = iconKey.SetStringValue("", dir+"\\MultiMC.exe")
	if err != nil {
		return err
	}

	shellKey, _, err := registry.CreateKey(key, "shell", registry.WRITE)
	if err != nil {
		return err
	}

	defer shellKey.Close()

	openKey, _, err := registry.CreateKey(shellKey, "open", registry.WRITE)
	if err != nil {
		return err
	}

	defer openKey.Close()

	commandKey, _, err := registry.CreateKey(openKey, "command", registry.WRITE)
	if err != nil {
		return err
	}

	defer commandKey.Close()

	executable, err := os.Executable()
	if err != nil {
		return err
	}

	err = commandKey.SetStringValue("", "\""+executable+"\" \"%1\"")
	if err != nil {
		return err
	}

	return nil
}

// CheckAdmin determines if we have administrative privileges.
// Ref: https://coolaj86.com/articles/golang-and-windows-and-admins-oh-my/
func CheckAdmin() (bool, error) {
	var sid *windows.SID

	// Although this looks scary, it is directly copied from the
	// official windows documentation. The Go API for this is a
	// direct wrap around the official C++ API.
	// See https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-checktokenmembership
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false, err
	}

	// This appears to cast a null pointer so I'm not sure why this
	// works, but this guy says it does and it Works for Me™:
	// https://github.com/golang/go/issues/28804#issuecomment-438838144
	token := windows.Token(0)

	member, err := token.IsMember(sid)
	if err != nil {
		return false, err
	}

	return member, nil
}
