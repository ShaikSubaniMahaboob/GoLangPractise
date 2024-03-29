package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	DBCONNECTION string = "host=localhost user=jiten password=admin dbname=sample port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	PORT         string = ":50080"
)

func usage() {
	//fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}
func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	db, err := database.GetConnection(DBCONNECTION)
}
