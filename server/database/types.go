package database

import (
	"MAS/shared"

	"gorm.io/gorm"
)

//Settings: Describes the settings used during a match.
//	BaseHp	-> The health of the base unit.
//	VisRange-> The range of vision of all units.
//	MaxStep	-> The max amount of steps that a match can last.
//	AltWin	-> Alternative win conditions: 0=economic(Who has more money), 1=military(Who has more units), 2=Glory(Who has killed more teams or monsters or units)
//	DefHp	-> The health of the defender unit.
//	DefAtk	-> The attack damage of the defender unit.
//	DefRange-> The attack range of the defender unit.
//	BuildHp	-> The health of the builder unit.
//	BuildAtk-> The attack damage of the builder unit.
//	BuildCap-> The number of resources the unit can carry.
//	MonstHp	-> The health of the monster unit.
//	MonstAtk-> The attack damage of the monster unit.
type Settings struct {
	gorm.Model
	BaseHp    uint `json:"BaseHp" binding:"required,gte=1,lte=50"`
	VisRange  uint `json:"VisRange" binding:"required,gte=1,lte=50"`
	MaxStep   uint `json:"MaxStep" binding:"required,gte=1"`
	AltWin    uint `json:"AltWin" binding:"required,gte=1,lte=255"`
	DefHp     uint `json:"DefHp" binding:"required,gte=1,lte=50"`
	DefAtk    uint `json:"DefAtk" binding:"required,gte=1,lte=50"`
	DefCost   uint `json:"DefCost" binding:"required,gte=1,lte=50"`
	DefRange  uint `json:"DefRange" binding:"required,gte=1,lte=50"`
	BuildHp   uint `json:"BuildHp" binding:"required,gte=1,lte=50"`
	BuildAtk  uint `json:"BuildAtk" binding:"required,gte=1,lte=50"`
	BuildCost uint `json:"BuildCost" binding:"required,gte=1,lte=50"`
	BuildCap  uint `json:"BuildCap" binding:"required,gte=1,lte=50"`
	MonstHp   uint `json:"MonstHp" binding:"required,gte=1,lte=50"`
	MonstAtk  uint `json:"MonstAtk" binding:"required,gte=1,lte=50"`
	LogFreq   uint `json:"LogFreq" binding:"required,gte=1"`
}

//Map: Describes the whole map with its stationary units.
//This type is used to save the map in the database.
type Map struct {
	ID      uint
	MapInfo MapDescJSON `json:"MapInfo" binding:"required" grom:"type:json"`
}

//Match: used to contain all information about a given match.
type Match struct {
	gorm.Model
	MapID uint
	Map   Map
	Steps []Step
	Teams TeamsJSON
}

//Step: describes one step in a match. Contains all information needed for the UI.
// current  -> contains all Units that needed to be drawn on the UI.
//TeamData  -> contains all information about every team that needs to be redrawn on the UI.
type Step struct {
	gorm.Model
	MatchID  uint
	State    UnitsJSON `json:"State" binding:"required" grom:"type:json"`
	TeamData TeamsJSON `json:"TeamData" binding:"required" grom:"type:json"`
	Effects  VFxsJSON  `json:"Effects" binding:"required" grom:"type:json"`
}

//Stat: describes a team performance during a match.
//	MatchID 	-> The ID of the match
//	Details		-> Contains all the stats about every team during one match.
type Stat struct {
	gorm.Model
	MatchID uint
	Details StatsJSON `json:"Details" binding:"required" grom:"type:json"`
}

//-------- JSON Types --------

type (
	MapDescJSON shared.TypeMap
	UnitsJSON   shared.TypeUnits
	TeamsJSON   shared.TypeTeams
	VFxsJSON    shared.TypeVFxs
	StatsJSON   shared.TypeStats
)
