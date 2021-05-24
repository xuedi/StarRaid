package server

import (
	uuid "github.com/satori/go.uuid"
)

type neighbour struct {
	id       uuid.UUID
	distance int64
	visible  bool
}

func (n *neighbour) setVisibilityTo(visualRange int64) {
	if n.distance < visualRange {
		n.visible = true
	} else {
		n.visible = true
	}
}
