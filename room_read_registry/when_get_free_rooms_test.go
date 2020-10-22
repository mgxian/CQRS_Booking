package room_read_registry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kata/cqrs_booking/room_read_registry"
	"kata/cqrs_booking/utils"
	"time"
)

var _ = Describe("when get free rooms", func() {
	It("room is free when that day is not booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrivalDate := utils.DateFor("2020-10-20")
		departureDate := utils.DateFor("2020-10-21")
		oneDay := time.Hour * 24
		r.BookRoom("shanghai", arrivalDate, departureDate)
		freeRooms := r.FreeRooms(arrivalDate.Add(oneDay), departureDate.Add(oneDay))

		// Then
		Expect(freeRooms).Should(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free when that day is booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrivalDate := utils.DateFor("2020-10-20")
		departureDate := utils.DateFor("2020-10-21")
		r.BookRoom("shanghai", arrivalDate, departureDate)
		freeRooms := r.FreeRooms(arrivalDate, departureDate)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free for one day booking when room already booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrivalDate := utils.DateFor("2020-10-20")
		departureDate := utils.DateFor("2020-10-23")
		r.BookRoom("shanghai", arrivalDate, departureDate)

		wantArrival := utils.DateFor("2020-10-21")
		wantDeparture := utils.DateFor("2020-10-22")
		freeRooms := r.FreeRooms(wantArrival, wantDeparture)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free for more than one day booking when room already booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrivalDate := utils.DateFor("2020-10-20")
		departureDate := utils.DateFor("2020-10-26")
		r.BookRoom("shanghai", arrivalDate, departureDate)

		wantArrival := utils.DateFor("2020-10-19")
		wantDeparture := utils.DateFor("2020-10-22")
		freeRooms := r.FreeRooms(wantArrival, wantDeparture)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})
})
