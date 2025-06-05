package modules

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hieunlt/themis/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/matzefriedrich/parsley/pkg/types"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

var c config

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("CONFIG_FILE", ".env")
	viper.SetDefault("SERVER_PORT", 3000)
	configFile := viper.GetString("CONFIG_FILE")
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot find config file %s", configFile)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
}

func ConfigureDBClient(registry types.ServiceRegistry) error {
	return registry.Register(newDBClient, types.LifetimeSingleton)
}

func newDBClient() *ent.Client {
	client, err := ent.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", c.DBDriver, err)
	}
	defer client.Close()
	return client
}
