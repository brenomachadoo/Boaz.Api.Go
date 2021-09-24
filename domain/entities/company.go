package entities

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID              uint           `json:"Id" gorm:"column:Id;not null;primaryKey;autoIncrement"`
	CorporateId     string         `json:"CorporateId" gorm:"column:CorporateId;size:14;index;not null;unique"`
	CorporateName   string         `json:"CorporateName" gorm:"column:CorporateName;size:100;z;not null;unique"`
	CorporateNameId string         `json:"CorporateNameId" gorm:"column:CorporateNameId;size:100;not null;unique"`
	CreatedAt       time.Time      `json:"CreatedAt" gorm:"column:CreatedAt;not null;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `json:"UpdatedAt" gorm:"column:UpdatedAt;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt `gorm:"column:DeletedAt"`
}

// Set User's table name to be `profiles`
func (Company) TableName() string {
	return "Company"
}

/* EXEMPLO DA ESTRTURA COMPLETA

package main

import (
	"fmt"
	"time"

	faker "github.com/bxcodec/faker/v3"
)

type Set struct {
	Id         uint64    `faker:"-"`
	GameId     uint64    `gorm:"not null;index:idx_sets_code_game_id,unique"`
	Name       string    `gorm:"size:255;not null" faker:"last_name,keep"`
	Code       string    `gorm:"size:255;not null;index:idx_sets_code_game_id,unique" faker:"unique,len=3,keep"`
	ReleasedAt time.Time `gorm:"type:DATETIME NOT NULL"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" faker:"-"`
	Game       *Game     `gorm:"foreignkey:GameId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" faker:"keep"`
}

type Game struct {
	Id         uint64    `faker:"-"`
	Name       string    `gorm:"size:255;not null;index:,unique" faker:"first_name,unique,keep"`
	BundleSize uint8     `gorm:"not null;default:1" faker:"bundle_size"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" faker:"-"`
}

func main() {
	game := Game{Name: "Magic The Gathering"}
	set := Set{Name: "Core Set 2021"}
	faker.FakeData(&game)
	faker.FakeData(&set)

	fmt.Println(game.Name == "Magic The Gathering")     // Expect true
	fmt.Println(set.Name == "Core Set 2021")            // Expect true
	// fmt.Println(set.Game.Name == "Magic The Gathering") // Expect true, but get nil pointer dereference

	set2 := Set{Name: "Core Set 2021", Game: &Game{}}
	faker.FakeData(&set)

	fmt.Println(set2.Name == "Core Set 2021") // Expect true
	fmt.Println(set2.Game.Name != "")         // Expect true, but get false
}

*/
