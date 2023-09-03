package cmd

import (
	"github.com/Shemistan/Lesson_6/api"
	"github.com/Shemistan/Lesson_6/services"
	"github.com/Shemistan/Lesson_6/storage"
)

const (
	host string = "127.0.0.1"
	port int    = 3232
	ttl  int    = 30
)

func main() {
	app := initApp()

	runCommands(&app)
}

func runCommands(api *api.IApi) {

}

func initApp() api.IApi {
	conn := storage.NewConn()
	appStorage := storage.New(host, port, ttl, conn)
	appService := services.New(appStorage)

	return api.New(appService)
}
