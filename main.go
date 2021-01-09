package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing provided")
		os.Exit(128)
	}

	var pack string

	url, err := url.ParseRequestURI(os.Args[1])
	if err == nil && url.Scheme == "curseforge" {
		paths := strings.Split(url.Path, "/")
		pack = GetPackURL(paths[3], paths[5])
	} else {
		pkg := LoadXML(os.Args[1])
		pack = GetPackURL(pkg.Project.ID, pkg.Project.File)
	}

	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", "-a", "MultiMC", "--args", "import", pack}
	case "freebsd", "linux", "netbsd", "openbsd":
		args = []string{"multimc", "--import", pack}
	case "windows":
		args = []string{"MultiMC.exe", "--import", pack}
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

// GetPackURL - Request the download url from Curseforge's API
func GetPackURL(id string, file string) string {
	resp, err := http.Get("https://addons-ecs.forgesvc.net/api/v2/addon/" + id + "/file/" + file + "/download-url")
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

	url, err := url.Parse(string(body))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return url.String()
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
