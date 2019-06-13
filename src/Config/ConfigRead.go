package Config

import (
	"encoding/xml"
	"os"
)

type Configuration struct {
	URL               string `xml:"url"`
	ContainUrl        string `xml:"containUrl"`
	ExecutionInterval int    `xml:"executionInterval"`
}

var conf *Configuration

func GetConfig() *Configuration {
	if conf == nil {
		xmlFile, err := os.Open("./src/Config/conf.xml")
		if err != nil {
			panic(err)
		}
		defer xmlFile.Close()
		if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
			panic(err)
		}
	}
	return conf
}
