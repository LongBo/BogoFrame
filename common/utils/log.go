package utils

import (
	"log"
	"os"
	"time"
)

const dir = "bin/log"

var Nlog *log.Logger

func init() {

	var fh *os.File
	var err error

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Println(err)
		}
	}
	file := dir + "/" + time.Now().Format("20060102") + ".log"

	fh, err = os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)

	if err != nil {
		log.Println(err)
	}

	Nlog = log.New(fh, "LOG:", log.Ldate|log.Ltime|log.Lshortfile)
}
