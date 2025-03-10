package main

import (
	"fmt"
	"net/http"
	"os"
	"context" // New import
"database/sql" // New import

	
	 "time"
	"flag"
	
	"log"
	_ "github.com/lib/pq"
)

// verision of our application

const version="1.0.0"

type config  struct{
 port int
 env string
 db struct {
	dsn string
 }
}
type application struct{
	config config
	log *log.Logger
}

func main() {
	var conf config
	flag.IntVar(&conf.port,"port",4000,"port number for our port")
	flag.StringVar(&conf.env,"env","development", "Environment (development|staging|production)")
flag.StringVar(&conf.db.dsn ,"db-dsn",os.Getenv("MOVIE_GO"),"PostgreSQL DSN")
	flag.Parse()
	
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	db ,err:=openDb(conf)
	if err!=nil{
		log.Fatal(err)
	}

	defer db.Close()
	logger.Printf("database connection pool established ")
	
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
err = srv.ListenAndServe()
if err !=nil{
	logger.Fatal(err)
}


}
func openDb(conf config)(*sql.DB,error){
	db ,err:=sql.Open("postgres",conf.db.dsn)
	if err !=nil{
		return nil ,err
	}
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err=db.PingContext(ctx)
	if err!=nil{
 return nil,err

	}
	return db,nil
}