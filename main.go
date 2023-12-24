package main

import (
	"go-server-compare/servers"
	"log/slog"
)

func main() {
	slog.Info("Initializing router")
	//         10s test                port            -c 50                   -c 200                  -c 1000
	go servers.NewGin().Init()        //8080       457238 |  8,76MB/s      417839 |  8,01MB/s      381100 | 7,30MB/s
	go servers.NewHttpMux().Init()    //8081       539316 | 10,34MB/s      584001 | 11,19MB/s      470207 | 9,00MB/s
	go servers.NewEcho().Init()       //8082       557608 | 10,69MB/s      509854 |  9,77MB/s      463519 | 8,89MB/s
	go servers.NewMartini().Init()    //8083       434394 |  8,33MB/s      438705 |  8,41MB/s      397508 | 7,62MB/s
	go servers.NewGoji().Init()       //8084       552464 | 10,59MB/s      584054 | 11,19MB/s      450340 | 8,62MB/s
	go servers.NewGorilla().Init()    //8085       554316 | 10,63MB/s      572607 | 10,98MB/s      446838 | 8,56MB/s
	go servers.NewFiber().Init()      //8086       608215 | 11,66MB/s      555507 | 10,65MB/s      498507 | 9,55MB/s
	go servers.NewChi().Init()        //8087       550486 | 10,56MB/s      567778 | 10,88MB/s      485359 | 9,30MB/s
	go servers.NewHttpRouter().Init() //8088       443655 |  8,5MB/s       418705 |  8,02MB/s      368643 | 7,07MB/s
	var forever chan struct{}
	slog.Info("Infinite loop")
	<-forever
}
