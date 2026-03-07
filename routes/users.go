package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicao/minimal-goapi/models"
	"github.com/nicao/minimal-goapi/utils"
)

// signup godoc
// @Summary      Cadastro de usuário
// @Description  Cria uma nova conta de usuário
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.UserCredentials  true  "Email e senha"
// @Success      201          {object}  map[string]interface{}
// @Failure      400          {object}  map[string]string
// @Router       /signup [post]
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data" + fmt.Sprint(err)})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save patrão" + fmt.Sprint(err)})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"id":      user.ID})

}

// login godoc
// @Summary      Login
// @Description  Autentica usuário e retorna token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.UserCredentials  true  "Email e senha"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /login [post]
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data" + fmt.Sprint(err)})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	tokenJwt, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "auth error: " + fmt.Sprint(err)})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successefull!",
		"token":   tokenJwt})
}
