package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Driver de conexão com PostgreSQL
)

// Tipo de configuração do banco de dados
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// Carrega as configurações do arquivo JSON
func LoadConfig() (*DBConfig, error) {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuração do banco de dados: %w", err)
	}

	var config DBConfig
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar as configurações do banco de dados: %w", err)
	}
	println(&config)
	return &config, nil
}

// Converte as configurações do DBConfig em uma string de conexão válida para o PostgreSQL
func (config *DBConfig) ToDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)
}

// Abre a conexão com o banco de dados PostgreSQL
func ConectarDB() (*sql.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	dbURL := config.ToDBURL()
	fmt.Println(dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %w", err)
	}

	return db, nil
}
