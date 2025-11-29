package server

import (
	"log"
	// "sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeGin() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.Config{
		// AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	config.AllowOriginFunc = func(origin string) bool {
		if origin == "null" {
			log.Printf("CORS Accepted: null origin for file://\n")
			return true
		}
		if len(origin) >= 17 && (origin[:17] == "http://localhost:" || origin == "http://localhost") {
			log.Printf("CORS Accepted: %s\n", origin)
			return true
		}
		if len(origin) >= 15 && (origin[:15] == "http://127.0.0.1:" || origin == "http://127.0.0.1") {
			log.Printf("CORS Accepted: %s\n", origin)
			return true
		}

		log.Printf("CORS Rejected: %s\n", origin)
		return false
	}

	r.Use(cors.New(config))

	// var wg sync.WaitGroup
	// var errChan = make(chan error, 3)

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	log.Println("Go routine 1")
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	log.Println("Go routine 2")

	// 	//
	// }()

	// wg.Wait()
	// close(errChan)

	// for err := range errChan {
	// 	if err != nil {
	// 		log.Fatalf("Initialization failed: %v", err)
	// 	}
	// }
	WebSocketRoutes(r)

	go func() {
		log.Println("Starting background message handler.")
		HandleMessages()
	}()

	log.Println("Now serving Gin.")
	r.Run(":8080")
}
