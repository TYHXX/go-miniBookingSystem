package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TYHXX/go-miniBookingSystem/internal/config"
	handerls "github.com/TYHXX/go-miniBookingSystem/internal/handlers"
	"github.com/TYHXX/go-miniBookingSystem/internal/models"
	"github.com/TYHXX/go-miniBookingSystem/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handerls.NewRepo(&app)
	handerls.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handerls.Repo.Home)
	// http.HandleFunc("/about", handerls.Repo.About)
	// http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting appliaction on port %s", portNumber))

	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal((err))
}

func addValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 0.0

	f, err := devideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))

}

func devideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cant devide by zero")
		return 0, err
	}
	result := x / y

	return result, nil
}
