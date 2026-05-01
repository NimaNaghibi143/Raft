package main

import (
	"fmt"
	"sync"
)

type statemachine struct {
	db     *sync.Map
	server int
}

type commnadKind unit8

const (
	setCommand commnadKind = iota
	getCommand
)

type command struct {
	kind  commnadKind
	key   string
	value string
}

func (s *statemachine) Apply(cmd []byte) ([]byte, error) {
	c := decodeCommand(cmd)

	switch c.kind {
	case setCommand:
		s.db.Store(c.key, c.value)
	case getCommand:
		value, ok := s.db.Load(c.key)
		if !ok {
			return nil, fmt.Errorf("key not found")
		}
		return []byte(value.(string)), nil
	default:
		return nil, fmt.Errorf("Unknown command: %x", cmd)
	}

	return nil, nil
}
