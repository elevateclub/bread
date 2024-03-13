package event

import (
	"time"
)

type Interface interface {
	Event() (*Object, error)
}

type Object struct {
	Name   string
	Time   time.Time
	Author string
}

func Now() *Object {
	return &Object{Time: time.Now()}
}

func New(called string, at time.Time, by string) *Object {
	return &Object{Name: called, Time: at, Author: By}
}
