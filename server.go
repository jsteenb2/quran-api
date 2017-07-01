package main

import (
	"log"
	"os"
	"strconv"

	"net/http"

	"github.com/boltdb/bolt"
	m "github.com/jsteenb2/boltDBquran/models"
	"github.com/labstack/echo"
)

var (
	db      *bolt.DB
	qBucket = []byte("quran")
)

func main() {
	dir, osErr := os.Getwd()
	if osErr != nil {
		log.Fatal(osErr)
	}

	var err error
	db, err = bolt.Open(dir+"/quran.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/api", surahFind)

	e.Logger.Fatal(e.Start(":3333"))
}

func surahFind(c echo.Context) error {
	suraNum, _ := strconv.Atoi(c.QueryParam("surah"))
	surah, err := m.GetSurah(qBucket, []byte{byte(suraNum)}, db)
	if err != nil {
		log.Println(err)
	}
	log.Println(surah.Name)
	return c.String(http.StatusOK, surah.String())
}
