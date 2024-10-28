package controllers

import (
	"fmt"
	"gorm/initializers"
	"gorm/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
func GetPersons(c *gin.Context) {

	users := models.Persons{}
	initializers.DB.Find(&users)
	c.JSON(http.StatusOK, users)

}*/



func GetPersons(c *gin.Context) {
    persons := models.Persons{}
    var page, limit int
    var totalRecords int64
    var err error

    // Establece el límite de registros por página a 15 si no se especifica un límite
    limit = 15

    // Lee el parámetro de consulta 'page' y 'limit' si están disponibles
    if c.Query("page") != "" {
        page, err = strconv.Atoi(c.Query("page"))
        if err != nil || page < 1 {
            page = 1 // Página predeterminada si se proporciona una página no válida
        }
    } else {
        page = 1 // Página predeterminada si no se proporciona ningún parámetro de página
    }

    // Calcula el offset
    offset := (page - 1) * limit

    // Contar el número total de registros
    initializers.DB.Model(&models.Persons{}).Count(&totalRecords)

    // Calcular el total de páginas
    totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

    // Realiza la consulta con límite y desplazamiento
    result := initializers.DB.Limit(limit).Offset(offset).Find(&persons)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    // Devuelve los resultados en formato JSON, incluyendo total de páginas y registros
    c.JSON(http.StatusOK, gin.H{
        "data":        persons,
        "totalPages":  totalPages,
        "totalRecords": totalRecords,
        "currentPage": page,
    })
}


func CreatePerson(c *gin.Context){
	var body struct {
		Name         string   `json:"name"`
		Email        string   `json:"email"`
		Age          uint     `json:"age"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	person:= models.Person{Name: body.Name, Email: body.Email, Age: body.Age}
	result := initializers.DB.Create(&person)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create person",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Person created successfully",
	})
}

func getPersonById(c *gin.Context) (models.Person, error) {

	userId, err := strconv.Atoi(c.Param("ID"))

	if err != nil {
		fmt.Println("ID must be a number ")
		return models.Person{}, err
	}

	var person models.Person
	err = initializers.DB.First(&person, userId).Error
	if err != nil {
		fmt.Println("Error trying to retrieve person")
		return person, err
	}

	return person, nil
}

func UpdatePerson(c *gin.Context) {

	var body struct {
		Name         string   `json:"name"`
		Email        string   `json:"email"`
		Age          uint     `json:"age"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	person, err := getPersonById(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	updates := make(map[string]interface{})
	
	if body.Name != person.Name{
		updates["name"] = body.Name
	}
	if body.Email != person.Email{
		updates["email"] = body.Email
	}
	if body.Age != person.Age{
		updates["age"] = body.Age
	}

	// save in db updated person
	if err := initializers.DB.Model(&person).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error trying to save person"})
		return
	}

	// Response with updated person
	c.JSON(http.StatusOK, person)
}

func DeletePerson(c *gin.Context) {

	person, err := getPersonById(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	} else {
		initializers.DB.Delete(&person)
		c.JSON(http.StatusOK, person)
	}

}

func SearchPersonByNameOrEmail(c *gin.Context) {
    var searchPersons []models.Person
    query := c.Query("query")

    
    if err := initializers.DB.Where("name LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&searchPersons).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Retornar la lista de personas encontradas
    c.JSON(http.StatusOK, searchPersons)
}