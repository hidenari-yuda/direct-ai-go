package handler

import (
	"context"

	"github.com/hidenari-yuda/direct-ai-go/infra/database"
	"github.com/hidenari-yuda/direct-ai-go/infra/driver"
	"github.com/hidenari-yuda/direct-ai-go/pb"
	"github.com/hidenari-yuda/direct-ai-go/usecase"
	"github.com/hidenari-yuda/direct-ai-go/usecase/interactor"
)

type KeywordServiceServer struct {
	pb.UnimplementedKeywordServiceServer
	KeywordInteractor interactor.KeywordInteractor
	Db                *database.Db
	Firebase          usecase.Firebase
}

func NewKeywordSercviceServer(KeywordInteractor interactor.KeywordInteractor) *KeywordServiceServer {
	return &KeywordServiceServer{
		KeywordInteractor: KeywordInteractor,
		Db:                database.NewDb(),
		Firebase:          driver.NewFirebaseImpl(),
	}
}

// create chat group
func (s *KeywordServiceServer) Create(ctx context.Context, req *pb.Keyword) (*pb.Keyword, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

// update chat group
func (s *KeywordServiceServer) Update(ctx context.Context, req *pb.Keyword) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// delete chat group
func (s *KeywordServiceServer) Delete(ctx context.Context, req *pb.ChatIdRequest) (*pb.ChatBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}

	res, err := s.KeywordInteractor.Delete(req.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &pb.ChatBoolResponse{Error: res}, nil
}

// get chat group by id
func (s *KeywordServiceServer) GetById(ctx context.Context, req *pb.ChatIdRequest) (*pb.Keyword, error) {

	res, err := s.KeywordInteractor.GetById(req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// get chat group by user id
func (s *KeywordServiceServer) GetListByUser(ctx context.Context, req *pb.ChatIdRequest) (*pb.KeywordList, error) {

	res, err := s.KeywordInteractor.GetListByUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.KeywordList{KeywordList: res}, nil
}
