package database

import (
	"fmt"
	"log"

	"github.com/kennedybaraka/app-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PGHOST     = config.Env("PGHOST")
	PGPASSWORD = config.Env("PGPASSWORD")
	PGPORT     = config.Env("PGPORT")
	PGUSER     = config.Env("PGUSER")
	PGDATABASE = config.Env("PGDATABASE")
	Data       *gorm.DB
)

func ConnectToPostgreSql() {
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", PGUSER, PGPASSWORD, PGHOST, PGPORT, PGDATABASE)
	c, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	Data = c
	fmt.Println("[#] DATABASE CONNECTION HAS BEEN ESTABLISHED.")

}
