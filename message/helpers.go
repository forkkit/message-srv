package message

import (
	proto "github.com/microhq/message-srv/proto/message"
)

type sortedEvents struct {
	events []*proto.Event
}

func (s sortedEvents) Len() int {
	return len(s.events)
}

func (s sortedEvents) Less(i, j int) bool {
	return s.events[i].Created < s.events[j].Created
}

func (s sortedEvents) Swap(i, j int) {
	s.events[i], s.events[j] = s.events[j], s.events[i]
}
