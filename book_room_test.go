package cqrs_booking_test

import (
	"github.com/stretchr/testify/assert"
	"kata/cqrs_booking/room_command"
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_read_registry"
	"kata/cqrs_booking/room_write_registry"
	"kata/cqrs_booking/utils"
)

type bookRoomFeature struct {
	roomQueryService   *room_query.RoomQueryService
	roomCommandService *room_command.RoomCommandService
}

func (b *bookRoomFeature) GivenRoom(room string) error {
	inMemoryRoomReadRegistry := room_read_registry.NewInMemoryRoomReadRegistry()
	inMemoryRoomReadRegistry.AddRoom(room)
	b.roomQueryService = room_query.NewRoomQueryService(inMemoryRoomReadRegistry)

	w := room_write_registry.NewInMemoryRoomWriteRegistry()
	b.roomCommandService = room_command.NewRoomCommandService(w, inMemoryRoomReadRegistry)
	return nil
}

func (b *bookRoomFeature) RoomIsFree(room, arrivalDateString, departureDateString string) error {
	arrivalDate := utils.DateFor(arrivalDateString)
	departureDate := utils.DateFor(departureDateString)
	freeRooms := b.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return assertExpectedAndActual(assert.Contains, freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}

func (b *bookRoomFeature) BookRoom(client, room, arrivalDateString, departureDateString string) error {
	arrivalDate := utils.DateFor(arrivalDateString)
	departureDate := utils.DateFor(departureDateString)
	booking := room_write_registry.NewBooking(client, room, arrivalDate, departureDate)
	b.roomCommandService.BookRoom(booking)
	return nil
}

func (b *bookRoomFeature) RoomIsNotFree(room, arrivalDateString, departureDateString string) error {
	arrivalDate := utils.DateFor(arrivalDateString)
	departureDate := utils.DateFor(departureDateString)
	freeRooms := b.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return assertExpectedAndActual(assert.NotContains, freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}
