package cqrs_booking_test

import (
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_read_registry"
	"time"
)

type freeRoomFeature struct {
	freeRooms                []room_read_registry.Room
	inMemoryRoomReadRegistry *room_read_registry.InMemoryRoomReadRegistry
	roomQueryService         *room_query.RoomQueryService
}

func (f *freeRoomFeature) roomWasBookedAt(room string, bookings *godog.Table) error {
	f.inMemoryRoomReadRegistry = room_read_registry.NewInMemoryRoomReadRegistry()
	f.inMemoryRoomReadRegistry.AddRoom(room)

	for _, booking := range bookings.Rows[1:] {
		arrival := booking.Cells[0].Value
		departure := booking.Cells[1].Value
		arrivalDate, _ := time.Parse("2006-1-2", arrival)
		departureDate, _ := time.Parse("2006-1-2", departure)
		f.inMemoryRoomReadRegistry.BookRoom(room, arrivalDate, departureDate)
	}
	f.roomQueryService = room_query.NewRoomQueryService(f.inMemoryRoomReadRegistry)
	return nil
}

func (f *freeRoomFeature) getFreeRoomsAt(arrival, departure string) error {
	arrivalDate, _ := time.Parse("2006-1-2", arrival)
	departureDate, _ := time.Parse("2006-1-2", departure)
	f.freeRooms = f.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return nil
}

func (f *freeRoomFeature) theFreeRoomsContainsRoom(room string) error {
	return assertExpectedAndActual(assert.Contains, f.freeRooms, room_read_registry.NewRoom(room), "Expected free rooms contain room %s", room)
}

func (f *freeRoomFeature) theFreeRoomsNotContainsRoom(room string) error {
	return assertExpectedAndActual(assert.NotContains, f.freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	freeRoomFeature := &freeRoomFeature{}
	ctx.Step(`^the (.*) room was booked as the following:$`, freeRoomFeature.roomWasBookedAt)
	ctx.Step(`^get free rooms arrival at (\d+-\d+-\d+) and departure at (\d+-\d+-\d+)$`, freeRoomFeature.getFreeRoomsAt)
	ctx.Step(`^the free rooms should contains the (.*) room$`, freeRoomFeature.theFreeRoomsContainsRoom)
	ctx.Step(`^the free rooms should not contains the (.*) room$`, freeRoomFeature.theFreeRoomsNotContainsRoom)

}
