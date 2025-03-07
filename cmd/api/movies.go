package main

import (
	
	"fmt"
	
	"net/http"
	"time"

	"github.com/aman1521coder/simple_project/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Year int32 `json:"year"`
		Runtime int32 `json:"runtime"`
	
		Genres []string `json:"genres"`
}
err:=app.readJSON(w,r,&input)
if err!=nil{
app.badRequest(w,r,err)
}
fmt.Fprintf(w, "%+v\n", input)

}


func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
id,err:=app.readIdParam(r)
if err!=nil{
	app.notFoundResponse(w,r)
	return
} 
fmt.Println(id)
movie:=data.Movie{
	Id: id,
	CreatedAt: time.Now(),
	Title: "amans life",
	Runtime: 20,
	Genere: []string{"drama","romance","war"},
	Version:1,

}
err=app.writeJSON(w,http.StatusOK,envelope{"movie":movie},nil)
if err!=nil{
	app.log.Fatal(err)
	app.serverErrorResponse(w,r,err)
	
}
}
