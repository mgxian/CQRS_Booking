package room_read_registry

import (
	"time"
)

type InMemoryRoomReadRegistry struct {
	freeRooms map[string]map[time.Time]bool
}

func NewInMemoryRoomReadRegistry() *InMemoryRoomReadRegistry {
	return &InMemoryRoomReadRegistry{
		freeRooms: make(map[string]map[time.Time]bool, 0),
	}
}

func (i *InMemoryRoomReadRegistry) FreeRooms(arrival time.Time, departure time.Time) []Room {
	result := make([]Room, 0)
	for room, _ := range i.freeRooms {
		if i.IsRoomFree(room, arrival, departure) {
			result = append(result, NewRoom(room))
		}
	}
	return result
}

func (i *InMemoryRoomReadRegistry) IsRoomFree(room string, arrival, departure time.Time) bool {
	bookedDate := arrival
	oneDay := time.Hour * 24
	for bookedDate.Before(departure) {
		if !i.isRoomFreeOn(room, bookedDate) {
			return false
		}
		bookedDate = bookedDate.Add(oneDay)
	}
	return true
}

func (i *InMemoryRoomReadRegistry) AddRoom(name string) {
	if _, ok := i.freeRooms[name]; ok {
		return
	}
	i.freeRooms[name] = make(map[time.Time]bool, 0)
}

func (i *InMemoryRoomReadRegistry) BookRoom(name string, arrival, departure time.Time) {
	if _, ok := i.freeRooms[name]; !ok {
		return
	}

	bookedDate := arrival
	oneDay := time.Hour * 24
	for bookedDate.Before(departure) {
		i.freeRooms[name][bookedDate] = true
		bookedDate = bookedDate.Add(oneDay)
	}
}

func (i *InMemoryRoomReadRegistry) isRoomFreeOn(room string, date time.Time) bool {
	if booked, ok := i.freeRooms[room][date]; ok {
		return !booked
	}
	return true
}
