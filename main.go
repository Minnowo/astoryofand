package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"minno/astoryofand/crypto"
	"minno/astoryofand/orders"

	"github.com/gin-gonic/gin"
)

func renderFormPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}

func placeOrder(c *gin.Context) {

	var o orders.Order

	if e := c.Bind(&o); e != nil || !orders.CheckValidOrder(o) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	jsonData, err := json.MarshalIndent(&o, "", "  ")

	fmt.Printf("%s\n", jsonData)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !crypto.WritePGPOrder(jsonData) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func main() {

	crypto.FailIfPGPDirNotExists()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.GET("/", renderFormPage)
	router.POST("/api/order", placeOrder)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	router.Run(":8080")
}
