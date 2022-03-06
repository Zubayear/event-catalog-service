package external

import (
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
	log "go-micro.dev/v4/logger"
)

type Host struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// HostProvider Provide the dependency of Host
func HostProvider() (*Host, error) {
	enc := yaml.NewEncoder()
	c, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	if err != nil {
		log.Errorf("Error loading config %v", err)
		return nil, err
	}

	err = c.Load(file.NewSource(file.WithPath("./config.yaml")))
	if err != nil {
		return nil, err
	}

	host := &Host{}
	if err := c.Get("hosts", "database").Scan(host); err != nil {
		return nil, err
	}

	return host, nil
}
