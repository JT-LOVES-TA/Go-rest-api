package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id        int     `json:"id"`
	State     string  `json:"state"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Price     float64 `json:"price"`
}

func main() {

	file, _ := ioutil.ReadFile("products.json")
	data := Products{}

	_ = json.Unmarshal([]byte(file), &data)

	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, data.Products)
	})
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		c.IndentedJSON(http.StatusOK, data.Products[i])
	})
	router.Run("localhost:8080")
}
