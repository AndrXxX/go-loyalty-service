package queue

type queueJob interface {
	Execute() error
}
