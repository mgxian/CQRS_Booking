package cqrs_booking_test

import (
	"github.com/stretchr/testify/assert"
	"kata/cqrs_booking/room_command"
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_read_registry"
	"kata/cqrs_booking/room_write_registry"
	"time"
)

type bookRoomFeature struct {
	roomQueryService         *room_query.RoomQueryService
	roomCommandService       *room_command.RoomCommandService
}

func (b *bookRoomFeature) givenRoom(room string) error {
	inMemoryRoomReadRegistry := room_read_registry.NewInMemoryRoomReadRegistry()
	inMemoryRoomReadRegistry.AddRoom(room)
	b.roomQueryService = room_query.NewRoomQueryService(inMemoryRoomReadRegistry)

	w := room_write_registry.NewInMemoryRoomWriteRegistry()
	b.roomCommandService = room_command.NewRoomCommandService(w, inMemoryRoomReadRegistry)
	return nil
}

func (b *bookRoomFeature) theRoomIsFree(room, arrivalDateString, departureDateString string) error {
	arrivalDate, _ := time.Parse("2006-1-2", arrivalDateString)
	departureDate, _ := time.Parse("2006-1-2", departureDateString)
	freeRooms := b.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return assertExpectedAndActual(assert.Contains, freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}

func (b *bookRoomFeature) bookRoom(client, room, arrivalDateString, departureDateString string) error {
	arrivalDate, _ := time.Parse("2006-1-2", arrivalDateString)
	departureDate, _ := time.Parse("2006-1-2", departureDateString)
	booking := room_command.NewBooking(client, room, arrivalDate, departureDate)
	b.roomCommandService.BookRoom(booking)
	return nil
}

func (b *bookRoomFeature) theRoomIsNotFree(room, arrivalDateString, departureDateString string) error {
	arrivalDate, _ := time.Parse("2006-1-2", arrivalDateString)
	departureDate, _ := time.Parse("2006-1-2", departureDateString)
	freeRooms := b.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return assertExpectedAndActual(assert.NotContains, freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}
