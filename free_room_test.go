package cqrs_booking_test

import (
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_read_registry"
	"kata/cqrs_booking/utils"
)

type freeRoomFeature struct {
	freeRooms        []room_read_registry.Room
	roomQueryService *room_query.RoomQueryService
}

func (f *freeRoomFeature) GivenRoomWasBooked(room string, bookings *godog.Table) error {
	inMemoryRoomReadRegistry := room_read_registry.NewInMemoryRoomReadRegistry()
	inMemoryRoomReadRegistry.AddRoom(room)

	for _, booking := range bookings.Rows[1:] {
		arrivalDateString := booking.Cells[0].Value
		departureDateString := booking.Cells[1].Value
		arrivalDate := utils.DateFor(arrivalDateString)
		departureDate := utils.DateFor(departureDateString)
		inMemoryRoomReadRegistry.BookRoom(room, arrivalDate, departureDate)
	}
	f.roomQueryService = room_query.NewRoomQueryService(inMemoryRoomReadRegistry)
	return nil
}

func (f *freeRoomFeature) WhenGetFreeRooms(arrivalDateString, departureDateString string) error {
	arrivalDate := utils.DateFor(arrivalDateString)
	departureDate := utils.DateFor(departureDateString)
	f.freeRooms = f.roomQueryService.FreeRooms(arrivalDate, departureDate)
	return nil
}

func (f *freeRoomFeature) FreeRoomsContainsRoom(room string) error {
	return assertExpectedAndActual(assert.Contains, f.freeRooms, room_read_registry.NewRoom(room), "Expected free rooms contain room %s", room)
}

func (f *freeRoomFeature) FreeRoomsNotContainsRoom(room string) error {
	return assertExpectedAndActual(assert.NotContains, f.freeRooms, room_read_registry.NewRoom(room), "Expected free rooms not contain room %s", room)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	freeRoomFeature := &freeRoomFeature{}
	ctx.Step(`^the (.*) room was booked as the following:$`, freeRoomFeature.GivenRoomWasBooked)
	ctx.Step(`^get free rooms arrival at (\d+-\d+-\d+) and departure at (\d+-\d+-\d+)$`, freeRoomFeature.WhenGetFreeRooms)
	ctx.Step(`^the free rooms should contains the (.*) room$`, freeRoomFeature.FreeRoomsContainsRoom)
	ctx.Step(`^the free rooms should not contains the (.*) room$`, freeRoomFeature.FreeRoomsNotContainsRoom)

	bookRoomFeature := &bookRoomFeature{}
	ctx.Step(`^the (.*) room$`, bookRoomFeature.GivenRoom)
	ctx.Step(`^the (.*) room is free arrival at (\d+-\d+-\d+) and departure at (\d+-\d+-\d+)$`, bookRoomFeature.RoomIsFree)
	ctx.Step(`^(.*) book the (.*) room arrival at (\d+-\d+-\d+) and departure at (\d+-\d+-\d+)$`, bookRoomFeature.BookRoom)
	ctx.Step(`^the (.*) room is not free arrival at (\d+-\d+-\d+) and departure at (\d+-\d+-\d+)$`, bookRoomFeature.RoomIsNotFree)
}
