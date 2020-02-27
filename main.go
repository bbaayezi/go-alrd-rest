package main

import (
	"fmt"
	"go-alrd-rest/httpd/handler"
	"go-alrd-rest/secret"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func main() {
	// DB connection
	db, err := gorm.Open("postgres", secret.DBString)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("---- Error connecting database: ", err, ", returning")
		return
	}
	// close the db connection when exit
	defer db.Close()

	// test db

	r := gin.Default()

	r.GET("/latest", handler.GetLatest(db))
	r.GET("/overview", handler.GetOverview(db))
	r.GET("/byYear", handler.GetYearlySummary(db))
	r.GET("/partnerCountry", handler.GetPartnerCountry(db))
	r.GET("/keywords", handler.GetKeywords(db))
	r.GET("/publisher", handler.GetPublisher(db))
	r.GET("/contentType", handler.GetContentType(db))

	r.Run()
}
