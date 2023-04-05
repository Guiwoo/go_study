package singleton

import (
	"fmt"
	"sync"
)

/**
For Some Components it only makes sense t o have one system
- Database repository
- Object Factory

the construction call is expensive
- We only do it once
- We give everyone the same instance

Want to prevent anyone creating additional copies
*/

// Load memory for once

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}
func readData(path string) (map[string]int, error) {
	m := make(map[string]int)
	return m, nil
}

// sync.Once , init() -- thread safety
// laziness => whenever the client wants the instance call it

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData(".\\cpitals.txt")
		db := &singletonDatabase{}
		if e == nil {
			db.capitals = caps
		}
		instance = db
	})
	return instance
}

/**
Problem With Singleton
Depends on real data DIP violation
*/

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

type Database interface {
	GetPopulation(name string) int
}

func GetTotalPopulation2(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

var _ Database = (*DummyDatabase)(nil)

/**
Lazy one -time initialization using sync.Once
DIP violation : Depends on interface not concrete type
Singleton is not scary pattern but have to be careful to use
TOo storng dependency on singleton with coupling
*/

func Start2() {
	db := DummyDatabase{}
	pop := GetTotalPopulation2(&db, []string{"alpha", "gamma"})
	fmt.Println("Population of seoul is ", pop)
}

func Start() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("POpulation of seoul is ", pop)
}
