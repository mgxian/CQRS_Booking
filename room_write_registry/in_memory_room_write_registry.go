package room_write_registry

import "kata/cqrs_booking/room_command"

type InMemoryRoomWriteRegistry struct {
	bookings map[string]room_command.Booking
}

func NewInMemoryRoomWriteRegistry() *InMemoryRoomWriteRegistry {
	return &InMemoryRoomWriteRegistry{
		bookings: make(map[string]room_command.Booking,0),
	}
}

func (i *InMemoryRoomWriteRegistry) BookRoom(booking room_command.Booking) {
	i.bookings[booking.ClientID()] = booking
}

func (i *InMemoryRoomWriteRegistry) GetBooking(clientID string) (room_command.Booking) {
	return i.bookings[clientID]
}
