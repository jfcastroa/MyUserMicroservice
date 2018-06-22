package main


import (

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/hex"
	"crypto/md5"

)

type (

	User struct {
		//ID   int    `json:"ID"`
		Nombre string `json:"Nombre"`
		Email string `json:"Email"`
		Password string `json:"Password"`
		Verificado string `json:"Verificado"`
		NoTel string `json:"NoTel"`
		Pais string `json:"Pais"`
		Ciudad string `json:"Ciudad"`
		Direccion string `json:"Direccion"`
	}
	handler struct {
		db map[string]*User
	}

)



//create user
func (h *handler) createUser(c echo.Context) error {

//var name =c.Request().Form.Get("Nombre")

	user := new(User)
	err := c.Bind(user)

	if  err != nil {
		return err
	}



	//jsonData := map[string]string{"Nombre": "Nic", "Email": "juanfercas2002@gmail.com", "Password": createHash("1234"),"Verificado": "true","NoTel":"32465366","Pais":"Colombia","Ciudad":"Bogota","Direccion":"Cra 4254 # 104 -56"}
	jsonValue, _ := json.Marshal(user)

	response, err := http.Post("http://localhost:50483/api/Users", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")

	return c.JSON(http.StatusCreated, user)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}





func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler

	//e.POST("/users", h.createUser)
	//e.GET("/users/:id", getUser)
	// Server

	// Start server
	e.Logger.Fatal(e.Start(":1323"))


}