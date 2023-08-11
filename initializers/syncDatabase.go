package initializers

import "github.com/jetsadawwts/go-jwt/models"


func init() {
	LoadEnvVariables()
	ConnectToDB()
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
