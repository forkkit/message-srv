package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/errors"
	"github.com/microhq/message-srv/message"
	proto "github.com/microhq/message-srv/proto/message"

	"golang.org/x/net/context"
)

type Message struct{}

func (m *Message) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	if req.Event == nil {
		return errors.BadRequest("go.micro.srv.message", "invalid event")
	}

	if len(req.Event.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Event.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	if len(req.Event.Id) == 0 {
		req.Event.Id = uuid.New().String()
	}

	req.Event.Created = time.Now().UnixNano()
	req.Event.Updated = 0

	if err := message.Create(req.Event); err != nil && err == message.ErrAlreadyExists {
		return errors.BadRequest("go.micro.srv.message", err.Error())
	} else if err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	return nil
}

func (m *Message) Update(ctx context.Context, req *proto.UpdateRequest, rsp *proto.UpdateResponse) error {
	if req.Event == nil {
		return errors.BadRequest("go.micro.srv.message", "invalid event")
	}

	if len(req.Event.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Event.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	if len(req.Event.Id) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid id")
	}

	if req.Event.Created == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid timestamp")
	}

	req.Event.Updated = time.Now().UnixNano()

	if err := message.Update(req.Event); err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	return nil
}

func (m *Message) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	if len(req.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	if len(req.Id) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid id")
	}

	if err := message.Delete(req.Id, req.Namespace, req.Channel); err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	return nil
}

func (m *Message) Search(ctx context.Context, req *proto.SearchRequest, rsp *proto.SearchResponse) error {
	if len(req.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	if req.Offset <= 0 {
		req.Offset = 0
	}

	events, err := message.Search(req.Query, req.Namespace, req.Channel, int(req.Limit), int(req.Offset), req.Reverse)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	rsp.Events = events

	return nil
}

func (m *Message) Stream(ctx context.Context, req *proto.StreamRequest, stream proto.Message_StreamStream) error {
	if len(req.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	ch, exit, err := message.Stream(req.Namespace, req.Channel)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	defer func() {
		close(exit)
		stream.Close()
	}()

	for {
		select {
		case e := <-ch:
			if err := stream.Send(&proto.StreamResponse{Event: e}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Message) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	if len(req.Namespace) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid namespace")
	}

	if len(req.Channel) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid channel")
	}

	if len(req.Id) == 0 {
		return errors.BadRequest("go.micro.srv.message", "invalid id")
	}

	event, err := message.Read(req.Id, req.Namespace, req.Channel)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.message", err.Error())
	}

	rsp.Event = event

	return nil
}
