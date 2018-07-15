package crypto

import (
	"strconv"
)

func IDEncrypt(id int64,key string) (string){
	strID := strconv.FormatInt(id,10)
	return AesEncrypt(strID,key)
}
func IDDecrypt(id,key string) (int64){

	strID := AesDecrypt(id,key)
	intID, err := strconv.ParseInt(strID, 10, 64)
	if(err != nil){
		panic(err)
	}
	return intID
}