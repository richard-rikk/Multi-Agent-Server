package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Some types need to be stored in the database as a json, however the shared package cannot contain
// the code needed for conversion since it is used by both server and the UI. Keeping the convertzation
// code in the shared package would make the UI program significantly bigger.
// However completely decoupling the types in the server and the UI introduces unnecessary complexity.
// Therefore, some convertzation types are introduced that are the same types as the shared ones with different names.
// This allows us to keep the code for storing information to the database on the server side and also it will not introduce
// complexity, since the shared type and its conterpart on the server side will be the same type with different names, so json.Marsal()
// function should be able to unpack the []byte message into either one.

// Value return json value, implement driver.Valuer interface
func (m MapDescJSON) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Make the Map struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (m *MapDescJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed!")
	}

	return json.Unmarshal(b, &m)
}

func (m UnitsJSON) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *UnitsJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed!")
	}

	return json.Unmarshal(b, &m)
}

func (m TeamsJSON) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *TeamsJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed!")
	}

	return json.Unmarshal(b, &m)
}

func (m VFxsJSON) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *VFxsJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed!")
	}

	return json.Unmarshal(b, &m)
}

func (m StatsJSON) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *StatsJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed!")
	}

	return json.Unmarshal(b, &m)
}
