package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

type config struct {
	AccessToken        string             `json:"AccessToken"`
	OrganizationFields organizationFields `json:"OrganizationFields"`
	SpaceFields        spaceFields        `json:"SpaceFields"`
}

type organizationFields struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

type spaceFields struct {
	Name string `json:"Name"`
}

const (
	templateFlagValue = "{{if .AccessToken}}{{.OrganizationFields.Name}}{{if .SpaceFields.Name}} / {{.SpaceFields.Name}}{{end}}{{end}}"
	templateFlagUsage = `The text/template to use for output.
	
Available values are:

	- {{.AccessToken}}
	- {{.OrganizationFields.Name}}
	- {{.OrganizationFields.URL}}
	- {{.SpaceFields.Name}}
`
	debugOutputValue = false
	debugOutputUsage = "Print errors and other output to aid in debugging"
)

var templateStr string
var debugOutputEnabled bool

var debugOutput = ioutil.Discard

func main() {
	flag.StringVar(&templateStr, "t", templateFlagValue, templateFlagUsage)
	flag.BoolVar(&debugOutputEnabled, "x", debugOutputValue, debugOutputUsage)
	flag.Parse()

	if debugOutputEnabled {
		debugOutput = os.Stdout
	}

	config, err := readCloudFoundryConfig()
	if err != nil {
		fmt.Fprintf(debugOutput, "error reading config file: %v\n", err)
		os.Exit(1)
	}

	tmpl, err := template.New("main").Parse(templateStr)
	if err != nil {
		fmt.Fprintf(debugOutput, "unable to parse template: %v\n", err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, config)
	if err != nil {
		fmt.Fprintf(debugOutput, "unable to output template: %v", err)
		os.Exit(1)
	}
}

func readCloudFoundryConfig() (config, error) {
	var config config

	cfHome := os.Getenv("CF_HOME")
	if cfHome == "" {
		var err error
		cfHome, err = os.UserHomeDir()
		if err != nil {
			return config, fmt.Errorf("unable to determine user home: %v", err)
		}
	}

	configFileContent, err := ioutil.ReadFile(path.Join(cfHome, ".cf", "config.json"))
	if err != nil {
		return config, fmt.Errorf("unable to open config.json: %v", err)
	}

	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		return config, fmt.Errorf("unable to parse config.json: %v", err)
	}

	return config, nil
}
