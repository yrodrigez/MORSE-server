package main
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
// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}