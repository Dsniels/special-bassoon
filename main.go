package main

import (
	"log/slog"
	"net/http"
	"time"

	"fixable.com/fixable/internal/app"
	"fixable.com/fixable/internal/data"
	"fixable.com/fixable/internal/router"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		app.Logger.Printf("error newApp: %v", err)
		panic(err)
	}
	flag := make(chan struct{})
	app.Logger.Printf("Migrando db")
	data.SeedData(app.Db)
	routes := router.InitRoutes(app)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      routes,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	}

	slog.Info("server running", slog.String("URL", "http://localhost:8000"))
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic(err)
		}
		flag <- struct{}{}
	}()

	<-flag

}
