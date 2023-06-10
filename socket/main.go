package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func writer(ws *websocket.Conn, id string) {
	defer func() {
		ws.Close()
	}()

	for {
		select {}
	}
}

type Connection struct {
	name string
}

var connections = make(map[string][]*Connection)

type Req struct {
	Name string `query:"name"`
}

func hello(c echo.Context) error {
	id := c.Param("id")
	req := &Req{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
	}
	fmt.Println(req)
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		val, ok := connections[id]
		if !ok {
			fmt.Println("There is no room will create")
			connections[id] = make([]*Connection, 0)
			connections[id] = append(connections[id], &Connection{name: req.Name})
		} else {
			fmt.Printf("There is room and there are %d,\n", len(val))
			val = append(val, &Connection{name: req.Name})
			connections[id] = val
		}

		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s has entered Current room user is %d", req.Name, len(connections[id]))))

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		defer func() {
			delete(connections, id)
		}()
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ws/:id", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
