package minerva

import (
	"errors"
	"log"
	"math/big"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Capital so that it can be exported
type Link struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ShortURL string `gorm:"size:16,uniqueIndex"`
	LongURL  string `gorm:"type:text"`
	Expiry   *time.Time
}

// Connecting to db function
// Takes in gormdb as param
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

// Base 26+26+10 logic goes here
func ShorturlGenerator(id uint) string {
	id64 := int64(id) + 3844
	return big.NewInt(id64).Text(62)
}

// Create and insert url into db
func Create(longURL string, expiry string, db *gorm.DB) (string, error) {
	var expiresAt *time.Time

	if expiry != "" {
		t, err := time.ParseInLocation("2006-01-02T15:04", expiry, time.Local)
		if err != nil {
			return "", err
		}
		expiresAt = &t
	}

	l := Link{
		LongURL: longURL,
		Expiry:  expiresAt,
	}

	if err := db.Create(&l).Error; err != nil {
		return "", err
	}

	shortURL := ShorturlGenerator(l.ID)

	if err := db.Model(&l).Update("short_url", shortURL).Error; err != nil {
		return "", err
	}

	return shortURL, nil
}

func Fetch(shortURL string, db *gorm.DB) (string, error) {
	var link Link
	err := db.
		Where("short_url = ?", shortURL).
		First(&link).
		Error

	if link.Expiry != nil && time.Now().After(*link.Expiry) {
		err = errors.New("link expired")
	}

	return link.LongURL, err
}

func DashboardFetch(db *gorm.DB) ([]Link, error) {
	var links []Link
	err := db.
		Find(&links).
		Error

	return links, err
}
