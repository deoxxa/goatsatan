package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Goat struct {
	Id               int
	Name             string
	HairLength       int64
	LegCount         int
	Dead             bool
	Alive            bool
	FavouriteGrassId int
	FavouriteGrass   *Grass `sql:"-"`
	Owners           []Satan
}

func (g Goat) IsZombie() bool {
	return g.Dead && g.Alive
}

type Grass struct {
	Id      int
	Colour  string
	Dryness int
}

type Satan struct {
	Id              int
	PitchforkLength int
	Temperature     uint64
	GoatId          int
	Goat            *Goat `sql:"-"`
}

func main() {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	gorm.DefaultCallback.Create().After("gorm:before_create").Register("app:before", func(scope *gorm.Scope) {
		fmt.Printf("%s before create\n", scope.TableName())
	})

	gorm.DefaultCallback.Create().After("gorm:after_create").Register("app:after", func(scope *gorm.Scope) {
		fmt.Printf("%s before create\n", scope.TableName())
	})

	db.LogMode(true)

	fmt.Printf("making a table maybe i don't know\n")

	if err := db.AutoMigrate(&Goat{}, &Grass{}, &Satan{}).Error; err != nil {
		panic(err)
	}

	fmt.Printf("we have a table\n")

	g := Goat{
		Name:       "terry",
		HairLength: 12,
		LegCount:   5,
		FavouriteGrass: &Grass{
			Colour:  "lavender",
			Dryness: 6,
		},
		Owners: []Satan{
			{
				PitchforkLength: 9001,
				Temperature:     666,
			},
		},
	}

	fmt.Printf("%#v\n", g)

	if err := db.Create(&g).Error; err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", g)

	fmt.Printf("goat is a zombie? %#v\n", g.IsZombie())
}
