/*
      __    __
     /\ \  /\ \
     \ `\`\\/'/   __        __       ___
      `\ `\ /'  /'__`\    /'_ `\    / __`\
        `\ \ \ /\ \L\.\_ /\ \L\ \  /\ \L\ \
          \ \_\\ \__/.\_\\ \____ \ \ \____/
           \/_/ \/__/\/_/ \/___L\ \ \/___/
      /'\_/`\               /\____/
     /\      \     ___    _ \_/__/____     __              ____     __    _ __   __  __     __    _ __
     \ \ \__\ \   / __`\ /\`'__\ /',__\  /'__`\  _______  /',__\  /'__`\ /\`'__\/\ \/\ \  /'__`\ /\`'__\
      \ \ \_/\ \ /\ \L\ \\ \ \/ /\__, `\/\  __/ /\______\/\__, `\/\  __/ \ \ \/ \ \ \_/ |/\  __/ \ \ \/
       \ \_\\ \_\\ \____/ \ \_\ \/\____/\ \____\\/______/\/\____/\ \____\ \ \_\  \ \___/ \ \____\ \ \_\
        \/_/ \/_/ \/___/   \/_/  \/___/  \/____/          \/___/  \/____/  \/_/   \/__/   \/____/  \/_/

*/
package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":80", "http service address")
var homeTemplate = template.Must(template.ParseFiles("../static/client.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}