package room_write_registry

import "time"

type Booking struct {
	client, room       string
	arrival, departure time.Time
}

func (b Booking) Arrival() time.Time {
	return b.arrival
}

func (b Booking) Departure() time.Time {
	return b.departure
}

func (b Booking) Name() string {
	return b.room
}

func (b Booking) ClientID() string {
	return b.client
}

func NewBooking(client, room string, arrival, departure time.Time) Booking {
	return Booking{client, room, arrival, departure}
}
