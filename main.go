package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unit-converter/utils"
	"unit-converter/views"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Views
	r.Get("/", templ.Handler(views.Home()).ServeHTTP)
	r.Get("/length", templ.Handler(views.LengthTab()).ServeHTTP)
	r.Get("/weight", templ.Handler(views.WeightTab()).ServeHTTP)
	r.Get("/temperature", templ.Handler(views.TemperatureTab()).ServeHTTP)

	// Unit converter route
	r.Post("/convert", func(w http.ResponseWriter, r *http.Request) {
		amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
		fromUnitStr := r.FormValue("fromUnit")
		toUnitStr := r.FormValue("toUnit")
		result, err := utils.ConvertUnits(amount, fromUnitStr, toUnitStr)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprintf(w, "%v %v = %v", amount, fromUnitStr, result)
	})

	http.ListenAndServe(":8080", r)
}
