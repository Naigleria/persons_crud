package initializers

import (
	
	"gorm/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Person{})
	//DB.AutoMigrate(&models.UsersCharacters{})
	//DB.AutoMigrate(&models.Character{})
	//DB.AutoMigrate(&models.UserResponse{})
	
}
