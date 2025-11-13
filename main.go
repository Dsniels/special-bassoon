package main

import (
	"log/slog"
	"net/http"
	"time"

	"fixable.com/fixable/internal/app"
	"fixable.com/fixable/internal/data"
	"fixable.com/fixable/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load() 
	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	flag := make(chan struct{})
	
	data.SeedData(app.Db)
	routes := router.InitRoutes(app)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      routes,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	}

	slog.Info("server running", slog.String("page", "http://localhost:8000"), slog.String("admin page", "http://localhost:8000/admin/servicios"))
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			panic(err)
		}
		flag <- struct{}{}
	}()

	<-flag

}
