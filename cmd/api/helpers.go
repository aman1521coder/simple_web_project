package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
type envelope map[string]interface{}
func (app *application) readIdParam(r * http.Request)(int64,error){
	params:=httprouter.ParamsFromContext(r.Context())
	id,err:=strconv.ParseInt(params.ByName("id"),10,64)
	
	if err!=nil||id<1{
		return 0,errors.New("invalid id paramter")

	}
	return id,nil
	
	

} 
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes:=1_048_576
	 r.Body=http.MaxBytesReader(w,r.Body,int64(maxBytes))
	 dec:=json.NewDecoder(r.Body)
	 dec.DisallowUnknownFields()
		err:=dec.Decode(dst)
		if err!=nil{
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError
			var invalidUnmarshalError *json.InvalidUnmarshalError
			switch  {
			case errors.As(err,&syntaxError):
			 return fmt.Errorf("the body u sent have a bad syntax json  %d ",syntaxError.Offset)
			 if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
				}
				return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
			 
			case errors.Is(err,io.ErrUnexpectedEOF):
				return errors.New("the body contain badly-formed json")	
			case errors.Is(err,io.EOF):
			return errors.New("body must not empty")
			case errors.As(err,&invalidUnmarshalError):
			  panic(err)
			default:
				return err
		
		
		}
		
		

			

		}
		return nil
	
}
func (app *application)writeJSON(w http.ResponseWriter,status int,data envelope,headers http.Header)error{
js,err:=json.MarshalIndent(data,"","\t")
	if err!=nil{
		return err
	}
	js=append(js,'\n')
	for key,value:=range headers{
		w.Header()[key]=value
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil

}