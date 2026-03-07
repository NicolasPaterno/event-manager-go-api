package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicao/minimal-goapi/models"
)

// registerForEvent godoc
// @Summary      Inscrever em evento
// @Description  Registra o usuário autenticado em um evento
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "ID do evento"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /events/{id}/register [post]
func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find event with id: " + strconv.Itoa(int(eventId))})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event "})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered "})
}

// cancelRegistration godoc
// @Summary      Cancelar inscrição
// @Description  Remove o registro do usuário em um evento
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "ID do evento"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "nao deu pra cancelar registro patrão" + fmt.Sprint(err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "cancelou patrão"})

}
