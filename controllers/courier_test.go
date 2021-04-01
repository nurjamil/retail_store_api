package controllers

import (
	"retailStore/config"
	"testing"

	"github.com/labstack/echo"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetCourierControllers(t *testing.T) {
	// var testCases = []struct {
	// 	name                 string
	// 	path                 string
	// 	expectStatus         int
	// 	expectBodyStartsWith string
	// }{
	// 	{
	// 		name:                 "berhasil",
	// 		path:                 "/couriers",
	// 		expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
	// 		expectStatus:         http.StatusOK,
	// 	},
	// }

	// e := InitEcho()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// rec := httptest.NewRecorder()
	// e = e.NewContext(req, rec)

	// for _, testCase := range testCases {
	// 	c.SetPath(textCase.path)

	// 	// Assertions
	// 	if assert.NoError(t, GetCouriersController(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 		body := rec.Body.String()
	// 		// assert.Equal(t, userJSON, rec.Body.String())
	// 		assert.True(t, strings.HasPrefix(body,
	// 			testCase.expectBodyStartsWith))
	// 	}
	// }
}
