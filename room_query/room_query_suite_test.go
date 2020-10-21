package room_query_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCqrsBooking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CqrsBooking Suite")
}
