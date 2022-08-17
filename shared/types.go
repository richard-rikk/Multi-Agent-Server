package shared

// Contains all types that can be used to communicate between the server and the UI.

//TypeUnit: describes a unit, where the Loc is the id of the div on the UI and []byte array describes the unit's properties.
//	0 -> Who owns the unit. 0=Nobody, 1-5=one of the teams.
//	1 -> What kind of unit is this. 0=Base, 1=Obstacle, 2=Mine, 3=Resource, 4=Monster, 5=Defender, 6=Builder
//	2 -> Health of the unit. For some units it is zero. For example Resource and Obstacle
//	3 -> Attack damage of unit. For some units it is zero. For example Resource and Obstacle
//	4 -> The amount of resources on the unit, only apply for the Base and Builder units.
type TypeUnit struct {
	Loc  uint
	Desc TypeUnitDesc
}

type TypeUnits struct {
	Units []TypeUnit
}

type TypeUnitDesc [5]byte

//TypeSettings describes the Settings messages, where the key is an constant from ids.go and int is a value.
type TypeSettings map[string]uint

//TypeMap describes the Map state for the UI, also contains all information to rebuild the map.
//	Name: The name of the map
//	ActiveID: The last clicked editing button. For example: Clear button.
//	Cols: The number of columns in the map.
//	Rows: The number of rows in the map.
//	Desc: The units on the map
type TypeMap struct {
	ID    uint       `json:"ID"`
	Name  string     `json:"Name" validate:"required"`
	Cols  int        `json:"Cols" validate:"required,gte=10,lte=300"`
	Rows  int        `json:"Rows" validate:"required,gte=10,lte=300"`
	Info  [5]int     `json:"Info" validate:"required"`
	Units []TypeUnit `json:"MapDesc" validate:"required"`
}

//TypeMatch describes the current match state for the UI, also used for initialize the match
//	Map: The describtion of the map.
//	StepCnt: Count of all the steps.
type TypeMatch struct {
	Map     TypeMap
	Teams   TypeTeams
	StepCnt uint
}

//TypeStep: Describes one step of a match.
//	Teams: All the data about every team in the match.
//	VFxs: All the visual effects needed to be drawn on the UI.
//	Units: All units on the map.
type TypeStep struct {
	Teams TypeTeams
	VFxs  TypeVFxs
	Units TypeUnits
}

//Team describes a team most important information for the UI.
type TypeTeam struct {
	ID          byte
	Name        string
	BaseHp      uint
	BuilderCnt  uint
	DefenderCnt uint
	Resources   uint
}

type TypeTeams struct {
	Teams []TypeTeam
}

//VFx describes a visual effect for the UI to display during the match
//	Fields -> The list of ids of the effected fields. These correspong to the UI element ids.
//	Effect -> The effect to be displayed on the fields.
//		0: Defender firing effect
type TypeVFx struct {
	Fields []uint
	Effect byte
}

type TypeVFxs struct {
	VFxs []TypeVFx
}

//TypeStat describes a team performance during a match.
//	TeamName 	-> The name of the team.
//	TeamID		-> The team ID
//	MonKilled 	-> Monsters killed during the match by the team.
//	Resources 	-> Resources killed saved every LogFreq steps.
//	BuildUnits	-> 0: Builder units, 1: Defender units
type TypeStat struct {
	TeamID      byte
	TeamName    string
	MonKilled   uint
	UnitsKilled uint
	Resources   []uint
	BuildUnits  [2]uint
}

type TypeStats struct {
	Stats []TypeStat
}
