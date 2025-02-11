package main

import (
	"log"
	"net/http"

	"gymshark/packcalculator/internal/handler"
	"gymshark/packcalculator/internal/middleware"
	"gymshark/packcalculator/internal/service"
)

func main() {
	// Initialize dependencies
	calculator := &service.DefaultPackCalculator{}
	packService := service.NewPackService(calculator, []int{250, 500, 1000, 2000, 5000})
	packHandler := handler.NewPackHandler(packService)
	corsHandler := middleware.NewCORSMiddleware(packHandler)

	// Setup routes
	http.Handle("/calculate-packs", corsHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	// Start server
	log.Println("Server starting on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
