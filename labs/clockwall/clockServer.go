// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"os"

)

func handleConn(c net.Conn, Time string, channel chan string) {
	defer c.Close()

	for {
		msg := Time+<-channel
		time.Sleep(time.Second)
		io.WriteString(c, msg)
	}
}

func main() {

	var a string
	flag.StringVar(&a)
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	channelA := make(chan string, 1)
	channelB := make(chan string)
	X := os.Getenv("TZ")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		
		go handleConn(conn, channelB, X) // handle connections concurrently
		go clockRun(channelA, channelB)
	}
	close(channelA)
	close(channelB)
}

func clockRun(InChannel chan string, OutChannel chan string){

	for{
		t, err := timeIn(time.Now(), <-InChannel)
		if err != nil{
			return
		}
		OutChannel <- t.Format("12:00:00\n")
	}
}

func TimeRun(t time.Time, name string) (time.Time, error) {
	location, err := time.LoadLocation(name)
	if err == nil {
		t=t.In(loc)
		
	}
	return t, err
}