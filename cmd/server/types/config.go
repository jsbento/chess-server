package types

import (
	"errors"
	"os"
)

type ServerConfig struct {
	MongoHost string
	MongoDB   string
}

func (c *ServerConfig) Validate() error {
	if c.MongoHost == "" {
		return errors.New("mongo host is required")
	}
	if c.MongoDB == "" {
		return errors.New("mongo database is required")
	}

	return nil
}

func (c *ServerConfig) Load() {
	c.MongoHost, c.MongoDB = os.Getenv("MONGO_DB_URI"), os.Getenv("APP_DB")
}

func (c *ServerConfig) LoadAndValidate() error {
	c.Load()
	return c.Validate()
}
