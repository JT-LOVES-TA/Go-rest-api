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
	Id               int     `json:"id"`
	State            int     `json:"state"`
	AuctionStartTime string  `json:"auctionStartTime"`
	AuctionEndTime   string  `json:"auctionEndTime"`
	Price            float64 `json:"price"`
	Winner           Winner  `json:"winner"`
}

type Winner struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Verified bool   `json:"verified"`
	Bio      string `json:"bio"`
	Picture  string `json:"picture"`
	Header   string `json:"header"`
	Address  string `json:"address"`
}

func main() {

	file, _ := ioutil.ReadFile("products.json")
	data := Products{}

	_ = json.Unmarshal([]byte(file), &data)
	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.IndentedJSON(http.StatusOK, data.Products)
	})
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.IndentedJSON(http.StatusOK, data.Products[i])
	})
	router.Run("localhost:8080")
}
