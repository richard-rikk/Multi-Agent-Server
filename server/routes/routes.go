package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"MAS/server/database"
	"MAS/shared"

	"github.com/go-playground/validator/v10"
)

//*** SETTINGS ***

func getSettings(c *gin.Context) {
	setting, err := database.SettingsGet()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := shared.TypeSettings{
		shared.ID_SET_BASE_HP:     setting.BaseHp,
		shared.ID_SET_VIS_RANGE:   setting.VisRange,
		shared.ID_SET_MAX_STEP:    setting.MaxStep,
		shared.ID_SET_ALT_WIN_CON: setting.AltWin,
		shared.ID_SET_DEF_HP:      setting.DefHp,
		shared.ID_SET_DEF_ATK:     setting.DefAtk,
		shared.ID_SET_DEF_COST:    setting.DefCost,
		shared.ID_SET_DEF_RANGE:   setting.DefRange,
		shared.ID_SET_BLD_HP:      setting.BuildHp,
		shared.ID_SET_BLD_ATK:     setting.BuildAtk,
		shared.ID_SET_BLD_COST:    setting.BuildCost,
		shared.ID_SET_BLD_CAP:     setting.BuildCap,
		shared.ID_SET_MNT_HP:      setting.MonstHp,
		shared.ID_SET_MNT_ATK:     setting.MonstAtk,
		shared.ID_SET_LOG_FREQ:    setting.LogFreq,
	}

	c.JSON(http.StatusOK, response)
}

func postSettings(c *gin.Context) {
	var newSetting database.Settings

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&newSetting); err != nil {
		c.AbortWithError(http.StatusNotAcceptable, err)
		return
	}

	if err := database.SettingsUpdate(&newSetting); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, newSetting)
}

//*** MAPS ***

func postMap(c *gin.Context) {
	var newMap database.MapDescJSON

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newMap); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	//Have to validate the JSON struct separately.
	if err := validator.New().Struct(newMap); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err.(validator.ValidationErrors)[0]})
		return
	}

	if err := database.MapSave(newMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func getAllMaps(c *gin.Context) {

	maps, err := database.MapGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	result := []database.MapDescJSON{}
	for _, m := range maps {
		m.MapInfo.ID = m.ID
		result = append(result, m.MapInfo)
	}

	c.JSON(http.StatusOK, result)
}

func getMapById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	m, e := database.MapGet(id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, m.MapInfo)

}

func deleteMapById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := database.MapDelete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func getMatchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	//match, e := database.MatchGet(id)
	//
	//if e != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK,
	//	shared.TypeMatch{
	//		Map:     shared.TypeMap(match.Map.MapInfo),
	//		StepCnt: uint(len(match.Steps)),
	//	})

	m, e := database.MapGet(id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK,
		shared.TypeMatch{
			Map:     shared.TypeMap(m.MapInfo),
			StepCnt: 10,
		})

}

//CreateRouter: Creates all routes for the application
func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	//Server static files
	r.StaticFS("/index", http.Dir("assets"))

	//API routes
	r.GET("/settings", getSettings)
	r.POST("/settings", postSettings)

	r.GET("/maps", getAllMaps)
	r.GET("/maps/:id", getMapById)
	r.POST("/maps", postMap)
	r.DELETE("/maps/:id", deleteMapById)

	r.GET("/match/:id", getMatchById)

	return r

}
