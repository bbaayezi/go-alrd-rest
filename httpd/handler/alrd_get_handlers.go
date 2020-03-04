package handler

import (
	"fmt"
	"go-alrd-rest/dbquery"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Latest struct {
	Title  string `gorm:"title" json:"title"`
	Author string `gorm:"author" json:"author"`
	Date   string `gorm:"date" json:"date"`
}

type Overview struct {
	Publications   int    `gorm:"publications" json:"publications"`
	Authors        int    `gorm:"authors" json:"authors"`
	TotalCitations int    `gorm:"total_citations" json:"citations"`
	StartYear      string `gorm:"start_year" json:"startYear"`
	EndYear        string `gorm:"end_year" json:"endYear"`
}

func GetLatest(db *gorm.DB) gin.HandlerFunc {
	var err error
	latestPapers := []Latest{}
	latestRows, err := db.Raw(dbquery.GetLatestPublications).Rows()
	if err != nil {
		fmt.Println(err)
	} else {
		defer latestRows.Close()
		for latestRows.Next() {
			// scan to an latest object
			latest := Latest{}
			db.ScanRows(latestRows, &latest)
			latestPapers = append(latestPapers, latest)
		}
	}

	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   latestPapers,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetOverview(db *gorm.DB) gin.HandlerFunc {
	var overview Overview
	err := db.Raw(dbquery.GetOverview).Scan(&overview).Error
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   overview,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetYearlySummary(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetYearlySummary, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetPartnerCountry(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetPartnerCountry, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetKeywords(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetKeywords, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetAuthKeywords(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetAuthKeywords, result)
	return func(c *gin.Context) {
		if err == nil {
			// TODO: wash data
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetPublisher(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetPublisher, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetPublicationName(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetPublicationName, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetContentType(db *gorm.DB) gin.HandlerFunc {
	result := map[string]int{}
	err := GetNumberedMap(db, dbquery.GetContentType, result)
	return func(c *gin.Context) {
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   result,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal server error",
				"data":   err,
			})
		}
	}
}

func GetNumberedMap(db *gorm.DB, queryString string, mapHolder map[string]int) error {
	var err error
	rows, err := db.Raw(queryString).Rows()
	if err == nil {
		defer rows.Close()
		var key string
		var value int
		for rows.Next() {
			rows.Scan(&key, &value)
			mapHolder[key] = value
		}
	} else {
		fmt.Println(err)
		return err
	}
	return nil
}
