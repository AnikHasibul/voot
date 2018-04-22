// AUTHOR: http://github.com/AnikHasibul/
// LICENSE: No license

// #VooT DB
//
// EXAMPLE::
//
// ===========
//
//
// package main
//
//import "voot"
//
// // Declaring Global variable `Key`(key to store) as database to store
//
// var Key = "Hello VooT"
//
//
//
// // Intializing DB to store the variable `Key`
//
// var DB  = voot.NewDB(&voot.VooT{Name:"DBname",Data:&Key})
//
//
//func main {
//	defer DB.SaveAndClose()
//	//^^^DON'T FORGET TO SAVE AND CLOSE YOUR DB AFTER YOUR PROGRAM RETURNS
//
//	// All or your codes go here...
//
//	fmt.Println("Hello VooT!")
//	return
//}
//
//
//
//
//More example: $GOPATH/src/voot/_example/
//
//https://github.com/AnikHasibul/voot/ {_example}
//
//==============
package voot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

type VooT struct {
	Name string      //Database name
	Data interface{} //Key to store (&Key)
}

var hEALTH = true
var bUFF = true
var done = make(chan bool)
var quit = make(chan os.Signal, 1)
var Healthz = &hEALTH // Returns the database connection closed or not `true = connected`
var sv, cv bool

const tr = true

func (v VooT) Health() bool {
	return hEALTH
}
func (v VooT) Kill() {
	hEALTH = false
}

func NewDB(Db *VooT) *VooT {
	go Db.saveVoot()
	signal.Notify(quit, os.Interrupt)
	go Db.closeVoot()
	for {
		if sv && cv {
			return Db
		} else {
			time.Sleep(10 * time.Millisecond)
		}

	}
	return Db
}

var i int

func (v VooT) saveVoot() {
	sv = true
	for {
		if !hEALTH {
			log.Println("Database has been closed!")
			log.Println("Rejecting the last buffer!")
			break
		}
		JsonoByto, err := json.Marshal(v.Data)
		if err != nil {
			err := ioutil.WriteFile("./"+v.Name+".VooT.Err", []byte(sfmtVooT(v.Data)), 0666)
			if err != nil {
				log.Println("ERR:", err)
			}
			log.Println("ERR[EncOnSave]:", err)
		} else {
			bUFF = true
			err := ioutil.WriteFile("./"+v.Name+".VooT", JsonoByto, 0666)
			if err != nil {
				log.Println("ERR:[SaveOnBuffer]", err)
				time.Sleep(5000 * time.Millisecond)
			}
			bUFF = false
			time.Sleep(3000 * time.Millisecond)
		}

	}
	return

}

func (v VooT) closeVoot() {
	cv = true
	<-quit
	if !hEALTH {
		return
	}
	hEALTH = false
	log.Println("Signal from VooT :: Shutdown!")
	log.Println("Closing Database...")
	JsonoByto, err := json.Marshal(v.Data)
	if err != nil {
		err := ioutil.WriteFile("./"+v.Name+".VooT.Err", []byte(sfmtVooT(v.Data)), 0666)
		if err != nil {
			log.Println("ERR:", err)
		}
		log.Println("ERR[EncOnSave]:", err)
		os.Exit(0)
		return
	}
	defer func() {
		for bUFF {
			log.Println("Waiting For Buffer!")
			time.Sleep(1000 * time.Millisecond)
		}
		log.Println("Leaving in 3 seconds...")
		time.Sleep(3000 * time.Millisecond)
		err := ioutil.WriteFile("./"+v.Name+".VooT", JsonoByto, 0666)
		if err != nil {
			log.Println("ERR:", err)
		}
		log.Println("That's all I know!")
		log.Println("Bye!")
		os.Exit(0)
	}()
	os.Exit(0)
	close(done)
}

func (v VooT) SaveAndClose(exit bool) {
	hEALTH = false
	if exit {
		log.Println("Signal from VooT :: SaveAndClose! [SHUTDOWN]")
	} else {
		log.Println("Signal from VooT :: SaveAndClose!")
	}
	JsonoByto, err := json.Marshal(v.Data)
	if err != nil {
		err := ioutil.WriteFile("./"+v.Name+".VooT.Err", []byte(sfmtVooT(v.Data)), 0666)
		if err != nil {
			log.Println("ERR:", err)
		}
		log.Println("ERR[EncOnSave]:", err)
		if exit {
			os.Exit(0)
		}
		return
	}
	defer func() {
		for bUFF {
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(300 * time.Millisecond)
		err := ioutil.WriteFile("./"+v.Name+".VooT", JsonoByto, 0666)
		if err != nil {
			log.Println("ERR:", err)
		}
		if exit {
			os.Exit(0)
		}
	}()
}

func sfmtVooT(args ...interface{}) string {
	var argsStr = []string{}

	for _, v := range args {
		if _, ok := v.(error); ok {
			argsStr = append(argsStr, fmt.Sprintf("%v", v))
			continue
		}
		argsStr = append(argsStr, fmt.Sprintf("%#v", v))
	}

	return strings.Join(argsStr, ", ")
}
