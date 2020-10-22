package room_write_registry

type InMemoryRoomWriteRegistry struct {
	bookings map[string]Booking
}

func NewInMemoryRoomWriteRegistry() *InMemoryRoomWriteRegistry {
	return &InMemoryRoomWriteRegistry{
		bookings: make(map[string]Booking, 0),
	}
}

func (i *InMemoryRoomWriteRegistry) BookRoom(booking Booking) {
	i.bookings[booking.ClientID()] = booking
}

func (i *InMemoryRoomWriteRegistry) GetBooking(clientID string) Booking {
	return i.bookings[clientID]
}
