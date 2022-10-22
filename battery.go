package main

import (
	"math"
)

type Battery struct {
	ID, columnID, elevatorID, floorRequestButtonID, callButtonID        int
	status                                                              string
	columnsList                                                         []Column
	floorRequestsButtonsList                                            []FloorRequestButton
	servedFloors []int
}

// Creates the columns and buttons 
func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	b := new(Battery)
	b.ID = _id
	b.status = "online"

	b.columnID = 1
	b.elevatorID = 1
	b.floorRequestButtonID = 1
	b.callButtonID = 1

	if _amountOfBasements > 0 {
		b.createBasementFloorRequestButtons(_amountOfBasements)
		b.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--

	}

	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn)

	return b

}

// Create the column for the basement and add to the servedFLoors list
func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	b.servedFloors = make([]int, 0)
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		b.servedFloors = append(b.servedFloors, floor)
		floor--

	}

	column := NewColumn(b.columnID, "online", _amountOfBasements, _amountOfElevatorPerColumn, b.servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	b.columnID++
}

// Create columns and add to the servedFloors list
func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := int(math.Ceil((float64(_amountOfFloors)/(float64(_amountOfColumns)))))
	floor := 1

	for i := 0; i < _amountOfColumns; i++ {
		b.servedFloors = make([]int, 0)
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if floor <= _amountOfFloors {
				b.servedFloors = append(b.servedFloors, floor)
				floor++
			}

		}

		column := NewColumn(b.columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, b.servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		b.columnID++

	}
}

// Adds the requests of people on floor levels to the floorRequestsButtonsList list
func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, buttonFloor, "OFF", "Up")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, *floorRequestButton)
		buttonFloor--
		b.floorRequestButtonID++

	}
}

// Adds the requests of people at the basement into the floorRequestsButtonsList list
func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID,buttonFloor, "OFF", "Down")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, *floorRequestButton)
		buttonFloor--
		b.floorRequestButtonID++
	}

}

// Select the best available column for the scenario 
func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			return &column
			
		}

	}
	return nil
}

func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()

	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator

}
