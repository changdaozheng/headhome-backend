package websocket

import (
    "fmt"
    "time"
    "github.com/gorilla/websocket"
)

type Client struct {
    CrId        string
    Conn        *websocket.Conn
    Hub         *WSHub
    Unixtime    int64
}

func (c *Client) KeepAlive() {
    defer func() {
        c.Hub.Unregister <- c
        fmt.Printf("WS: %s disconnected\n", c.CrId)
        c.Conn.Close()
    }()
    
    fmt.Print("WS: ", c.CrId, " connected at ", time.Unix(c.Unixtime, 0), "\n")

    for {
        _, _, err := c.Conn.ReadMessage()
        if err != nil {
            return
        }
    }
}