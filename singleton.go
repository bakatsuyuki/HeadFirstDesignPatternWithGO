package main

import (
	"github.com/google/uuid"
)

type Singleton struct{ text string }

func (s Singleton) toString() string {
	return s.text
}

var _singletonInstance = &Singleton{uuid.New().String()}

func SingletonInstance() *Singleton {
	return _singletonInstance
}
