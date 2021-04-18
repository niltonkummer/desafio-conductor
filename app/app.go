package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/niltonkummer/desafio-conductor/app/setup"

	"github.com/niltonkummer/desafio-conductor/app/config"

	"github.com/niltonkummer/desafio-conductor/app/model"

	"github.com/niltonkummer/desafio-conductor/app/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/niltonkummer/desafio-conductor/docs"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

var instance *app

type app struct {

	// Config
	conf *config.Config

	// Components
	router *chi.Mux
	db     *gorm.DB
	log    zerolog.Logger

	// Repository
	account     repository.Account
	transaction repository.Transaction
}

func NewApplication(config *config.Config) *app {

	instance = &app{
		conf: config,
	}

	instance.log = zerolog.New(os.Stdout)

	err := setupDB(instance)
	if err != nil {
		instance.log.Fatal().Err(err).Str("file", config.DBDialect.Name()).Msg("setupDB")
	}

	err = setupRepository(instance)
	if err != nil {
		instance.log.Fatal().Err(err).Msg("setupRepository")
	}

	setupRoutes(instance)

	return instance
}

func (a *app) StartServer() {

	srv := &http.Server{
		Addr:              a.conf.HttpListen,
		Handler:           a.router,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 10,
		ErrorLog:          log.New(os.Stderr, "", log.LstdFlags),
	}

	idleConnections := make(chan struct{})
	go func() {
		signStop := make(chan os.Signal, 1)
		signal.Notify(signStop, syscall.SIGINT, syscall.SIGTERM)
		<-signStop

		a.log.Info().Msg("[SERVER] Shutting down...")

		// force quit
		go func() {
			<-signStop
			os.Exit(2)
		}()

		// create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		// start http shutdown
		if err := srv.Shutdown(ctx); err != nil {
			a.log.Error().Caller().Err(err).Msg("[SERVER] Shutdown!")
		}

		close(idleConnections)
	}()

	a.log.Info().Msgf("[SERVER] Listen at %s", a.conf.HttpListen)
	err := srv.ListenAndServe()
	if err != nil {
		a.log.Fatal().Err(err).Msg("[SERVER] ListenAndServe")
	}

	a.log.Info().Msg("[SERVER] waiting idle connections...")
	<-idleConnections
	a.log.Info().Msg("[SERVER] Bye")
}

func setupDB(a *app) error {

	db, err := setup.SetupDB(a.conf.DBDialect)
	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func setupRepository(a *app) error {

	err := a.db.AutoMigrate(&model.Account{}, &model.Transaction{})
	if err != nil {
		return err
	}

	a.account = repository.CreateAccountRepository(a.db)
	a.transaction = repository.CreateTransactionRepository(a.db)

	return nil
}

func setupRoutes(a *app) {

	a.router = chi.NewRouter()

	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)

	a.router.Route("/conductor", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/api", func(r chi.Router) {
				r.Mount("/contas", SetupRoutes(a))
			})
		})
	})

	a.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition"
	))

}
