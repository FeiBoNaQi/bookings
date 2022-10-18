package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FeiBoNaQi/bookings/internal/config"
	"github.com/FeiBoNaQi/bookings/internal/handlers"
	"github.com/FeiBoNaQi/bookings/internal/helpers"
	"github.com/FeiBoNaQi/bookings/internal/models"
	"github.com/FeiBoNaQi/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":9981"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("start application on port %s\r\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func Run() error {
	// what am I going to push in the session
	gob.Register(models.Reservation{})

	// change to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create tempalte cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
