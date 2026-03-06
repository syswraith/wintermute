package minerva

import (
	"log"
	"math/big"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// capital so that it can be exported
type Link struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ShortURL string `gorm:"size:16,uniqueIndex"`
	LongURL  string `gorm:"type:text"`
}

// connecting to db function
// takes in gormdb as param
func Connect() *gorm.DB {
	log.Println("called")

	dsn := "user:password@tcp(127.0.0.1:3306)/wintermute?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect")

	}

	log.Println("db connected")

	return db
}

// base 26+26+10 logic goes here
func ShorturlGenerator(id uint) string {
	id64 := int64(id) + 3844
	return big.NewInt(id64).Text(62)
}

// create and insert url into db
func Create(longURL string, db *gorm.DB) (string, error) {
	db.AutoMigrate(&Link{})

	l := Link{LongURL: longURL}

	err := db.Create(&l).Error

	if err != nil {
		log.Fatal("create failed")
	}

	shortURL := ShorturlGenerator(l.ID)

	err = db.Model(&l).Update("short_url", shortURL).Error

	return shortURL, err
}

func Fetch(shortURL string, db *gorm.DB) (string, error) {
	var link Link
	err := db.
		Where("short_url = ?", shortURL).
		First(&link).
		Error

	return link.LongURL, err
}

func DashboardFetch(db *gorm.DB) ([]Link, error) {
	var links []Link
	err := db.
		Find(&links).
		Error

	return links, err

}
