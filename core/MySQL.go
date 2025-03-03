package core

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_"github.com/go-sql-driver/mysql"
)

type Config struct {
    Database DatabaseConfig `toml:"database"`
    Server   ServerConfig   `toml:"server"`
}

type DatabaseConfig struct {
    DNS string `toml:"dns"`
}

type ServerConfig struct {
    Port  int  `toml:"port"`
    Debug bool `toml:"debug"`
}

type ConectionMySQL struct{
	DB *sql.DB
	Err string
}

func MySQLConection() *ConectionMySQL {
	error := ""

	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		error = fmt.Sprintf("Error al leer el archivo toml: ", err)
	}

	//user := os.Getenv("USERNAME")
	//pwd := os.Getenv("PASSWORD123")
	//db_name := os.Getenv("DATABASE123")
	//host := os.Getenv("HOST123")

	db, err := sql.Open("mysql", config.Database.DNS)

	if err != nil {
		error = fmt.Sprintf("Error al establecer la conexión con la BD:", err)
	}

	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		error = fmt.Sprintf("Error al hacer ping en la BD:", err)
	}

	
	fmt.Println("Corriendo servidor")

	return &ConectionMySQL{DB: db, Err: error}
}

func(conection *ConectionMySQL) ExecPreparedQuerys(query string, values ...interface{})(sql.Result, error){
	stmt, err := conection.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error al preparar la consulta", err)
	}
	defer stmt.Close()
	
	results, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la consulta", err)
	}
	return results, nil
}

func(conection ConectionMySQL) FetchRows(query string, values... interface{})(*sql.Rows, error){
	rows, err := conection.DB.Query(query, values...)
	if err != nil{
		return nil, fmt.Errorf("Erro al conseguir las filas afectadas")
	}
	return rows, nil
}