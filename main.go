package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sony/sonyflake"
)

var db *gorm.DB
var sf *sonyflake.Sonyflake

func init() {
	// setup snowflake
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return 1, nil
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	//open a db connection
	var err error
	db, err = gorm.Open("sqlite3", "./tracking.db")
	if err != nil {
		panic("failed to connect database")
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "trackingservice_" + defaultTableName
	}

	//Migrate the schema
	db.AutoMigrate(&Tracking{})
}

func main() {
	router := gin.Default()

	businessV1 := router.Group("/api/v1/business")
	RegisterTrackingApiController(businessV1)
	router.Run()
}
