package config

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// SetupConfigSources sets up Viper configuration.
//
// If specific file path is provided but fails when loading it will return an error.
//
// In case of default config location it will not fail if file does not exist,
// but will in any other case such as parse error.
//
// Config precedence (each item takes precedence over the item below it):
// . Flags
// . Env variables
// . Config file
//
// Environment variables are prefixed with `IKE` and have fully qualified names, for example
// in case of `develop` command and its `port` flag corresponding environment variable is
// `IKE_DEVELOP_PORT`.
func SetupConfigSources(configFile string, notDefault bool) error {
	viper.Reset()
	viper.SetEnvPrefix("GHC")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetTypeByDefaultValue(true)

	ext := path.Ext(configFile)
	viper.SetConfigName(strings.TrimSuffix(path.Base(configFile), ext))
	if !contains(SupportedExtensions(), strings.TrimPrefix(path.Ext(ext), ".")) {
		return fmt.Errorf("'%s' extension is not supported. Use one of [%s]", ext, strings.Join(SupportedExtensions(), ", "))
	}
	viper.SetConfigType(ext[1:])
	viper.AddConfigPath(path.Dir(configFile))

	if err := viper.ReadInConfig(); err != nil {
		if notDefault {
			return err
		}

		if _, fileDoesNotExist := err.(viper.ConfigFileNotFoundError); !fileDoesNotExist {
			return err
		}
	}
	return nil
}

// SupportedExtensions returns a slice of all supported config format (as file extensions)
func SupportedExtensions() []string {
	return viper.SupportedExts
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
