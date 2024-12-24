package AdminService

import (
	"github.com/RicliZz/app_invest/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type SearchParams struct {
	NotConfirmed bool `json:"NotConfirmed" example:"false"`
	HaveStartups bool `json:"HaveStartups" example:"false"`
}

type AdminService struct {
	repoAdmin repository.AdminRepository
}

func NewAdminService(repoAdmin repository.AdminRepository) *AdminService {
	return &AdminService{repoAdmin: repoAdmin}
}

// @Summary Найти пользователя по ID
// @Description ID
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "ID пользователя"
// @Success 200 {object} userModel.User "Найденный юзер"
// @Failure 400 {string} string "Пользователь по ID не найден"
// @Security BearerAuth
// @Router /admin/search-one/{id} [post]
func (s *AdminService) SearchOneUser(c *gin.Context) {
	userId := c.Param("id")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Println("Введён неверный формат ID")
		return
	}
	users, err := s.repoAdmin.GetUserById(userIdInt)
	if err != nil {
		log.Println("Ошибка поиска пользователя по ID:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Найти всех пользователей по параметрам
// @Description ФИО/Название стартапа
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param searchParams body SearchParams true "Параметры поиска"
// @Success 200 {object} []userModel.User "Найденные пользователи"
// @Failure 400 {string} string "Неверные параметры"
// @Security BearerAuth
// @Router /admin/get-all [post]
func (s *AdminService) SearchAllUsers(c *gin.Context) {
	var params SearchParams
	//Десериализация в JSON
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Проверка на параметр notConfirmed (Поиск людей, кто не подтверждён модерацией)
	if params.NotConfirmed == true {
		users, err := s.repoAdmin.GetUsersNotConfirmed()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
	//Проверка на параметр startup (Поиск людей, у кого есть хотя бы один стартап)
	if params.HaveStartups == true {
		users, err := s.repoAdmin.GetUsersWithStartups()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}

}
