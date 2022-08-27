package config

import (
	"fmt"
	"os"
	"time"

	"github.com/apotox/strikememongo"
)

type Configs struct {
	MONGO_URL        string
	DATABASE_NAME    string
	ENCRYNOTE_DOMAIN string
	STAGE            string
	MongoTestServer  *strikememongo.Server
}

func (c *Configs) startInMemoryDb() {

	mongoTestServer, err := strikememongo.StartWithOptions(&strikememongo.Options{
		MongoVersion:   "4.0.5",
		StartupTimeout: 10 * time.Second,
	})
	if err != nil {
		fmt.Println("Error starting MongoDB server: ", err)
		panic(err)
	}
	c.MONGO_URL = mongoTestServer.URI()
	c.MongoTestServer = mongoTestServer

}

var CONTEXT_TIMEOUT = 20 * time.Second
var _currentConfig *Configs

func getEnvVar(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetConfig() *Configs {

	if _currentConfig != nil {
		return _currentConfig
	}

	_currentConfig = &Configs{
		STAGE:            getEnvVar("STAGE", "dev"),
		ENCRYNOTE_DOMAIN: getEnvVar("ENCRYNOTE_DOMAIN", "localhost:3000"),
		DATABASE_NAME:    getEnvVar("DATABASE_NAME", fmt.Sprintf("encrynote-%s", getEnvVar("STAGE", "dev"))),
		MONGO_URL:        getEnvVar("MONGO_URL", ""),
	}

	if _currentConfig.STAGE == "dev" || _currentConfig.STAGE == "test" {
		_currentConfig.startInMemoryDb()
	}

	return _currentConfig
}
