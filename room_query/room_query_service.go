package room_query

import (
	"kata/cqrs_booking/room_read_registry"
	"time"
)

type RoomReadRegistry interface {
	FreeRooms(arrival, departure time.Time) []room_read_registry.Room
	BookRoom(name string, arrival, departure time.Time)
}

type RoomQueryService struct {
	rr RoomReadRegistry
}

func (s *RoomQueryService) FreeRooms(arrival time.Time, departure time.Time) []room_read_registry.Room {
	return s.rr.FreeRooms(arrival, departure)
}

func NewRoomQueryService(rr RoomReadRegistry) *RoomQueryService {
	return &RoomQueryService{
		rr: rr,
	}
}
