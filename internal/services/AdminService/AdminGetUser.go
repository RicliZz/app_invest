package AdminService

import (
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AdminService struct {
	repoUser repository.UserRepository
}

func NewAdminService(repoUser repository.UserRepository) *AdminService {
	return &AdminService{repoUser: repoUser}
}

// @Summary findUsers
// @Description Найти юзер(а/ов) по фамилии(СДЕЛАТЬ ЕЩЁ ПО ПОЧТЕ,ИМЕНИ....)
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   opt  query  string  false  "Фамилия"
// @Success 200 {object} []userModel.User "Найденный юзер"
// @Failure 400 {string} string ""
// @Security BearerAuth
// @Router /admin/all-users [get]
func (s *AdminService) AdminGetUser(c *gin.Context) {
	paramUser := c.Query("opt")
	findedUsers, err := s.repoUser.GetUsersByOpt(&paramUser)
	if err != nil {
		log.Println("Траблы с нахождением юезров от админа")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, findedUsers)
}
