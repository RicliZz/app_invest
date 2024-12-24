package AdminService

import (
	"github.com/RicliZz/app_invest/pkg/Utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

// @Summary Зайти под пользователем
// @Description Только админ
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param id path int true "ID пользователя, под которым надо войти"
// @Success 200 {object} map[string]string "JWT токен"
// @Failure 400 {string} string "Неверные параметры"
// @Security BearerAuth
// @Router /admin/as-user/{id} [post]
func (s *AdminService) LoginAsAnotherUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	anotherUser, err := s.repoAdmin.GetUserById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := Utils.CreateJWT(secret, anotherUser.ID, anotherUser.Role)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Token": token})
}
