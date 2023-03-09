package postgres

import (
	"fmt"
	"ramazon/pkg/utils"
	"sync"

	_ "github.com/lib/pq" //pq for connection
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instanceG *gorm.DB
	once      sync.Once
)

// DB ...
func DB() *gorm.DB {
	once.Do(func() {
		dsn, err := utils.ConnectionURLBuilder("postgres")
		if err != nil {
			panic(err)
		}

		instanceG, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("GORM connections: %v", err.Error())
			panic(err)
		}
	})

	return instanceG
}
