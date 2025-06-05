package modules

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hieunlt/themis/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/matzefriedrich/parsley/pkg/types"
	"github.com/spf13/viper"
)

// config stores the configuration parameters for the application.
type config struct {
	// DBDriver specifies the database driver to use (e.g., "mysql", "sqlite3").
	DBDriver string `mapstructure:"DB_DRIVER"`
	// DBSource specifies the data source name for the database connection.
	DBSource string `mapstructure:"DB_SOURCE"`
}

// c is a global variable that holds the configuration for the application.
var c config

// init is a special function that is automatically executed when the package is initialized.
func init() {
	// Automatically load environment variables from the environment.
	viper.AutomaticEnv()
	// Set default values for configuration parameters.
	viper.SetDefault("CONFIG_FILE", ".env")
	viper.SetDefault("SERVER_PORT", 3000)
	// Get the configuration file path from the environment or default value.
	configFile := viper.GetString("CONFIG_FILE")
	// Set the configuration file to be used by Viper.
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	// If there is an error while reading the configuration file, log a fatal error.
	if err != nil {
		log.Fatalf("cannot find config file %s", configFile)
	}

	err = viper.Unmarshal(&c)
	// If there is an error while unmarshaling the configuration, log a fatal error.
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
}

// ConfigureDBClient configures the database client and registers it with the service registry.
func ConfigureDBClient(registry types.ServiceRegistry) error {
	return registry.Register(newDBClient, types.LifetimeSingleton)
}

// newDBClient creates a new database client using the configuration parameters.
func newDBClient() *ent.Client {
	client, err := ent.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", c.DBDriver, err)
	}
	return client
}
