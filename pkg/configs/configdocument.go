package configs

import (
	"encoding/json"
	"strings"

	"github.com/spf13/viper"
)

// DocumentConfig represents a table configuration.
type DocumentConfig struct {
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Authors     []Author `json:"authors"`
	} `json:"metadata"`
}

// NewDocumentConfig creates a DocumentConfig from string. Supports JSON, TOML and YAML string format.
func NewDocumentConfig(documentString string, format ConfigFormat) (DocumentConfig, error) {
	// TODO: Sec: Validation
	document := DocumentConfig{}
	reader := strings.NewReader(documentString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(format.String())
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&document, getCustomDecoder()); err != nil {
		return DocumentConfig{}, err
	}

	return document, nil
}

// JSON converts config to json.
func (config *DocumentConfig) JSON() (string, error) {
	// TODO: Sec: Validation
	bytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
