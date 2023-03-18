package websocket

import (
    "fmt"
    "strconv"
    "net/http"

	"github.com/gin-gonic/gin"
    
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

var Hub *WSHub

func serveWs(c *gin.Context) {
    CrId := c.Param("CrId")
    unixtime, _ := strconv.ParseInt(c.Param("unixtime"),10 ,64)

    //Creating a new client for the entered route /CgWs/:CrId
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        fmt.Println("WS: error in establishing new connection")
        return
    }

    client := &Client{
        Conn: conn,
        Hub: Hub,
        CrId: CrId,
        Unixtime: unixtime,
    }

    Hub.Register <- client
    client.KeepAlive()
}

func InitWS(router *gin.Engine) {
	Hub = NewWSHub()
	go Hub.Start()

	router.GET("/CgWS/:CrId/:unixtime", serveWs)
}