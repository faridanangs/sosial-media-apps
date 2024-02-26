package ws

type Room struct {
	ID    string           `json:"id"`
	Name  string           `json:"name"`
	Users map[string]*User `json:"users"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *User
	UnRegister chan *User
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *User),
		UnRegister: make(chan *User),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]
				if _, ok := r.Users[cl.ID]; !ok {
					r.Users[cl.ID] = cl
				}
			}
		case cl := <-h.UnRegister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Users[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Users) != 0 {
						h.Broadcast <- &Message{
							Content:  cl.Username + " meninggalkan ruang obrolan",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Users, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Users {
					cl.Message <- m
				}
			}
		}
	}
}
