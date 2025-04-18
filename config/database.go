package config

type DatabaseConfig struct {
	Host string `default: "localhost"`
	Port string `default: "5432"`
	User string `default: "developer"`
	Pass string `default: "supersecretpassword"`
	Name string `default: "rdf_auth_db_test"`
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host: "localhost",
		Port: "5432",
		User: "developer",
		Pass: "supersecretpassword",
		Name: "rdf_auth_db_test",
	}
}
