package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"retailStore/config"
	"retailStore/lib/seeders"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	model, _ := seeders.ItemSeed()
	jsonString, _ := json.Marshal(model)
	Item1Json, _ := json.Marshal(model[0])

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonString)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/items/:id")
	c.SetParamNames("1")

	//Assertions
	if assert.NoError(t, GetItemWIthParamsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, Item1Json, rec.Body.String())
	}
}
