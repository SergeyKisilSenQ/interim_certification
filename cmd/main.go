package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handlers "interim_certification/pkg/api/handler"
	"interim_certification/pkg/storage/repo_func"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	Cs := repo_func.NewCityStorage()
	repo_func.ReadFile(Cs)

	go func() {
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Get("/city/{cityId_}", handlers.GetCity(Cs))
		r.Post("/create", handlers.AddCity(Cs))
		r.Delete("/delete", handlers.DeleteCity(Cs))
		r.Put("/{cityId_}", handlers.UpdateCityPopulation(Cs))
		r.Get("/region", handlers.GetRegion(Cs))
		r.Get("/district", handlers.GetDistrict(Cs))
		r.Get("/population", handlers.GetPopulation(Cs))
		r.Get("/foundation", handlers.GetFoundation(Cs))
		http.ListenAndServe("localhost:8080", r)
	}()

	go func() {
		defer wg.Done()
		<-sig

		file, err := os.OpenFile("cities.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		for i, _ := range Cs {
			var record []string
			record = append(record, strconv.Itoa(Cs[i].Id), Cs[i].Name, Cs[i].Region, Cs[i].District, strconv.Itoa(Cs[i].Population), strconv.Itoa(Cs[i].Foundation))
			err := writer.Write(record)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
		writer.Flush()
		fmt.Println("Выхожу из программы")
	}()
	wg.Wait()
}
