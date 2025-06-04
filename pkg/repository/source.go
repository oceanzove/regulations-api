package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"regulations-api/models"
)

func NewRegulationsDatabase(config models.ConfigService, environment models.Environment) *sqlx.DB {
	fmt.Println("start database connected")
	database, err := NewPostgresDB(&PostgresDBConfig{
		Host:     config.RegulationsDB.Host,
		Port:     config.RegulationsDB.Port,
		Username: config.RegulationsDB.Username,
		Password: environment.DBPassword,
		DBName:   config.RegulationsDB.DBName,
		SSLMode:  config.RegulationsDB.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize regulation db: %s", err.Error())
	}
	fmt.Println("database connected")
	return database
}
