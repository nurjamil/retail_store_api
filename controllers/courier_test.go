package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"retailStore/config"
	"retailStore/lib/seeders"
	"retailStore/models"
	"testing"
	"strings"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()

	return e
}

func TestCreateCourier(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	model, _ := seeders.CouriersSeed()
	
	seeders.CouriersSeed()
	model, err := seeders.CouriersSeed()
	couriersJSON, _ := json.Marshal(model)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(couriersJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateCourierController(c)) && err == nil {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Condition(t, func() bool {
			var dat models.Courier
			var b []byte = rec.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			} else {
				return true
			}
		}, rec.Body.String())
	}
}