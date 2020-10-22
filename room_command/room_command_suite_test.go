package room_command_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoomCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RoomCommand Suite")
}
