package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"retailStore/config"
	"retailStore/lib/seeders"
	"retailStore/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	seeders.ItemSeed()
	model, err := seeders.UserSeed()
	userJSON, _ := json.Marshal(model)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) && err == nil {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Condition(t, func() bool {
			var dat models.UserResponse
			var b []byte = rec.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.ID == model.ID && dat.Data.Email == model.Email && dat.Data.PhoneNumber == model.PhoneNumber && dat.Data.Username == model.Username && dat.Data.Name == model.Name && dat.Data.Role == model.Role {
				return true
			}

			return false
		}, rec.Body.String())
	}
}

func TestLoginUser(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	seeders.ItemSeed()
	model, err := seeders.UserSeed()

	config.DB.Create(&model)
	userJSON, _ := json.Marshal(model)
	// userResponse := models.UserResponse{
	// 	Code:    201,
	// 	Status:  "success",
	// 	Message: "user account created",
	// 	Data:    model,
	// }

	//cartLists := createModelCartLists()
	//items := createModelItem()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginUserController(c)) && err == nil {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Condition(t, func() bool {
			var dat models.UserResponse
			var b []byte = rec.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.ID == model.ID && len(dat.Data.Token) > 10 {
				return true
			}

			return false
		}, rec.Body.String())
	}
}
