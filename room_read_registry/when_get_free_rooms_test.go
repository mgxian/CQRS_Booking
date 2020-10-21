package room_read_registry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kata/cqrs_booking/room_read_registry"
	"time"
)

var _ = Describe("when get free rooms", func() {
	It("room is free when that day is not booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrival, _ := time.Parse("2006-1-2", "2020-10-20")
		departure, _ := time.Parse("2006-1-2", "2020-10-21")
		oneDay := time.Hour * 24
		r.BookRoom("shanghai", arrival, departure)
		freeRooms := r.FreeRooms(arrival.Add(oneDay), departure.Add(oneDay))

		// Then
		Expect(freeRooms).Should(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free when that day is booked", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrival, _ := time.Parse("2006-1-2", "2020-10-20")
		departure, _ := time.Parse("2006-1-2", "2020-10-21")
		r.BookRoom("shanghai", arrival, departure)
		freeRooms := r.FreeRooms(arrival, departure)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free on everyday from arrival date to departure date", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrival, _ := time.Parse("2006-1-2", "2020-10-20")
		departure, _ := time.Parse("2006-1-2", "2020-10-23")
		r.BookRoom("shanghai", arrival, departure)

		oneDay := time.Hour * 24
		firstDaySinceArrival := arrival
		secondDaySinceArrival := firstDaySinceArrival.Add(oneDay)
		thirdDaySinceArrival := secondDaySinceArrival.Add(oneDay)
		forthDaySinceArrival := thirdDaySinceArrival.Add(oneDay)

		// Then
		Expect(r.IsRoomFreeOn("shanghai", firstDaySinceArrival)).Should(Equal(false))
		Expect(r.IsRoomFreeOn("shanghai", secondDaySinceArrival)).Should(Equal(false))
		Expect(r.IsRoomFreeOn("shanghai", thirdDaySinceArrival)).Should(Equal(false))
		Expect(r.IsRoomFreeOn("shanghai", forthDaySinceArrival)).Should(Equal(true))
	})

	It("room is not free when one day booking date between arrival date and departure date", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrival, _ := time.Parse("2006-1-2", "2020-10-20")
		departure, _ := time.Parse("2006-1-2", "2020-10-23")
		r.BookRoom("shanghai", arrival, departure)

		wantArrival, _ := time.Parse("2006-1-2", "2020-10-21")
		wantDeparture, _ := time.Parse("2006-1-2", "2020-10-22")
		freeRooms := r.FreeRooms(wantArrival, wantDeparture)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})

	It("room is not free when more than one day booking date between arrival date and departure date", func() {
		// Given
		r := room_read_registry.NewInMemoryRoomReadRegistry()
		r.AddRoom("shanghai")

		// When
		arrival, _ := time.Parse("2006-1-2", "2020-10-20")
		departure, _ := time.Parse("2006-1-2", "2020-10-26")
		r.BookRoom("shanghai", arrival, departure)

		wantArrival, _ := time.Parse("2006-1-2", "2020-10-19")
		wantDeparture, _ := time.Parse("2006-1-2", "2020-10-22")
		freeRooms := r.FreeRooms(wantArrival, wantDeparture)

		// Then
		Expect(freeRooms).ShouldNot(ConsistOf(room_read_registry.NewRoom("shanghai")))
	})
})
