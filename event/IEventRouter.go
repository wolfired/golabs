package event

type IEventRouter interface {
	Route(e IEvent)
}
