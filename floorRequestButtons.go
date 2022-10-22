package main


type FloorRequestButton struct {
	ID, floor int
	status, direction string


}
//FloorRequestButton is a button on the pannel at the lobby to request any floor
func NewFloorRequestButton(_id int, _floor int, _status string, _direction string) *FloorRequestButton {
	frq := new(FloorRequestButton)
	frq.ID = _id
	frq.status = _status
	frq.floor = _floor
	frq.direction = _direction

	return frq

}
