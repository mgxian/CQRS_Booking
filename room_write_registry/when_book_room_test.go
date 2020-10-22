package room_write_registry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kata/cqrs_booking/room_command"
	"kata/cqrs_booking/room_write_registry"
	"kata/cqrs_booking/utils"
)

var _ = Describe("when book room", func() {
	It("should save booking", func() {
		// Given
		inMemoryRoomWriteRegistry := room_write_registry.NewInMemoryRoomWriteRegistry()
		arrivalDate := utils.DateFor("2020-10-21")
		departureDate := utils.DateFor("2020-10-21")
		booking := room_command.NewBooking("will", "shanghai", arrivalDate, departureDate)

		// When
		inMemoryRoomWriteRegistry.BookRoom(booking)

		// Then
		Expect(inMemoryRoomWriteRegistry.GetBooking(booking.ClientID())).Should(Equal(booking))
	})
})
