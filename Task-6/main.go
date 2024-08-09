package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/", getHome)
	r.POST("/register", register)

	r.Run("localhost:8080")
}

type User struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	PassWord string `json:"-"`
}

// in-memory storage
// users []*Users

func getHome(c *gin.Context){
	c.IndentedJSON(200, gin.H{"message": "Welcome!"})
}

func register(c *gin.Context){
	var user User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.IndentedJSON(400, gin.H{"message": err.Error()})
	}


	
}