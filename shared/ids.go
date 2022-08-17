package shared

//UnitMapping: contains the names of all attributes of the TypeUnit type.
var (
	V_UNIT_MAP map[int]string = map[int]string{
		0: "Owner",
		1: "Type",
		2: "Health",
		3: "Attack",
		4: "Resources",
	}

	V_UNIT_MAP_TYPE map[int]string = map[int]string{
		0: "Base",
		1: "Obstacle",
		2: "Mine",
		3: "Resource",
		4: "Monster",
		5: "Defender",
		6: "Builder",
	}
)

const (
	// *** IDs on the HTML page ***

	//IDs for the Match Rewatch page
	ID_MATCH_GRID       string = "matchGrid"
	ID_MATCH_ID         string = "matchID"
	ID_MATCH_TEAM_NAME  string = "match-team-"
	ID_MATCH_TEAM_BASE  string = "match-base-"
	ID_MATCH_TEAM_BUILD string = "match-build-"
	ID_MATCH_TEAM_DEF   string = "match-def-"
	ID_MATCH_TEAM_RES   string = "match-res-"
	ID_MATCH_STEP_CURR  string = "match-curr-step"
	ID_MATCH_STEP_LAST  string = "match-last-step"

	//IDs for the Map Editor page
	ID_MAP_VIS       string = "mapIDVis"
	ID_MAP_NAME      string = "mapName"
	ID_MAP_GRID      string = "mapGrid"
	ID_MAP_ROWS      string = "mapRows"
	ID_MAP_COLS      string = "mapCols"
	ID_MAP_LIST      string = "mapList"
	ID_MAP_ADD_BASE  string = "addBaseBtn"
	ID_MAP_ADD_OBS   string = "addObsBtn"
	ID_MAP_ADD_MINE  string = "addMineBtn"
	ID_MAP_ADD_RESRC string = "addResrcBtn"
	ID_MAP_ADD_MON   string = "addMonBtn"
	ID_MAP_CLEAR_BTN string = "clearMapBtn"

	//IDs for the Analyze Match	page
	ID_GRAPH_RESRC string = "resourceGraph"

	//IDs for the Settings page.
	//	! = The given value has to correspond to one of the attribute names of type Settings in the database/types package
	ID_SET_LOAD_BTN    string = "loadSettingsBtn"
	ID_SET_GAME_FORM   string = "gameSettingsForm"
	ID_SET_BASE_HP     string = "BaseHp"    //!
	ID_SET_VIS_RANGE   string = "VisRange"  //!
	ID_SET_MAX_STEP    string = "MaxStep"   //!
	ID_SET_ALT_WIN_CON string = "AltWin"    //!
	ID_SET_DEF_HP      string = "DefHp"     //!
	ID_SET_DEF_ATK     string = "DefAtk"    //!
	ID_SET_DEF_COST    string = "DefCost"   //!
	ID_SET_DEF_RANGE   string = "DefRange"  //!
	ID_SET_BLD_HP      string = "BuildHp"   //!
	ID_SET_BLD_ATK     string = "BuildAtk"  //!
	ID_SET_BLD_COST    string = "BuildCost" //!
	ID_SET_BLD_CAP     string = "BuildCap"  //!
	ID_SET_MNT_HP      string = "MonstHp"   //!
	ID_SET_MNT_ATK     string = "MonstAtk"  //!
	ID_SET_LOG_FREQ    string = "LogFreq"   //!

	// *** CSS variable IDs and classes ***
	CSS_GRID_COL     string = "--grid-column-cnt"
	CSS_GRID_ROW     string = "--grid-row-cnt"
	CSS_MATCH_COL    string = "--match-row-cnt"
	CSS_MATCH_ROW    string = "--match-col-cnt"
	CSS_MAP_BG       string = "--map-bg-color"
	CSS_TILE_COLOR   string = "--map-tile-color"
	CSS_TILE_SIZE    string = "--map-tile-size"
	CSS_TILE_GAP     string = "--map-tile-gap"
	CSS_TEAM_1_COLOR string = "--team-color-1"
	CSS_TEAM_2_COLOR string = "--team-color-2"
	CSS_TEAM_3_COLOR string = "--team-color-3"
	CSS_TEAM_4_COLOR string = "--team-color-4"
	CSS_TEAM_5_COLOR string = "--team-color-5"
	CSS_UNIT         string = "unit-"
	CSS_TEAM         string = "team-"
	CSS_VFX          string = "vfx-"
	CSS_CELL         string = "cell"

	// *** Other values on the HTML page ***
	V_MAP_CONTENT string = "data-content" //In the map cell divs, a data-content property is provided. The possible values are the following.
	V_PLOT        string = "Plotly"

	// *** API ***
	API_GET    string = "GET"
	API_POST   string = "POST"
	API_DELETE string = "DELETE"

	// *** JS callable function names ***
	FUNC_GET_E         string = "getElementById"
	FUNC_GET_ATTRIBUTE string = "getAttribute"
	FUNC_SET_ATTRIBUTE string = "setAttribute"
	FUNC_CLASS_NAME    string = "className"
	FUNC_CLASS_LIST    string = "classList"
	FUNC_CLASS_TOGGLE  string = "toggle"
	FUNC_INNER_HTML    string = "innerHTML"

	// *** KEYS ***
	//The server and the UI communicate with  map[string]interface{} types, so the actual types don't have to be exposed to the UI.
	//	! = The given value has to correspond to one of the attribute names of type Map in the database/types package.
	KEY_MAP_NAME string = "Name"    //!
	KEY_MAP_ROWS string = "Rows"    //!
	KEY_MAP_COLS string = "Cols"    //!
	KEY_MAP_DESC string = "MapDesc" //!
)
