package engine

import (
	"MAS/shared"
	"MAS/wasm/states"
	"MAS/wasm/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
)

//RenderEmptyMap creates the HTML grid layout given rows and columns with CSS, also initialize the map state.
func RenderEmptyMap(this js.Value, inputs []js.Value) interface{} {

	grid := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_GRID)
	root := js.Global().Get("document").Get("documentElement")
	row, errCon1 := strconv.Atoi(js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_ROWS).Get("value").String())
	col, errCon2 := strconv.Atoi(js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_COLS).Get("value").String())
	name := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_NAME).Get("value").String()

	go func() {

		if errCon1 != nil || errCon2 != nil || name == "" {
			utils.ShowMessage("Map dimension is not a number or incorrect name for map!")
			return
		}

		n := row * col
		var gridContent strings.Builder
		gridContent.Grow(64)

		for i := 0; i < n; i++ {
			fmt.Fprintf(&gridContent, "<div class='%s' id='cell-%d' %s='-1' onclick=ModifyCell('cell-%d')></div>\n",
				shared.CSS_CELL, i, shared.V_MAP_CONTENT, i)
		}

		grid.Set(shared.FUNC_INNER_HTML, gridContent.String())

		root.Get("style").Call("setProperty", shared.CSS_GRID_COL, col)
		root.Get("style").Call("setProperty", shared.CSS_GRID_ROW, row)

	}()

	// The map has been rendered so the map state is needed and show these values on the UI as well.
	go states.MapInit(name, col, row)

	return nil

}

//LoadMap loads a map from the server to the Map Editor.
func LoadMap(mapData shared.TypeMap) {
	// Reinit the map data localy.
	states.MapReinitialize(mapData)

	// Draw the map on the map editor with the given description.
	go func() {
		mapSt := states.MapGetState()
		umap := states.MapGetUnitMaping()
		grid := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_GRID)
		root := js.Global().Get("document").Get("documentElement")
		n := uint(mapSt.Cols * mapSt.Rows)
		var gridContent strings.Builder
		var innerHTML strings.Builder
		gridContent.Grow(64)
		innerHTML.Grow(32)

		for i := uint(0); i < n; i++ {
			unit, inMap := umap[i]
			dataContent := -1
			innerHTML.Reset()

			if inMap {
				fmt.Fprintf(&innerHTML, unitHTML(unit))
				dataContent = int(unit[1])
			}

			fmt.Fprintf(&gridContent, "<div class='%s' id='cell-%d' %s='%d'  onclick=ModifyCell('cell-%d')> %s </div>\n",
				shared.CSS_CELL, i, shared.V_MAP_CONTENT, dataContent, i, innerHTML.String())
		}

		grid.Set(shared.FUNC_INNER_HTML, gridContent.String())
		root.Get("style").Call("setProperty", shared.CSS_GRID_COL, mapSt.Cols)
		root.Get("style").Call("setProperty", shared.CSS_GRID_ROW, mapSt.Rows)
	}()

}

//LoadMatch loads the match given the mapdata. Builds the map and initializes the match state.
func LoadMatch(mapData shared.TypeMatch) {
	//Init match state
	go states.MatchInit(mapData)

	//Init and units map
	go func() {
		//Create a mapping of units
		units := map[uint]shared.TypeUnitDesc{}
		for _, unit := range mapData.Map.Units {
			units[unit.Loc] = unit.Desc
		}

		root := js.Global().Get("document").Get("documentElement")
		grid := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MATCH_GRID)
		n := uint(mapData.Map.Cols * mapData.Map.Rows)

		var gridContent strings.Builder
		var innerHTML strings.Builder
		gridContent.Grow(64)
		innerHTML.Grow(32)

		for i := uint(0); i < n; i++ {
			unit, inMap := units[i]
			innerHTML.Reset()

			if inMap {
				fmt.Fprintf(&innerHTML, unitHTML(unit))
			}

			fmt.Fprintf(&gridContent, "<div class='%s' id='match-cell-%d' > %s </div>\n", shared.CSS_CELL, i, innerHTML.String())
		}

		grid.Set(shared.FUNC_INNER_HTML, gridContent.String())
		root.Get("style").Call("setProperty", shared.CSS_MATCH_COL, mapData.Map.Cols)
		root.Get("style").Call("setProperty", shared.CSS_MATCH_ROW, mapData.Map.Rows)
	}()
}

