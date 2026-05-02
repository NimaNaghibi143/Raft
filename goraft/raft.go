package goraft

import "net/rpc"

type StateMachine interface {
	Apply(cmd []byte) ([]byte, error)
}

type ApplyResult struct {
	Result []byte
	Error  error
}

type Entry struct {
	Command []byte
	Term    uint64

	// Set by the primary so it can learn about the result of
	// applying this command to the state machine
	result chan ApplyResult
}

type ClusterMember struct {
	Id      uint64
	Address string

	// Index of the next log entry to send
	nextIndex uint64
	// Hghest log entry known to be replicated
	matchIndex uint64

	// Whos was voted for in the most recent term
	votedFor uint64

	// TCP connection
	rcpClient *rpc.Client
}
