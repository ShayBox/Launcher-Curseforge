package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

// Package - XML structure the CurseForge CCIP files
type Package struct {
	XMLName xml.Name `xml:"package"`
	Project struct {
		ID   string `xml:"id,attr"`
		File string `xml:"file,attr"`
	} `xml:"project"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No path provided")
		os.Exit(128)
	}

	pkg := LoadXML(os.Args[1])
	url := GetURL(pkg)

	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", "-a", "MultiMC", "--args", "import", url}
	case "freebsd", "linux", "netbsd", "openbsd":
		args = []string{"multimc", "--import", url}
	case "windows":
		args = []string{"MultiMC.exe", "--import", url}
	}

	LoadMultiMC(args)
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

// GetURL - Request the download url from Twitch's API
func GetURL(pkg Package) string {
	url := "https://addons-ecs.forgesvc.net/api/v2/addon/" + pkg.Project.ID + "/file/" + pkg.Project.File + "/download-url"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(body)
}

// LoadMultiMC - Execute MultiMC with --import (url)
func LoadMultiMC(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd.Wait()
}
