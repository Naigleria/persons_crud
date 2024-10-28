package initializers

import (
	"fmt"
	//"log"
	"os"
	"gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	//get env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// build db conection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	} else {
		fmt.Println("Successful db conection")
	}
	/*	
	err = generateTestPersons(DB, 50) // Insertará 50 registros de prueba
    if err != nil {
        log.Fatalf("Error generating test persons: %v", err)
    }

    fmt.Println("Test records generated successfully")*/
}

func generateTestPersons(db *gorm.DB, count int) error {
    persons := models.Persons{}

    for i := 1; i <= count; i++ {
        person := models.Person{
            Name:  fmt.Sprintf("Person %d", i),
            Email: fmt.Sprintf("person%d@example.com", i),
            Age:   uint(20 + (i % 10)), // Edad entre 20 y 29
        }
        persons = append(persons, person)
    }

    // Insertar los registros en la base de datos
    return db.Create(&persons).Error
}