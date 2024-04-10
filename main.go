package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the struct for the config file
type Config struct {
	// Paths
	ImportPath  string `mapstructure:"paths.importPath"`
	ArchivePath string `mapstructure:"paths.archivePath"`
	ErrorPath   string `mapstructure:"paths.errorPath"`

	// Scheduler
	LookupIntervall int `mapstructure:"scheduler.lookupInterval"`

	// Logging
	LogFile   string `mapstructure:"logging.logFile"`
	LogLevel  int    `mapstructure:"logging.logLevel"`
	AddSource bool   `mapstructure:"logging.addSource"`

	// Database
	DbHost     string `mapstructure:"database.host"`
	DbPort     int    `mapstructure:"database.port"`
	DbUser     string `mapstructure:"database.user"`
	DbPassword string `mapstructure:"database.password"`
	DbName     string `mapstructure:"database.name"`

	// Source
	SourceName string `mapstructure:"source.name"`
}

// NewConfig creates a new config struct
func NewConfig() *Config {
	return &Config{}
}

// SetupConfig reads the config file and sets up a watcher for changes
func SetupConfig(conf *Config) error {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("could not read config file: %w", err)
	}

	test := viper.GetString("paths.importPath")
	fmt.Println("Value with viper.GetString: ", test)

	//viper.WatchConfig()

	err = viper.Unmarshal(conf)
	if err != nil {
		return fmt.Errorf("unable to decode configuration into struct, %v", err)
	}
	return nil
}

func main() {
	conf := NewConfig()
	err := SetupConfig(conf)
	if err != nil {
		panic(fmt.Errorf("unable to setup the configuration, %v", err))
	}
	fmt.Println("Value from unmarshaled config struct: ", conf.ImportPath)
}
