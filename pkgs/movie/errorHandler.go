package movie

import "log"


func HandleErrorByPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}


func HandleErrorByReturn(err error) error{
	if err != nil {
		return err
	}
	return nil
}


func HandleErrorByPrint(err error) {
	if err != nil {
		log.Println(err)
	}
}