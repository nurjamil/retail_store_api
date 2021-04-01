package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"retailStore/config"
	"retailStore/lib/db"
	"retailStore/lib/seeders"
	"retailStore/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func createModelCartLists() []models.ShoppingCartList {
	cartLists := make([]models.ShoppingCartList, 1)
	cartLists[0] = models.ShoppingCartList{
		ItemID:   1,
		Quantity: 3,
	}
	return cartLists
}

// func createModelItem() []models.Item {
// 	items := make([]models.Item, 1)
// 	items[0] = models.Item{
// 		ID:             1,
// 		ItemCategoryID: 2,
// 	}
// 	return items
// }

func TestPostItemToShoppingCart(t *testing.T) {
	// Setup
	config.InitDBTest()
	//config.DropTable() //reset tables
	config.InitialMigration()

	// seeders for insert categories, paymentservices, and couries. for dev purposes
	seeders.Seed()
	seeders.ItemSeed()
	model, err := seeders.UserSeed()

	config.DB.Create(&model)
	model.ShoppingCart.UserID = model.ID
	config.DB.Create(&model.ShoppingCart)

	db.LoginUser(&model)
	cartLists := createModelCartLists()
	cartLists[0].ShoppingCartID = model.ShoppingCart.ID
	//items := createModelItem()
	itemsJSON, _ := json.Marshal(cartLists[0])

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(itemsJSON)))
	req.Header.Set("Authorization", "Bearer "+model.Token)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println(c.Request().Header)
	//Assertions
	if assert.NoError(t, PostItemToShoppingCartController(c)) && err == nil {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Condition(t, func() bool {
			var dat models.ShoppingCartAPI
			var b []byte = rec.Body.Bytes()
			if err := json.Unmarshal(b, &dat); err != nil {
				return false
			}
			if dat.Data.ShoppingCartList[0].ItemID == cartLists[0].ItemID && dat.Data.ShoppingCartList[0].Quantity == cartLists[0].Quantity && dat.Data.ShoppingCartList[0].ShoppingCartID == cartLists[0].ShoppingCartID {
				return true
			}

			return false
		}, rec.Body.String())
	}
}
