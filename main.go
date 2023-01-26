package main

import (
	"dream/infrastructure/datastore"
	"dream/infrastructure/router"
	"dream/registry"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("Hi dream")
	dbSource := ""
	db := datastore.NewDB(&dbSource)
	db.Logger.LogMode(logger.Info)
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	reg := registry.NewRegistry(db)

	r := gin.New()
	r = router.NewRouter(r, reg.NewAppHandler())

	r.Run("localhost:8081")
}
