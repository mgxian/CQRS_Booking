package room_command

import (
	"kata/cqrs_booking/room_query"
	"time"
)

type RoomCommandService struct {
	rw RoomWriteRegistry
	rr room_query.RoomReadRegistry
}

func NewRoomCommandService(rw RoomWriteRegistry, rr room_query.RoomReadRegistry) *RoomCommandService {
	return &RoomCommandService{rw, rr}
}

func (s *RoomCommandService) BookRoom(booking Booking) {
	s.rw.BookRoom(booking)
	s.rr.BookRoom(booking.Name(), booking.Arrival(), booking.Departure())
}

type RoomWriteRegistry interface {
	BookRoom(booking Booking)
}

type Booking struct {
	client, room       string
	arrival, departure time.Time
}

func (b Booking) Arrival() time.Time {
	return b.arrival
}

func (b Booking) Departure() time.Time {
	return b.departure
}

func (b Booking) Name() string {
	return b.room
}

func (b Booking) ClientID() string {
	return b.client
}

func NewBooking(client, room string, arrival, departure time.Time) Booking {
	return Booking{client, room, arrival, departure}
}
