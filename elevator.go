package main

import (
	"sort"
)

type Elevator struct {
	ID, amountOfFloors, currentFloor       int
	status, direction                      string
	floorRequestsList, completedRequestsList []int
	door                                   Door
}


func NewElevator(_id int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	e := new(Elevator)
	e.ID = _id
	e.status = _status
	e.floorRequestsList = make([]int, 0)
	e.amountOfFloors = _amountOfFloors
	e.currentFloor = _currentFloor
	e.direction = ""
	e.door = Door{_id, "closed"}

	return e
}

// Move elevator to the direction based on what floor we are currently
func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		e.sortFloorList()
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		if !contains(e.completedRequestsList, destination){
			e.completedRequestsList = append(e.completedRequestsList, destination)
		}
		e.floorRequestsList = e.floorRequestsList[1:]

	}
	e.status = "idle"
	e.direction = "empty"

}

// Sort list function to pick up other requests mid destination
func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Sort(sort.IntSlice(e.floorRequestsList))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}

}

// Door operations
func (e *Elevator) operateDoors() {
	if e.status == "stopped" || e.status == "idle" {
		e.door.status = "open"

		if len(e.floorRequestsList) < 1 {
			e.direction = ""
			e.status = "idle"
		}
	}
}

// Add new requests to the floorRequest list to pick up
func (e *Elevator) addNewRequest(requestedFloor int) {
	if !contains(e.floorRequestsList, requestedFloor) {
		e.floorRequestsList = append([]int{requestedFloor}, e.floorRequestsList...)

	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
