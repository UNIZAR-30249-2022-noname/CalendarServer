package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *TestSuite) setupRouterTest() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	return r
}

func (suite *TestSuite) SetupTest() {
	suite.router = suite.setupRouterTest()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestPingRoute() {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	suite.router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)
	suite.Equal("pong", w.Body.String())
}
