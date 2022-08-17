package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB

//https://stackoverflow.com/questions/51471973/gorm-automigrate-and-createtable-not-working
// Change db := d.db.AutoMigrate(&m) to db := d.db.AutoMigrate(m) to allow for the reflection to get the type name.
var _migrateList []interface{} = []interface{}{Settings{}, Match{}, Map{}, Step{}, Stat{}}

//Connect: this function has to be called before any other database operation. Given the variable it migrates the database.
// Migrates all schemas. It can lead to dataloss, so only use it when neccessary.
func Connect(migrate bool) error {
	dsn := "host=localhost user=rikk password=pw123 dbname=masdb port=5432 sslmode=disable TimeZone=Europe/Budapest"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_db = db

	if e != nil {
		return e
	}

	//Migrate the database if needed.
	if migrate {
		// Set up the schemas
		if err := _db.AutoMigrate(_migrateList...); err != nil {
			return err
		}

		// Clear any data
		if err := ClearDB(); err != nil {
			return err
		}

		// Init the database with current data
		if err := InitDB(); err != nil {
			return err
		}

	}

	return nil
}

func ClearDB() error {
	for _, table := range _migrateList {
		if result := _db.Unscoped().Where("1 = 1").Delete(table); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

//InitDB: Initializes the database with the necessary data
func InitDB() error {
	defaultSettings := Settings{
		BaseHp:    50,
		VisRange:  5,
		MaxStep:   50_000,
		AltWin:    1,
		DefHp:     5,
		DefAtk:    3,
		DefCost:   3,
		DefRange:  2,
		BuildHp:   2,
		BuildAtk:  1,
		BuildCost: 1,
		BuildCap:  2,
		MonstHp:   10,
		MonstAtk:  5,
		LogFreq:   1000,
	}

	if result := _db.Create(&defaultSettings); result.Error != nil {
		return result.Error
	}

	return nil
}

//SettingsUpdate: updates the settings in the database. If there are no settings, then creates it in the table.
func SettingsUpdate(new *Settings) error {
	old := Settings{}
	result := _db.First(&old)

	if result.Error != nil {
		if result := _db.Create(&new); result.Error != nil {
			return result.Error
		}

		return nil
	}

	new.ID = old.ID
	result = _db.Save(new)

	return result.Error
}

//SettingsGet: return current settings in the db.
func SettingsGet() (*Settings, error) {
	s := Settings{}
	result := _db.First(&s)

	return &s, result.Error
}

//MapSave: Saves a map into the database.
func MapSave(m MapDescJSON) error {
	result := _db.Save(&Map{MapInfo: m})

	return result.Error
}

//MapGetAll: returns all maps saved in the database
func MapGetAll() (m []Map, e error) {
	e = _db.Find(&m).Error
	return
}

//MapGet: returns a specific map
func MapGet(id int) (m Map, e error) {
	e = _db.Find(&m, id).Error
	return
}

//MapDelete: deletes map by ID
func MapDelete(id int) error {
	m := Map{}
	if r := _db.First(&m, id); r.Error != nil {
		return r.Error
	}

	if r := _db.Delete(&m); r.Error != nil {
		return r.Error
	}

	return nil
}

func MatchGet(id int) (m Match, e error) {
	e = _db.Find(&m, id).Error
	return
}
