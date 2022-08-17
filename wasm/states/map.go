package states

import (
	"MAS/shared"
	"errors"
	"syscall/js"
)

const ID_CLEAR byte = 5

//unit: Used by this state to effectenly delete and add new units without searching
type unit map[uint]shared.TypeUnitDesc

//MapState: contains all information for the Map Editor page, where:
//	ms: all information that is going to save on the server
//	activeID: which editor button is active at the moment.
//	units: a map for easy access, the content of this variable has to be copied into ms.Units before returning it.
type MapState struct {
	ms       shared.TypeMap
	activeID byte
	units    unit
}

var mapSt MapState

// MapUpdateUI: Only updates the counters for each unit on the UI.
func MapUpdateUI() {

	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ADD_BASE).Set("innerText", mapSt.ms.Info[0])
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ADD_OBS).Set("innerText", mapSt.ms.Info[1])
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ADD_MINE).Set("innerText", mapSt.ms.Info[2])
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ADD_RESRC).Set("innerText", mapSt.ms.Info[3])
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ADD_MON).Set("innerText", mapSt.ms.Info[4])
}

// MapUpdateUIAll: Update all elements on the UI including map name, row and col counts.
func MapUpdateUIAll() {
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_NAME).Set("value", mapSt.ms.Name)
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ROWS).Set("value", mapSt.ms.Rows)
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_COLS).Set("value", mapSt.ms.Cols)

	MapUpdateUI()
}

// InitMapState: Initialize the map state for the map editor page.
func MapInit(name string, col, row int) {
	mapSt = MapState{
		ms: shared.TypeMap{
			Name:  name,
			Cols:  col,
			Rows:  row,
			Info:  [5]int{0, 0, 0, 0, 0},
			Units: []shared.TypeUnit{},
		},
		units:    unit{},
		activeID: 0,
	}

	MapUpdateUI()

}

// MapStateSetActive: Follows which button has been clicked on the Map Editor page. Call this function with the following values:
//	0 = BaseBtn
//	1 = ObstacleBtn
//	2 = MineBtn
//	3 = ResourcesBtn
//	4 = MonsterBtn
//	5 = ClearBtn
func MapSetActive(this js.Value, inputs []js.Value) interface{} {
	mapSt.activeID = byte(inputs[0].Int())
	return nil
}

// MapStateGetActiveBtn: Returns the active button on the Map Editor page.
func MapGetActiveBtn() byte {
	return mapSt.activeID
}

//MapStateUpdate: updates the values in memory and on the UI. Returns error if a value is bad.
//Also manages the map layout in memory, but only the first two bytes, so it only manages the placement and the type of each unit.
//Call the function with cellContent = -1 if a unit is not being overwritten meaning the cell is empty on the map.
func MapUpdate(id uint, cellContent int) error {
	desc := shared.TypeUnitDesc{}

	if mapSt.activeID == 0 && mapSt.ms.Info[0] >= 5 {
		return errors.New("Maximum 5 players are supported on any map!")
	} else if mapSt.activeID == 5 {
		delete(mapSt.units, id)
	} else {
		mapSt.ms.Info[mapSt.activeID] += 1
		desc[1] = mapSt.activeID
		mapSt.units[id] = desc
	}

	if cellContent != -1 {
		mapSt.ms.Info[cellContent] -= 1
	}

	MapUpdateUI()

	return nil
}

// MapGetState: Returns the current map state from the Map Editor
func MapGetState() shared.TypeMap {
	u := []shared.TypeUnit{}
	for k, v := range mapSt.units {
		u = append(u, shared.TypeUnit{Loc: k, Desc: v})
	}

	mapSt.ms.Units = u

	return mapSt.ms
}

func MapGetUnitMaping() unit {
	return mapSt.units
}

// MapIsEmpty: Tells if the map has elements on it. It can be usefull when saving a map.
func MapIsEmpty() bool {
	return len(mapSt.units) == 0
}

// MapReinitialize: It reinitializes the map state given the map state from the server.
func MapReinitialize(mapData shared.TypeMap) {
	u := unit{}
	for _, unit := range mapData.Units {
		u[unit.Loc] = unit.Desc
	}

	mapSt = MapState{
		ms:       mapData,
		activeID: 0,
		units:    u,
	}

	MapUpdateUIAll()
}
