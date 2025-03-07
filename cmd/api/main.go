package main

import (
	"fmt"
	"net/http"
	"os"

	
	 "time"
	"flag"
	
	"log"
)

// verision of our application

const version="1.0.0"

type config  struct{
 port int
 env string
}
type application struct{
	config config
	log *log.Logger
}
func main() {
	var conf config
	flag.IntVar(&conf.port,"port",4000,"port number for our port")
	flag.StringVar(&conf.env,"env","development", "Environment (development|staging|production)")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	app:=&application{
		config:conf,

		log:logger,
	}

	srv:=&http.Server{
Addr: fmt.Sprintf(":%d", conf.port),
Handler: app.routes(),
IdleTimeout: time.Minute,
ReadTimeout: 10 * time.Second,
WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", conf.env, srv.Addr)
err := srv.ListenAndServe()
if err !=nil{
	logger.Fatal(err)
}


}