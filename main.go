package main

import (
	"gin_chat/global"
	"gin_chat/router"
)

func main() {
	global.SetupMelody()
	r := router.NewRouter()
	r.Run()
}

/*
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/html/*")
	r.Static("/assets", "./template/assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chatroom.html", nil)
	})

	m := melody.New()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		msg, err := my_pkg.ParseMessage(data)
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
		log.Println(msg)
		err = m.Broadcast(data)
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		id := s.Request.URL.Query().Get("id")
		name := s.Request.URL.Query().Get("name")
		msg := my_pkg.NewConnectMessage(id, name)
		log.Println(msg)
		err := m.Broadcast(msg.GetJSON())
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
	})

	m.HandleClose(func(s *melody.Session, i int, str string) error {
		id := s.Request.URL.Query().Get("id")
		name := s.Request.URL.Query().Get("name")
		err := m.Broadcast(my_pkg.NewCloseMessage(id, name).GetJSON())
		if err != nil {
			log.Printf("Melody Close ERROR : %v\n", err)
		}
		return nil
	})

	r.Run()
}
*/
