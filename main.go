package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Package - CCIP file structure
type Package struct {
	XMLName xml.Name `xml:"package"`
	Project struct {
		ID   string `xml:"id,attr"`
		File string `xml:"file,attr"`
	} `xml:"project"`
}

// AddonInfo - AddonInfo info json structure
type AddonInfo struct {
	Name        string `json:"name"`
	Attachments []struct {
		IsDefault bool   `json:"isDefault"`
		URL       string `json:"url"`
	} `json:"attachments"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing provided")
		os.Exit(128)
	}

	input := strings.Join(os.Args[1:], " ")

	var addon string
	var file string

	url, err := url.ParseRequestURI(input)
	if err == nil && url.Scheme == "curseforge" {
		addon = url.Query().Get("addonId")
		file = url.Query().Get("fileId")
	} else {
		pkg := LoadXML(input)
		addon = pkg.Project.ID
		file = pkg.Project.File
	}

	pack, err := GetPackURL(addon, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addonInfo, err := GetAddonInfo(addon)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occured getting addon info, proceeding")
	}

	var path string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		path = "/Applications/PolyMC.app/Contents/MacOS/icons/" + addonInfo.Name
		args = []string{"open", "-a", "PolyMC", "--args", "--import", pack}
	case "freebsd", "linux", "netbsd", "openbsd":
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}

		path = home + "/.local/share/polymc/icons/" + addonInfo.Name
		// Workaround for hacky PolyMC.deb package wrapper
		deb := "/opt/polymc/run.sh"
		if FileExists(deb) {
			args = []string{deb, "--import", pack}
		} else {
			args = []string{"polymc", "--import", pack}
		}
	case "windows":
		executable, err := os.Executable()
		if err != nil {
			fmt.Println(err)
		}

		polymc := filepath.Dir(executable)
		path = polymc + "\\icons\\" + addonInfo.Name
		args = []string{"PolyMC.exe", "--import", pack}
	}

	var attachmentURL string
	for _, attachment := range addonInfo.Attachments {
		if attachment.IsDefault {
			attachmentURL = attachment.URL
			break
		}
	}

	err = DownloadFile(path, attachmentURL)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error occured downloaading icon, proceeding")
	}

	RunCMD(args)
}

// LoadXML - Load XML from disk into variable
func LoadXML(fileName string) Package {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var pkg Package
	xml.Unmarshal(byteValue, &pkg)

	return pkg
}

// GetPackURL - Request the pack url from Curseforge's API
func GetPackURL(addon string, file string) (string, error) {
	resp, err := http.Get("https://addons-ecs.forgesvc.net/api/v2/addon/" + addon + "/file/" + file + "/download-url")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	url, err := url.Parse(string(body))
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

// GetAddonInfo - Get addon info
func GetAddonInfo(addon string) (AddonInfo, error) {
	resp, err := http.Get("https://addons-ecs.forgesvc.net/api/v2/addon/" + addon)
	if err != nil {
		return AddonInfo{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return AddonInfo{}, err
	}

	var addonInfo AddonInfo
	err = json.Unmarshal(body, &addonInfo)
	if err != nil {
		return AddonInfo{}, err
	}

	return addonInfo, nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// RunCMD - Execute CMD with args
func RunCMD(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd.Wait()
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
