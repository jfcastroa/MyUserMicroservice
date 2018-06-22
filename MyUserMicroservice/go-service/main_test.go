package main

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"github.com/stretchr/testify/assert"
	"net/http"
)

var (
	mockDB = map[string]*User{
		"jon@labstack.com": &User{"Jon Snow","juanfercas2002@gmail.com", createHash("1234"), "true","32465366","Colombia","Bogota","Cra 4254 # 104 -56"},
	}
	userJSON = `{"Nombre":"Jon Snow","Email":"juanfercas2002@gmail.com","Password":"` + createHash("1234") + `","Verificado":"true","NoTel":"32465366","Pais":"Colombia","Ciudad":"Bogota","Direccion":"Cra 4254 # 104 -56"}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}



