package main
import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func serveFileHandler() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/" + "client" + ".html")
	})
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	serveFileHandler()



	if err2 := http.ListenAndServe("", nil); err2 != nil {
		log.Fatal("ListenAndServe:", err2)
	}

}

