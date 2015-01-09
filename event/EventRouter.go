package event

type EventRouter struct {
	event_map map[string]IEvent
}

func (this *EventRouter) Route(e IEvent) {

}
