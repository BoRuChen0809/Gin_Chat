package global

import (
	model "gin_chat/model"
	"log"

	"github.com/olahol/melody"
)

var Melody *melody.Melody

func SetupMelody() {
	Melody = melody.New()

	Melody.HandleConnect(func(s *melody.Session) {
		id := s.Request.URL.Query().Get("id")
		name := s.Request.URL.Query().Get("name")
		msg := model.NewConnectMessage(id, name)
		err := Melody.Broadcast(msg.GetJSON())
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
	})

	Melody.HandleMessage(func(s *melody.Session, data []byte) {
		msg, err := model.ParseMessage(data)
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
		err = Melody.Broadcast(msg.GetJSON())
		if err != nil {
			log.Printf("Melody Connect ERROR : %v\n", err)
		}
	})

	Melody.HandleClose(func(s *melody.Session, i int, s2 string) error {
		id := s.Request.URL.Query().Get("id")
		name := s.Request.URL.Query().Get("name")
		err := Melody.Broadcast(model.NewCloseMessage(id, name).GetJSON())
		if err != nil {
			log.Printf("Melody Close ERROR : %v\n", err)
		}
		return nil
	})
}
