package room_query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kata/cqrs_booking/room_query"
	"kata/cqrs_booking/room_read_registry"
	"kata/cqrs_booking/utils"
	"time"
)

type StubRoomReadRegistry struct {
}

func (srr *StubRoomReadRegistry) BookRoom(name string, arrival, departure time.Time) {
}

func (srr *StubRoomReadRegistry) FreeRooms(arrival, departure time.Time) []room_read_registry.Room {
	return []room_read_registry.Room{
		room_read_registry.NewRoom("shanghai"),
	}
}

func NewStubRoomReadRegistry() *StubRoomReadRegistry {
	return &StubRoomReadRegistry{}
}

var _ = Describe("when query free rooms", func() {
	It("should contain the free room", func() {
		// Given
		rq := room_query.NewRoomQueryService(NewStubRoomReadRegistry())

		// When
		arrivalDate := utils.DateFor("2020-10-20")
		departureDate := utils.DateFor("2020-10-21")
		freeRooms := rq.FreeRooms(arrivalDate, departureDate)

		// Then
		Expect(freeRooms).Should(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})
})
