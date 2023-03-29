package websocket

import (
	"fmt"
)

type WSHub struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan map[string]interface{}
}

func NewWSHub() *WSHub {
    return &WSHub{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan map[string]interface{}),
    }
}

func (hub *WSHub) Start() { 
    //Close all active client connections when closing the hub
    defer func() {
        for client,_ := range hub.Clients{
            hub.Unregister <- client
            client.Conn.Close()
        }
    }()

    for {
        select {
        case client := <-hub.Register:
            hub.Clients[client] = true
            fmt.Println("WS: Size of Connection WSHub: ", len(hub.Clients))
            break
        case client := <-hub.Unregister:
            delete(hub.Clients, client)
            fmt.Println("WS: Size of Connection WSHub: ", len(hub.Clients))
            break
        case message := <-hub.Broadcast:
            for client, _ := range hub.Clients {
                //implement selective sending logic
                if client.CrId == message["CrId"]{
                    if err := client.Conn.WriteJSON(message); err != nil {
                        return
                    }
                    fmt.Printf("WS: Sent message to %s\n", client.CrId)
                }
            }
        }
    }

   
}