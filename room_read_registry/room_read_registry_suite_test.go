package room_read_registry_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoomReadRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RoomReadRegistry Suite")
}
