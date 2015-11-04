package main

import (
	"github.com/HackerLoop/rotonde-client.go"
	"github.com/HackerLoop/rotonde/shared"
	log "github.com/Sirupsen/logrus"
)

func main() {
	client := client.NewClient("ws://127.0.0.1:4224/")

	client.OnNamedEvent("testevent", func(m interface{}) bool {
		log.Info(m)
		return false
	})

	event := &rotonde.Definition{"goevent", "event", rotonde.FieldDefinitions{}}
	event.PushField("goname", "string", "")
	client.AddLocalDefinition(event)

	client.OnDefinition(func(m interface{}) bool {
		def := m.(rotonde.Definition)
		log.Info("Got definition")
		log.Info(def.Identifier)
		for _, f := range def.Fields {
			log.Info(f.Name, " ", f.Type)
		}
		return true
	})

	client.OnNamedDefinition("testaction", func(m interface{}) bool {
		client.SendAction("testaction", rotonde.Object{
			"field1": 42,
			"field2": "pouet",
			"field3": true,
		})
		return false
	})

	select {}
}
