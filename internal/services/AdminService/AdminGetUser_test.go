package AdminService

import (
	"github.com/RicliZz/app_invest/internal/models/userModel"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAdminRepo struct {
	mock.Mock
}

func (m *MockAdminRepo) GetUserById(userId int64) (*userModel.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*userModel.User), args.Error(1)
}

func (m *MockAdminRepo) GetUsersWithStartups() ([]userModel.User, error) {
	return []userModel.User{}, nil
}
func (m *MockAdminRepo) GetUsersNotConfirmed() ([]userModel.User, error) {
	return []userModel.User{}, nil
}

func TestSearchOneUser(t *testing.T) {
	mockRepo := new(MockAdminRepo)
	adminService := &AdminService{repoAdmin: mockRepo}

	expectedUser := &userModel.User{ID: 1, FirstName: "Ivan", LastName: "Ivanov", MiddleName: "Ivanovich"}
	mockRepo.On("GetUserById", int64(1)).Return(expectedUser, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/user/:id", adminService.SearchOneUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"firstName":"Ivan"`)
}
