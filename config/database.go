package config

import "fmt"

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		Getenv("DATABASE_USER", "root"),
		Getenv("DATABASE_PASSWORD", ""),
		Getenv("DATABASE_HOST", "localhost"),
		Getenv("DATABASE_PORT", "3306"),
		Getenv("DATABASE_NAME", "cakes_development"),
	)
}
