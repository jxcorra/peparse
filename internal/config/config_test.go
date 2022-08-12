package config_test

import (
	"testing"

	"github.com/jxcorra/peparse/internal/config"
)

func TestParseConfigurationEmptyNotAllowed(t *testing.T) {
	inputConfigData := `{}`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err == nil {
		t.Error("empty configuration")
	}
}

func TestParseConfigurationNoError(t *testing.T) {
	inputConfigData := `{
		"resources": [
		    {
			"url": "https://somesite1.com/",
			"search": [
			    {
				"key": {
				    "element": "a",
				    "class": "someclass1"
				},
				"withText": true,
				"parse": [
				    {
					"attr": "someattr1"
				    }
				]
			    }
			]
		    },
		    {
			"url": "https://somesite2.com/",
			"search": [
			    {
				"key": {
				    "element": "div",
				    "class": "someclass2"
				},
				"parse": [
				    {
					"attr": "someattr2"
				    }
				]
			    }
			]
		    }
		]
	    }`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err != nil {
		t.Error("error while parsing configuration")
	}
}

func TestParseConfigurationUrlRequired(t *testing.T) {
	inputConfigData := `{
		"resources": [
		    {
			"search": [
			    {
				"key": {
				    "element": "a",
				    "class": "someclass1"
				},
				"withText": true,
				"parse": [
				    {
					"attr": "someattr1"
				    }
				]
			    }
			]
		    }
		]
	    }`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err == nil {
		t.Error("url is required in configuration")
	}
}

func TestParseConfigurationNoSearch(t *testing.T) {
	inputConfigData := `{
		"resources": [
		    {
			"url": "https://somesite1.com/",
		    }
		]
	    }`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err == nil {
		t.Error("search is required in configuration")
	}
}

func TestParseConfigurationNoKey(t *testing.T) {
	inputConfigData := `{
		"resources": [
		    {
			"url": "https://somesite1.com/",
			"search": [
			    {
				"withText": true,
				"parse": [
				    {
					"attr": "someattr1"
				    }
				]
			    }
			]
		    }
		]
	    }`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err == nil {
		t.Error("key is required in configuration")
	}
}

func TestParseConfigurationNoParse(t *testing.T) {
	inputConfigData := `{
		"resources": [
		    {
			"url": "https://somesite1.com/",
			"search": [
			    {
				"key": {
				    "element": "a",
				    "class": "someclass1"
				},
				"withText": true,
			    }
			]
		    }
		]
	    }`

	_, err := config.ParseConfiguration([]byte(inputConfigData))
	if err == nil {
		t.Error("parse is required in configuration")
	}
}
