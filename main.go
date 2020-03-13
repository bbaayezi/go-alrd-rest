package main

import (
	"fmt"
	"go-alrd-rest/httpd/handler"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func main() {
	// DB connection
	dbHost := os.Getenv("ALRD_DB_HOST")
	dbPort := os.Getenv("ALRD_DB_PORT")
	dbUser := os.Getenv("ALRD_DB_USER")
	dbName := os.Getenv("ALRD_DB_NAME")
	dbPassword := os.Getenv("ALRD_DB_PASSWORD")
	dbSSLMode := os.Getenv("ALRD_DB_SSLMODE")
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode))
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
	r.GET("/authKeywords", handler.GetAuthKeywords(db))
	r.GET("/publicationName", handler.GetPublicationName(db))

	r.Run()
}
