package main

import (
	"path"
	"path/filepath"

	"github.com/docktermj/go-logger/logger"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Interface object for viper.BindFlagValue(...)
// ----------------------------------------------------------------------------

type myFlag struct {
	FlagName string
	Value    string
	Type     string
}

func (f myFlag) HasChanged() bool    { return false }
func (f myFlag) Name() string        { return f.FlagName }
func (f myFlag) ValueString() string { return f.Value }
func (f myFlag) ValueType() string   { return f.Type }

// ----------------------------------------------------------------------------
// Internal utility functions
// ----------------------------------------------------------------------------

// Given a file path, return just the filename with no path and no extension.
func baseFilename(filename string) string {
	return filepath.Base(filename[0 : len(filename)-len(path.Ext(filename))])
}

// Safe extraction of string option value.
func getFlagString(args map[string]interface{}, flag string, theDefault string) string {
	result := theDefault
	rawResult := args[flag]
	if rawResult != nil {
		result = rawResult.(string)
	}
	return result
}

// ----------------------------------------------------------------------------
// Internal setConfigurationXXX(...) functions.
// ----------------------------------------------------------------------------

type keysString struct {
	Config string
	Os     string
	Option string
}

type keysStringSlice struct {
	Config string
	Os     string
	Option string
}

type keysBool struct {
	Config string
	Os     string
	Option string
}

// Set a key's various values for string value.
func setConfigurationString(key keysString, args map[string]interface{}) {

	if key.Os != "" {
		viper.BindEnv(key.Config, key.Os)
	}

	// TODO: Once https://github.com/spf13/viper/issues/369 is fixed, restructure the following to use viper.BindFlagValue(...)
	if key.Option != "" {
		rawResult := args[key.Option]
		if rawResult != nil {
			// viper.BindFlagValue(key, myFlag{MyName: key, Value: rawResult.(string), Type: "string"})
			viper.Set(key.Config, rawResult.(string)) // Work-around for Bug #369
		}
	}
}

// Set a key's various values for boolean value.
func setConfigurationBool(key keysBool, args map[string]interface{}) {

	if key.Os != "" {
		viper.BindEnv(key.Config, key.Os)
	}

	// TODO: Once https://github.com/spf13/viper/issues/369 is fixed, restructure the following to use viper.BindFlagValue(...)
	if key.Option != "" {
		rawResult := args[key.Option]
		if rawResult != nil {
			// viper.BindFlagValue(key, myFlag{MyName: key, Value: rawResult.(bool), Type: "string"})
			viper.Set(key.Config, rawResult.(bool)) // Work-around for Bug #369
		}
	}
}

// ----------------------------------------------------------------------------
// Load configuration.
// ----------------------------------------------------------------------------

func LoadConfig(args map[string]interface{}) {

	// ------------------------------------------------------------------------
	// Load configuration file.
	// ------------------------------------------------------------------------

	// Set paths where the configuration may be found.

	commandlineOption := getFlagString(args, ETCD_OPTION_CONFIGURATION, "")
	if commandlineOption != "" { // Configuration file was specified in a commandline option.
		viper.SetConfigName(baseFilename(commandlineOption)) // Name of configuration file without extension.
		viper.AddConfigPath(filepath.Dir(commandlineOption))
	} else { // Configuration not specified in commandline option.  Look in designated places.
		viper.SetConfigName("go-hello-etcd") // Name of configuration file without extension.
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("$HOME/go/src/github.com/BixData/go-hello-etcd/")
		viper.AddConfigPath(".")
	}

	// Read the configuration file.

	err := viper.ReadInConfig()
	if err != nil {
		logger.Info("No configuration file used.")
	}

	// ------------------------------------------------------------------------
	// Load "string" configuration variables.
	// ------------------------------------------------------------------------

	stringKeys := []keysString{
		keysString{
			Config: ETCD_CONFIG_ADVERTISE_CLIENT_URLS,
			Os:     ETCD_OS_ADVERTISE_CLIENT_URLS,
			Option: ETCD_OPTION_ADVERTISE_CLIENT_URLS,
		},
		keysString{
			Config: ETCD_CONFIG_INITIAL_ADVERTISE_PEER_URLS,
			Os:     ETCD_OS_INITIAL_ADVERTISE_PEER_URLS,
			Option: ETCD_OPTION_INITIAL_ADVERTISE_PEER_URLS,
		},
		keysString{
			Config: ETCD_CONFIG_INITIAL_CLUSTER,
			Os:     ETCD_OS_INITIAL_CLUSTER,
			Option: ETCD_OPTION_INITIAL_CLUSTER,
		},
		keysString{
			Config: ETCD_CONFIG_INITIAL_CLUSTER_STATE,
			Os:     ETCD_OS_INITIAL_CLUSTER_STATE,
			Option: ETCD_OPTION_INITIAL_CLUSTER_STATE,
		},
		keysString{
			Config: ETCD_CONFIG_INITIAL_CLUSTER_TOKEN,
			Os:     ETCD_OS_INITIAL_CLUSTER_TOKEN,
			Option: ETCD_OPTION_INITIAL_CLUSTER_TOKEN,
		},
		keysString{
			Config: ETCD_CONFIG_LISTEN_CLIENT_URLS,
			Os:     ETCD_OS_LISTEN_CLIENT_URLS,
			Option: ETCD_OPTION_LISTEN_CLIENT_URLS,
		},
		keysString{
			Config: ETCD_CONFIG_LISTEN_PEER_URLS,
			Os:     ETCD_OS_LISTEN_PEER_URLS,
			Option: ETCD_OPTION_LISTEN_PEER_URLS,
		},
		keysString{
			Config: ETCD_CONFIG_NAME,
			Os:     ETCD_OS_NAME,
			Option: ETCD_OPTION_NAME,
		},
	}

	for _, value := range stringKeys {
		setConfigurationString(value, args)
	}

	// ------------------------------------------------------------------------
	// Load "bool" configuration variables.
	// ------------------------------------------------------------------------

	boolKeys := []keysBool{
		keysBool{
			Config: ETCD_CONFIG_DEBUG,
			Os:     ETCD_OS_DEBUG,
			Option: ETCD_OPTION_DEBUG,
		},
		keysBool{
			Config: ETCD_CONFIG_TRACE,
			Os:     ETCD_OS_TRACE,
			Option: ETCD_OPTION_TRACE,
		},
	}

	for _, value := range boolKeys {
		setConfigurationBool(value, args)
	}
}
