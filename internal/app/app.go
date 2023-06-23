package app

import (
	"encoding/json"
	"io/ioutil"
	"log"

	tgClient "github.com/NickolayGoloburdin/tlg_bot_golang/internal/clients/telegram"
	event_consumer "github.com/NickolayGoloburdin/tlg_bot_golang/internal/consumer/event-consumer"
	"github.com/NickolayGoloburdin/tlg_bot_golang/internal/events/telegram"
	files "github.com/NickolayGoloburdin/tlg_bot_golang/internal/storage/filestorage"
)

type Config struct {
	token string
}
type App struct {
	TgBotHost         string `json:"host"`
	SqliteStoragePath string `json:"sqladdress"`
	BatchSize         int    `json:"batchsize"`
	Token             string `json:"token"`
}

func NewApp(settings string) (c App) {
	content, err := ioutil.ReadFile(settings)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c

}
func (a App) Start() {

	// s, err := sqlite.New(a.SqliteStoragePath)
	// if err != nil {
	// 	log.Fatal("can't connect to storage: ", err)
	// }

	// if err := s.Init(context.TODO()); err != nil {
	// 	log.Fatal("can't init storage: ", err)
	// }
	filestorage := files.New(a.SqliteStoragePath)

	eventsProcessor := telegram.New(
		tgClient.New(a.TgBotHost, a.Token),
		filestorage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, a.BatchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
