package main

import (
	
	"fmt"
	
	"net/http"
	"time"

	"github.com/aman1521coder/simple_project/internal/data"
	"github.com/aman1521coder/simple_project/internal/validator"
)

func  ValidateMovie(v *validator.Validator,movie *data.Movie){
	v.Check(movie.Title != "", "title", "must be provided")
v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
v.Check(movie.Year != 0, "year", "must be provided")
v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
v.Check(movie.Runtime != 0, "runtime", "must be provided")
v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
v.Check(movie.Genere != nil, "genres", "must be provided")
v.Check(len(movie.Genere) >= 1, "genres", "must contain at least 1 genre")
v.Check(len(movie.Genere) <= 5, "genres", "must not contain more than 5 genres")
v.Check(validator.Unique(movie.Genere), "genres", "must not contain duplicate values")

}
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	
	v:=validator.New()
	
	var input struct {
		Title string `json:"title"`
		Year int32 `json:"year"`
		Runtime data.Runtime `json:"runtime"`
	
		Genres []string `json:"genres"`
}

err:=app.readJSON(w,r,&input)
if err!=nil{
app.badRequest(w,r,err)
}

movie := &data.Movie{
Title: input.Title,
Year: input.Year,
Runtime: input.Runtime,
Genere: input.Genres,
}
if ValidateMovie(v, movie); !v.Valid() {
	app.faildValidationResponse(w, r, v.Erorrs)
	return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
