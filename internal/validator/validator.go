package validator

import (
	
"regexp"
)
var (
EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
type Validator struct{
	Erorrs map[string]string
}
func New() *Validator{
	return &Validator{Erorrs: make(map[string]string)}
}
func (v *Validator)Valid()bool{
	return len(v.Erorrs)==0
}
func(v *Validator)AddError(key string,message string){
	if _,exists:=v.Erorrs[key];!exists{
	v.Erorrs[key]=message
	}

}
func (v *Validator)Check(ok bool,key string ,message string){
	if !ok{
		v.AddError(key,message)
	}
}
func In(value string,list ...string) bool{
	for i:=range list{
		if list[i]==value {
		return true
		}

	}
	return false
}
func Match (value string ,regx *regexp.Regexp) bool{
	return regx.MatchString(value)
}
func Unique(values []string)bool{
	uniqueValues := make(map[string]bool)
	for _, value := range values {
	uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}