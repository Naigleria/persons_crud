package main

import (
	"gorm/controllers"
	"gorm/initializers"

	//"gorm/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	//endpoint

	r := gin.Default()
	
	r.GET("/personas", controllers.GetPersons)
	r.GET("/personas/search", controllers.SearchPersonByNameOrEmail)
	r.POST("/personas", controllers.CreatePerson)
	r.PUT("/personas/:ID", controllers.UpdatePerson)
	r.DELETE("/personas/:ID", controllers.DeletePerson)
	
	
	r.Run()
}
