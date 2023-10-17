package Kernel

type databaseConfig struct {
	user     string
	password string
	db       string
}

func getDatabaseConfig() (*databaseConfig, error) {

	// Yaml.Read("/config/database/config.yaml")

	return nil, nil
}
