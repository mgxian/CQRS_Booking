package room_command_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"kata/cqrs_booking/room_command"
	"time"
)

var _ = Describe("when book room", func() {
	var (
		ctrl                  *gomock.Controller
		mockRoomWriteRegistry *MockRoomWriteRegistry
		mockRoomReadRegistry  *MockRoomReadRegistry
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRoomWriteRegistry = NewMockRoomWriteRegistry(ctrl)
		mockRoomReadRegistry = NewMockRoomReadRegistry(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("notify room read registry and room write registry", func() {
		arrivalDate, _ := time.Parse("2006-1-2", "2020-10-21")
		departureDate, _ := time.Parse("2006-1-2", "2020-10-21")
		booking := room_command.NewBooking("will", "shanghai", arrivalDate, departureDate)
		roomCommandService := room_command.NewRoomCommandService(mockRoomWriteRegistry, mockRoomReadRegistry)

		mockRoomWriteRegistry.EXPECT().BookRoom(gomock.Eq(booking))
		mockRoomReadRegistry.EXPECT().BookRoom(booking.Name(), arrivalDate, departureDate)
		roomCommandService.BookRoom(booking)
	})
})
