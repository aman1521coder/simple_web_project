package data

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)
type Runtime int32
var ErrorInvalidformat error=errors.New("invalid runtime format")
func (r Runtime)MarshalJSON()([]byte,error){
	jsonValue:=fmt.Sprintf("%d mins",r)
	qoutedJsonValue:=strconv.Quote(jsonValue)
	return []byte(qoutedJsonValue),nil
}
func (r *Runtime)UnmarshalJSON(jsonValue []byte) error{

unquoteValue,err:=strconv.Unquote(string(jsonValue))
if err!=nil{
	return ErrorInvalidformat
}
parts:=strings.Split(unquoteValue,"")
if len(parts)!=2&&parts[1]!="minus"{
	return ErrorInvalidformat
}
i,err:=strconv.ParseInt(parts[0],10,32)
if err!=nil{
	return ErrorInvalidformat
}
*r=Runtime(i)
return nil

}