package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nicao/minimal-goapi/models"
	"github.com/gin-gonic/gin"
)

// getEvents godoc
// @Summary      Lista todos os eventos
// @Description  Retorna a lista completa de eventos
// @Tags         events
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /events [get]
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "nao pode dar fetch"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": events})
}

// getEvent godoc
// @Summary      Busca evento por ID
// @Description  Retorna um evento específico pelo ID
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "ID do evento"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /events/{id} [get]
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "nao pode dar fetch"})
		return
	}

	context.JSON(http.StatusOK, event)
}

// postEvent godoc
// @Summary      Cria novo evento
// @Description  Cria um novo evento (requer autenticação)
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body      models.Event  true  "Dados do evento"
// @Success      201    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Security     BearerAuth
// @Router       /events [post]
func postEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	event.ID = 1
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "criado pai", "event": event})

}

// putEvent godoc
// @Summary      Atualiza evento
// @Description  Atualiza um evento existente (apenas o dono)
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "ID do evento"
// @Param        event  body      models.Event  true  "Dados do evento"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Security     BearerAuth
// @Router       /events/{id} [put]
func putEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var UpdatedEvent models.Event
	err = context.ShouldBindJSON(&UpdatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}

	UpdatedEvent.ID = eventId
	err = UpdatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "event updated",
		"id":      eventId,
	})
}

// deleteEvent godoc
// @Summary      Deleta evento
// @Description  Remove um evento (apenas o dono)
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "ID do evento"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /events/{id} [delete]
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "nao pode dar parse no int"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "nao deu pra deletar patrão" + fmt.Sprint(err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "deu boa pra apagar chefe, olha ai",
		"id":      eventId,
	})
}
