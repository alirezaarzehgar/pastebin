package dotenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alirezaarzehgar/pastebin/internal/config"
)

const (
	DefaultCouchDBPort = 5984
)

func loadCouchDBConfig(cfg *config.Config) error {
	cfg.CouchDB.Hostname = os.Getenv(KeyPersistenceCouchDBHostname)
	if cfg.CouchDB.Hostname == "" {
		return fmt.Errorf("%s is empty", KeyPersistenceCouchDBHostname)
	}

	port, err := parseInt(KeyPersistenceCouchDBPort, DefaultCouchDBPort)
	if err != nil {
		return err
	}
	cfg.CouchDB.Port = port

	cfg.CouchDB.Username = os.Getenv(KeyPersistenceCouchDBUsername)
	if cfg.CouchDB.Username == "" {
		return fmt.Errorf("%s is empty", KeyPersistenceCouchDBUsername)
	}
	cfg.CouchDB.Password = os.Getenv(KeyPersistenceCouchDBPassword)
	if cfg.CouchDB.Password == "" {
		return fmt.Errorf("%s is empty", KeyPersistenceCouchDBPassword)
	}

	insecure := os.Getenv(KeyPersistenceCouchDBInsecure)
	if insecure != "" {
		cfg.CouchDB.Insecure, err = strconv.ParseBool(insecure)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", KeyPersistenceCouchDBInsecure, err)
		}
	}

	metadataDBName := os.Getenv(KeyPersistenceCouchDBMetadataDBName)
	if metadataDBName == "" {
		return fmt.Errorf("%s is empty", KeyPersistenceCouchDBMetadataDBName)
	}
	cfg.CouchDB.MetadataDBName = metadataDBName

	return nil
}
