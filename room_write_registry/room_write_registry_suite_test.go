package room_write_registry_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoomWriteRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RoomWriteRegistry Suite")
}
