package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing...")
	for i := 1; 1 < 10; i++ {
		msg := "Hello" + string(48+i)
		fmt.Println("Sending to Client: ", msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			println("Error sending Msg", err)
			os.Exit(1)
		}
		var reply string
		err = websocket.Message.Receive(ws, reply)
		if err != nil {
			println("Error receiving Msg", err)
			os.Exit(1)
		}
		fmt.Println("Received back from client: ", reply)

	}
}

func main() {
	fmt.Println("start")
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(2)
	}
}
