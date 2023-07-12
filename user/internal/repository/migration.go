package repository

func Migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&User{},
		)

	if err != nil {
		panic(err)
	}
}
