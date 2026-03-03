package minerva

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/big"
)

// capital so that it can be exported 
type Link struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ShortURL string `gorm:"size:16,uniqueIndex"`
	LongURL  string `gorm:"type:text"`
}


// connecting to db function
// takes in gormdb as param
func Connect() (*gorm.DB) {
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
func Create(link string, db *gorm.DB) error {
    db.AutoMigrate(&Link{})

    l := Link{ LongURL: link, }

    if err := db.Create(&l).Error; err != nil {
        return err
    }

    short := ShorturlGenerator(l.ID)

    if err := db.Model(&l).Update("short_url", short).Error; err != nil {
        return err
    }

    return nil
}





