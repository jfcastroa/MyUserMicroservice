package main


import (

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"encoding/json"
	"fmt"
	"encoding/hex"
	"crypto/md5"
	"flag"
	"time"
	"go-service/Broker/Consumer"
	"go-service/Broker/Producer"
	"log"
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

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 5*time.Second, "lifetime of process before shutdown (0s=infinite)")
)

var (

	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")

	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar2", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
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



	if err := Producer.publish(*uri, *exchangeName, *exchangeType, *routingKey, jsonValue, *reliable); err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("published %dB OK", len(jsonValue))
	consumer, err := Consumer.NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if *lifetime > 0 {
		log.Printf("running for %s", *lifetime)
		time.Sleep(*lifetime)
	} else {
		log.Printf("running forever")
		select {}
	}

	log.Printf("shutting down")

	if err := consumer.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
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