package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
   var port = os.Getenv("PORT")

	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	
}