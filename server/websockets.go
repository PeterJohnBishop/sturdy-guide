package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSEvent struct {
	Event  string `json:"event"`
	Data   string `json:"data"`
	Sender string `json:"sender"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow any origin (for dev)
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	ID   string // user ID, session token, etc.
}

var clients = make(map[*Client]bool)

var broadcast = make(chan WSEvent)

const (
	pingPeriod = 30 * time.Second
	pongWait   = 60 * time.Second
	writeWait  = 10 * time.Second
)

func generateID() string {
	return uuid.New().String()
}

func handleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	newClient := &Client{
		Conn: ws,
		ID:   "ws_" + generateID(),
	}
	clients[newClient] = true
	log.Printf("Client %s connected: %s\n", newClient.ID, newClient.Conn.RemoteAddr())

	newClient.Conn.SetReadLimit(512)
	newClient.Conn.SetReadDeadline(time.Now().Add(pongWait))
	newClient.Conn.SetPongHandler(func(string) error {
		newClient.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	go pingClient(newClient.Conn)

	for {
		var event WSEvent
		event.Sender = newClient.ID
		err := newClient.Conn.ReadJSON(&event)
		if err != nil {
			log.Println("WebSocket read error:", err)
			delete(clients, newClient)
			break
		}
		broadcast <- event
	}
}

func pingClient(ws *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for range ticker.C {
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Println("Ping failed, closing connection:", err)
			ws.Close()
			break
		}
	}
}

func HandleMessages() {

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("WebSocket server ready for events at ws://localhost:8080/ws\n")
	for {
		event := <-broadcast

		switch event.Event {
		case "message":
			broadcastMessage(event)
		case "connect":
			handleConnection(event)
		case "disconnect":
			handleDisconnection(event)
		default:
			log.Printf("Unhandled event type: %s\n", event.Event)
		}
	}
}

func broadcastMessage(event WSEvent) {
	for client := range clients {
		log.Printf("message: %s\n", event.Data)
		event.Data = "Message sent." // Ensure data is set for each client
		err := client.Conn.WriteJSON(event)
		if err != nil {
			log.Println("WebSocket write error:", err)
			client.Conn.Close()
			delete(clients, client)
		}
	}
}

func handleConnection(event WSEvent) {
	log.Printf("%s.\n", event.Data)
}

func handleDisconnection(event WSEvent) {
	log.Printf("%s disconnected.\n", event.Data)
}

func WebSocketRoutes(r *gin.Engine) {
	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c)
	})
}
