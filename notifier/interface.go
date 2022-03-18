package notifier

import (
	"github.com/ElrondNetwork/notifier-go/data"
)

// PublisherService defines the behaviour of a publisher component which should be
// able to publish received events and broadcast them to channels
type PublisherService interface {
	BroadcastChan() chan<- data.BlockEvents
	BroadcastRevertChan() chan<- data.RevertBlock
	BroadcastFinalizedChan() chan<- data.FinalizedBlock
	IsInterfaceNil() bool
}

// TODO: evaluate adding Subscriber interface and split Hub interface if suitable
