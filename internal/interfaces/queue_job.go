package interfaces

type QueueJob interface {
	Execute() error
}
