package handlers

import "github.com/ElrondNetwork/notifier-go/data"

// LockService defines the behaviour of a lock service component.
// It makes sure that a duplicated entry is not processed multiple times,
// it lockes an item once it has been processed.
type LockService interface {
	IsBlockProcessed(blockHash string) (bool, error)
	HasConnection() bool
	IsInterfaceNil() bool
}

// Publisher defines the behaviour of a publisher component
type Publisher interface {
	Run()
	BroadcastChan() chan<- data.BlockEvents
	BroadcastRevertChan() chan<- data.RevertBlock
	BroadcastFinalizedChan() chan<- data.FinalizedBlock
}
