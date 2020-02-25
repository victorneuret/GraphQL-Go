package ORM

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/victorneuret/graphql-go/graph/model"
	"sync"
)

var (
	DB    *gorm.DB
	mutex sync.Mutex
)

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Todo{})
	DB.AutoMigrate(&model.User{})
}
