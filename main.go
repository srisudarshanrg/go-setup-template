package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/go-setup-template/server/database"
	"github.com/srisudarshanrg/go-setup-template/server/functions"
	"github.com/srisudarshanrg/go-setup-template/server/setup"
	"github.com/srisudarshanrg/go-setup-template/server/validations"
)

const portNumber = ":8500"

var session *scs.SessionManager

func main() {
	// session
	session = scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// database setup
	db, err := database.CreateDatabaseConn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// database and access
	setup.DBAccessHandlers(db)
	setup.SessionAccessHandlers(session)
	functions.DBAccessFunctions(db)
	validations.DBAccessFormValidations(db)

	// routes
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("server running on port number", portNumber)
	server.ListenAndServe()
}

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(SessionLoadAndSave)

	mux.Get("/", setup.Home)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func SessionLoadAndSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
