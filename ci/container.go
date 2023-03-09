package ci

import (
	"log"
	"strings"
	"sync"

	"ramazon/configs"
	"ramazon/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
)

var (
	once sync.Once

	cfg = configs.Load()
)

// MigrationsUp  for migration to database
func MigrationsUp() {
	url, err := utils.ConnectionURLBuilder("migration")
	if err != nil {
		log.Println("Error generating migration url: ", err.Error())
	}
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal("error in creating migrations: ", err.Error())
	}

	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Println("Error in migrating ", err.Error())
		}
	}
}
