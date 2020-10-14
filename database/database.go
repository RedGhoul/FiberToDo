package database

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/RedGhoul/fibertodo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
)

func InitDb() {
	configErr := godotenv.Load()
	if configErr != nil {
		log.Fatal(configErr.Error())
		log.Fatal("Error loading .env file")
	}

	DebugFlag, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	var newLogger logger.Interface
	if DebugFlag {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)
	}

	var err error
	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("DBUSER"), os.Getenv("DBPASSWORD")),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", os.Getenv("DBHOST"), os.Getenv("DBPORT")),
		Path:     os.Getenv("DBNAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	if DebugFlag {
		DBConn, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		DBConn, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	}

	if err != nil {
		log.Panic("Failed to connect to the database")
	}
	log.Println("Connected to DB")
	DBConn.AutoMigrate(&models.User{})
	DBConn.AutoMigrate(&models.TodoList{})
	DBConn.AutoMigrate(&models.TodoListItem{})
	log.Println("Ran Auto Migrate")
}
