package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"retailStore/config"
	"retailStore/lib/seeders"
	"retailStore/models"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetItemWithParamsId(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	model, _ := seeders.ItemSeed()
	for _, each := range model.Data {
		fmt.Println(each.ID)
	}

	//id 1 json
	newModel := models.ItemResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting items",
		Data:    make([]models.Item, 1),
	}
	newModel.Data[0] = model.Data[0]
	Item1Json, _ := json.Marshal(newModel)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/items/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//Assertions
	if assert.NoError(t, GetItemWIthParamsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(Item1Json)+"\n", rec.Body.String())
	}

}

func TestGetItemWithCategory(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	model, _ := seeders.ItemSeed()
	for _, each := range model.Data {
		fmt.Println(each.ID)
	}

	//category books json
	modelCategory2 := models.ItemResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting items",
		Data:    make([]models.Item, 2),
	}
	modelCategory2.Data[0] = model.Data[0]
	modelCategory2.Data[1] = model.Data[1]

	itemCategoryBooksJson, _ := json.Marshal(modelCategory2)

	e := echo.New()
	q := make(url.Values)
	q.Set("category_name", "Books")
	req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetItemController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(itemCategoryBooksJson)+"\n", rec.Body.String())
	}
}
