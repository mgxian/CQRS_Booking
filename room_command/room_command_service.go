package room_command

import (
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_write_registry"
)

type RoomCommandService struct {
	rw RoomWriteRegistry
	rr room_query.RoomReadRegistry
}

func NewRoomCommandService(rw RoomWriteRegistry, rr room_query.RoomReadRegistry) *RoomCommandService {
	return &RoomCommandService{rw, rr}
}

func (s *RoomCommandService) BookRoom(booking room_write_registry.Booking) {
	s.rw.BookRoom(booking)
	s.rr.BookRoom(booking.Name(), booking.Arrival(), booking.Departure())
}

type RoomWriteRegistry interface {
	BookRoom(booking room_write_registry.Booking)
}

