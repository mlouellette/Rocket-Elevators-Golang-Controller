package main

type CallButton struct {
	ID, floor int
	status, direction string

}
//Button on a floor or basement to go back to lobby
func NewCallButton(_id int, _status string, _floor int, _direction string) *CallButton {
	cb := new(CallButton)
	cb.ID = _id
	cb.status = _status
	cb.floor = _floor
	cb.direction = _direction

	return cb

}
