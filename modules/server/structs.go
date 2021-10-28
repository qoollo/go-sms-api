package server

import (
	"context"

	"github.com/minish144/go-sms-api/gen/pb"
	"github.com/minish144/go-sms-api/modules/controllers/messages"
)

type ApiServiceServer struct {
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) SendMessage(ctx context.Context, in *pb.Messages_SendRequest) (*pb.Messages_SendResponse, error) {
	return messages.Send(ctx, in)
}

func (s *ApiServiceServer) ListMessages(ctx context.Context, in *pb.Messages_ListRequest) (*pb.Messages_ListResponse, error) {
	return messages.List(ctx, in)
}
