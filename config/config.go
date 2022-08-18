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

func GetConfig() *Configs {

	if _currentConfig != nil {
		return _currentConfig
	}

	_currentConfig = &Configs{}

	_currentConfig.STAGE = os.Getenv("STAGE")
	if _currentConfig.STAGE == "" {
		_currentConfig.STAGE = "dev"
	}

	_currentConfig.ENCRYNOTE_DOMAIN = os.Getenv("ENCRYNOTE_DOMAIN")
	if _currentConfig.ENCRYNOTE_DOMAIN == "" {
		_currentConfig.ENCRYNOTE_DOMAIN = "localhost:3000"
	}

	if _currentConfig.STAGE == "test" {
		_currentConfig.startInMemoryDb()
	} else {
		_currentConfig.MONGO_URL = os.Getenv("MONGO_URL")
		if _currentConfig.MONGO_URL == "" {
			_currentConfig.MONGO_URL = "mongodb://admin:admin1234@mongo:3014/?directConnection=true"
		}
	}

	_currentConfig.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	if _currentConfig.DATABASE_NAME == "" {
		_currentConfig.DATABASE_NAME = fmt.Sprintf("encrynote-%s", _currentConfig.STAGE)
	}

	return _currentConfig
}
