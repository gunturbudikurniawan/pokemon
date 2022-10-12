package postgre

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var (
	ErrNoDatabaseConnection = errors.New("there is no connection")
	ErrColumnEmpty          = errors.New("column can't be empty")
	ErrTableEmpty           = errors.New("table name can't be empty")
	ErrParamInvalid         = errors.New("paramater is invalid")
)

type Postgre struct {
	//DB Configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	//DB connection
	DB *sqlx.DB
}

type PsqlDb struct {
	*Postgre
}

var (
	PSQL *PsqlDb
)

func InitPostgre() error {
	PSQL = new(PsqlDb)

	//init psql model
	PSQL.Postgre = &Postgre{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	//connect to db psql
	if err := PSQL.Postgre.OpenConnection(); err != nil {
		return err
	}

	//check ping
	if err := PSQL.DB.Ping(); err != nil {
		return err
	}

	// table migration
	query, err := ioutil.ReadFile("./migrations/migration.sql")
	if err != nil {
		return err
	}

	if _, err := PSQL.DB.Exec(string(query)); err != nil {
		panic(err)
	}

	return nil
}

func (p *Postgre) OpenConnection() error {
	//initialize path string
	path := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	//open postgree
	dbConnection, err := sqlx.Open("postgres", path)

	if err != nil {
		return err
	} else {
		p.DB = dbConnection
	}

	//test connection
	if err := p.DB.Ping(); err != nil {
		return err
	}

	return nil
}
