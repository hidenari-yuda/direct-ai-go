package handler

import (
	"context"
	"log"

	"github.com/hidenari-yuda/direct-ai-go/infra/database"
	"github.com/hidenari-yuda/direct-ai-go/infra/driver"
	"github.com/hidenari-yuda/direct-ai-go/pb"
	"github.com/hidenari-yuda/direct-ai-go/usecase"
	"github.com/hidenari-yuda/direct-ai-go/usecase/interactor"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	ChatInteractor interactor.ChatInteractor
	Db             *database.Db
	Firebase       usecase.Firebase
}

func NewChatSercviceServer(chatInteractor interactor.ChatInteractor) *ChatServiceServer {
	return &ChatServiceServer{
		ChatInteractor: chatInteractor,
		Db:             database.NewDb(),
		Firebase:       driver.NewFirebaseImpl(),
	}
}

// create chat
func (s *ChatServiceServer) Create(ctx context.Context, req *pb.Chat) (*pb.Chat, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update chat
func (s *ChatServiceServer) Update(ctx context.Context, req *pb.Chat) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// delete chat
func (s *ChatServiceServer) Delete(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.ChatInteractor.Delete(uint(req.Id))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// get chat  by id
func (s *ChatServiceServer) GetById(ctx context.Context, req *pb.ChatIdRequest) (*pb.Chat, error) {

	res, err := s.ChatInteractor.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get chat  by user id
func (s *ChatServiceServer) GetListByGroup(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatList, error) {

	res, err := s.ChatInteractor.GetListByGroup(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ChatList{ChatList: res}, nil
}

func (s *ChatServiceServer) GetStream(req *pb.GetChatStreamRequest, server pb.ChatService_GetStreamServer) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := make(chan pb.Chat)

	go func() {
		err := s.ChatInteractor.GetStream(ctx, stream)
		if err != nil {
			log.Println(err)
		}
	}()

	for {
		v := <-stream

		// createdAt := timestamppb.New(v.CreatedAt)
		if err := server.Send(&pb.Chat{
			From:      v.From,
			To:        v.To,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		},
		); err != nil {
			return err
		}
	}
}
