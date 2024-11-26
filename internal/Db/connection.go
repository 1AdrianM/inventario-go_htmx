package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnection() {
	/*	cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("not able to find config, error: %v", err)
		}

		DSN := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%d",
			cfg.DB.Host,
			cfg.DB.Username,
			cfg.DB.Password,
			cfg.DB.DName,
			cfg.DB.Port)

	*/
	DSN := "host=localhost  user=1AdrianM password=sshconnect2020 dbname=inventariogo port=5433"
	var errDb error
	Db, errDb = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if errDb != nil {
		log.Fatalf("not able to connect to database, error: %v", errDb)
	} else {
		log.Println("Db connected")
	}
}
