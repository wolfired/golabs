package event

type IEventRouter interface {
	Router(e IEvent)
}
