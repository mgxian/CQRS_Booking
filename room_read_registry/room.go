package room_read_registry

type Room struct {
	name string
}

func NewRoom(name string) Room {
	return Room{name: name}
}
