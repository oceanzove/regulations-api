package models

type ConfigService struct {
	Server        ServerConfig        `json:"server" binding:"required"`
	RegulationsDB RegulationsDBConfig `json:"regulations-database" binding:"required"`
}

type ServerConfig struct {
	Port         string `json:"server_port" binding:"required"`
	JWTSecretKey string `json:"jwt_secret_key" binding:"requires"`
}

type RegulationsDBConfig struct {
	Host     string `json:"db_host" binding:"required"`
	Port     string `json:"db_port" binding:"required"`
	Username string `json:"db_username" binding:"required"`
	DBName   string `json:"db_name" binding:"required"`
	SSLMode  string `json:"db_ssl_mode" binding:"required"`
}
