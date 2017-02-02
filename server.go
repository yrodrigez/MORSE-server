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

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL.Path[1:]= ", r.URL.Path[1:])
		http.ServeFile(w, r, "static/" + r.URL.Path[1:] + ".html")
	})

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

