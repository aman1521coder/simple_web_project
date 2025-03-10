package main

import (
	"fmt"
	"net/http"

	
)
func(app *application)logError (r *http.Request, err error){
	app.log.Print(err)
}
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request,status int ,message interface{}){
	env:=envelope{"error":message}
	err:=app.writeJSON(w,status,env,nil)
	if err !=nil{
		app.logError(r,err)
		w.WriteHeader(500)
	}
}
func (app *application)serverErrorResponse(w http.ResponseWriter, r *http.Request,err error){
app.logError(r,err)
message:="the server encountred a problem and could not process your request"
app.errorResponse(w,r,http.StatusInternalServerError,message)
}
func (app *application)notFoundResponse(w http.ResponseWriter, r *http.Request){
	message:="the requested resource could not be found "
	app.errorResponse(w,r,http.StatusNotFound,message)
}
func (app *application)methodNotAllowedResponse(w http.ResponseWriter, r *http.Request){
	message :=fmt.Sprintf(" this %s is not allowed  ")
	app.errorResponse(w,r,http.StatusMethodNotAllowed,message)
}
func (app *application)badRequest(w http.ResponseWriter, r *http.Request,err error){
app.errorResponse(w,r,http.StatusBadRequest,err.Error())
}
func (app *application)faildValidationResponse(w http.ResponseWriter, r *http.Request, err map[string]string){
app.errorResponse(w,r,http.StatusUnprocessableEntity,err)
}