package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
)

func Connect() (*sql.DB, error) {
	dbService := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	return sql.Open(dbService, dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
}

func Insert(tableName string, fields map[string]interface{}) (*sql.Rows, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	columns := make([]string, len(fields))
	values := make([]interface{}, len(fields))
	parameterised := make([]string, len(fields))
	i := 0
	for k, v := range fields {
		columns[i] = strcase.ToLowerCamel(k)
		values[i] = fmt.Sprint(v)
		parameterised[i] = "?"
		i++
	}

	return db.Query(fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", tableName, strings.Join(columns, ", "), strings.Join(parameterised, ", ")), values...)
}
