package initializers

func SyncDb() {
	DB.AutoMigrate(
	// &models.User{},
	)

}
