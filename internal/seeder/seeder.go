package seeder

import (
	"gorm.io/gorm"
	"errors"
)

type Seeder interface {
	Name() string
	Seed(engine *gorm.DB) error
}

var seeders = []Seeder{} 

func Execute(engine *gorm.DB, name string, count int) (err error) {
	err = errors.New("Seeder not found")
	for i := 0; i < count; i++ {
		for _, seeder := range seeders {
			if seeder.Name() == name {
				if err = seeder.Seed(engine); err != nil {
					return
				}
			}
		}
	}
	return
}

func RegisterSeeder(seeder Seeder) {
	seeders = append(seeders, seeder)
}
