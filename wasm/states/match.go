package states

import (
	"MAS/shared"
	"fmt"
	"syscall/js"
)

type MatchState struct {
	currStep uint
	lastStep uint
	units    shared.TypeUnits
	vfxs     shared.TypeVFxs
	teams    shared.TypeTeams
}

var matchSt MatchState

func MatchInit(data shared.TypeMatch) {

	matchSt = MatchState{
		currStep: 0,
		lastStep: data.StepCnt,
		units:    shared.TypeUnits{Units: data.Map.Units},
		vfxs:     shared.TypeVFxs{},
		teams:    data.Teams,
	}

	teamsInit(data.Teams)
	updateStep(true)

}

//teamsInit sets the team information on the UI and clears the rest of the information brackets.
func teamsInit(teams shared.TypeTeams) {
	for i := 0; i < 5; i++ {
		teamName := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_NAME, i))
		baseHp := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_BASE, i))
		defCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_DEF, i))
		buildCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_BUILD, i))
		resCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_RES, i))

		if i < len(teams.Teams) {
			teamName.Set("value", teams.Teams[i].Name)
			baseHp.Set("value", teams.Teams[i].BaseHp)
			defCnt.Set("value", teams.Teams[i].DefenderCnt)
			buildCnt.Set("value", teams.Teams[i].BuilderCnt)
			resCnt.Set("value", teams.Teams[i].Resources)
		} else {
			teamName.Set("value", "")
			baseHp.Set("value", "")
			defCnt.Set("value", "")
			buildCnt.Set("value", "")
			resCnt.Set("value", "")
		}

	}
}

//stepUpdate updates the step, if last is true it updates the final
// step value as well.
func updateStep(last bool) {
	js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MATCH_STEP_CURR).Set("value", matchSt.currStep)

	if last {
		js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MATCH_STEP_LAST).Set("value", matchSt.lastStep)
	}
}

//updateTeams updates the teams on the UI.
func updateTeams() {
	for i, team := range matchSt.teams.Teams {
		teamName := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_NAME, i))
		baseHp := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_BASE, i))
		defCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_DEF, i))
		buildCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_BUILD, i))
		resCnt := js.Global().Get("document").Call(shared.FUNC_GET_E, fmt.Sprintf("%s%d", shared.ID_MATCH_TEAM_RES, i))

		teamName.Set("value", team.Name)
		baseHp.Set("value", team.BaseHp)
		defCnt.Set("value", team.DefenderCnt)
		buildCnt.Set("value", team.BuilderCnt)
		resCnt.Set("value", team.Resources)
	}
}

//update updates all element on the Match page UI after a step.
func update() {
	go updateStep(false)
	go updateTeams()
}

//MatchStep updates the state of the match on the UI side. Including the units, vfxs, step count, and team data.
//Using this function deletes the information of the previous step.
func MatchStep(step shared.TypeStep) {
	matchSt.currStep += 1
	matchSt.teams = step.Teams
	matchSt.units = step.Units
	matchSt.vfxs = step.VFxs

	update()
}

//MatchGetPrevStep returns the state saved as the previous step.
func MatchGetPrevStep() (shared.TypeUnits, shared.TypeVFxs) {
	return matchSt.units, matchSt.vfxs
}