func getCellByID(id uint) js.Value {
	return js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("match-cell-%d", id))
}

//Step changes the UI during the match between two steps.
func Step(step shared.TypeStep) {
	var wg sync.WaitGroup
	prevUnits, prevVFXs := states.MatchGetPrevStep()

	//Clear the units from the previous step.
	go func() {
		wg.Add(1)
		defer wg.Done()

		for _, unit := range prevUnits.Units {
			cell := getCellByID(unit.Loc)
			cell.Set(shared.FUNC_INNER_HTML, "")
		}
	}()

	//Clear the VFXs of the previous match.
	go func() {
		wg.Add(1)
		defer wg.Done()

		for _, vfx := range prevVFXs.VFxs {
			for _, id := range vfx.Fields {
				cell := getCellByID(id)
				cell.Set(shared.FUNC_CLASS_NAME, shared.CSS_CELL)
			}
		}
	}()

	wg.Wait()

	//Add the current units
	go func() {
		for _, unit := range step.Units.Units {
			cell := getCellByID(unit.Loc)
			cell.Set(shared.FUNC_INNER_HTML, unitHTML(unit.Desc))
		}
	}()

	//Add the current VFXs
	go func() {
		for _, vfx := range step.VFxs.VFxs {
			for _, id := range vfx.Fields {
				cell := getCellByID(id)
				cell.Get(shared.FUNC_CLASS_LIST).Call(shared.FUNC_CLASS_TOGGLE, fmt.Sprintf("%s%d", shared.CSS_VFX, vfx.Effect))
			}
		}
	}()
}

//ModifyCell modifies a cell on the map on the Map Editor, based on what editor button has been clicked.
func ModifyCell(this js.Value, inputs []js.Value) interface{} {

	id := inputs[0].String()
	idN := uint(0)
	_, convErr := fmt.Sscanf(id, "cell-%d", &idN)
	cell := js.Global().Get("document").Call(shared.FUNC_GET_E, id)
	var innerHTML strings.Builder
	innerHTML.Grow(64)

	if convErr != nil {
		fmt.Println(convErr.Error())
		return nil
	}

	activeBtn := states.MapGetActiveBtn()
	content := 0

	if activeBtn == states.ID_CLEAR {
		innerHTML.WriteString("")
		content = -1
	} else {
		unit := shared.TypeUnitDesc{}
		unit[1] = activeBtn
		innerHTML.WriteString(unitHTML(unit))
		content = int(activeBtn)
	}

	prevContent, _ := strconv.Atoi(cell.Call(shared.FUNC_GET_ATTRIBUTE, shared.V_MAP_CONTENT).String())

	if err := states.MapUpdate(idN, prevContent); err != nil {
		utils.ShowMessage(err.Error())
		return nil
	}

	cell.Set(shared.FUNC_INNER_HTML, innerHTML.String())
	cell.Call(shared.FUNC_SET_ATTRIBUTE, shared.V_MAP_CONTENT, content)

	return nil
}

//unitHTML generates the HTML code for a unit and returns it as a string.
func unitHTML(unit shared.TypeUnitDesc) string {
	var title strings.Builder
	title.Grow(32)

	fmt.Fprintf(&title, "%s:%s\n", "Category", shared.V_UNIT_MAP_TYPE[int(unit[1])])
	for i := 1; i < len(unit); i++ {
		fmt.Fprintf(&title, "%s:%d\n", shared.V_UNIT_MAP[i], unit[i])
	}

	return fmt.Sprintf("<div class='%s%d %s%d' data-toggle='tooltip' data-html='true' title='%s'></div>",
		shared.CSS_TEAM, unit[0], shared.CSS_UNIT, unit[1], title.String())
}
