package main

type Door struct {
	ID int
	status string

}

// Door class
func NewDoor(_id int, _status string) *Door {
	d := new(Door)
	d.ID = _id
	d.status = _status

	return d

}
