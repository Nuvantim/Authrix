package main

import(
	"api/internal/server/http"
	"log"
)

func main(){
	serv := server.ServerGo()

	var envServ, err = config.GetServerConfig()
	if err != nil{
		log.Fatal(err)
	}

	done := make(chan bool, 1)

	go func() {
		serv.Listen(":"+envServ.Port)
	}()

	config.gracefulShutdown(serv, done)

	<-done

}