package server

import (
	"log"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeGin() {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:*"}, // development only
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	config.AllowOriginFunc = func(origin string) bool {
		return len(origin) > 0 && (origin == "http://localhost" ||
			(len(origin) > 17 && origin[:17] == "http://localhost:"))
	}

	r.Use(cors.New(config))

	var wg sync.WaitGroup
	var errChan = make(chan error, 3)

	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Println("Go routine 1")

		//
	}()

	go func() {
		defer wg.Done()
		log.Println("Go routine 2")

		//
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			log.Fatalf("Initialization failed: %v", err)
		}
	}

	log.Println("Now serving Gin.")
	r.Run(":8080")
}
