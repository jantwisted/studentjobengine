package ops

import (
	"log"
)

func  LogPrint(str string){
	log.Print(str)
}

func LogError(err error){
	if err != nil{
		log.Fatal(err)
	}
}
