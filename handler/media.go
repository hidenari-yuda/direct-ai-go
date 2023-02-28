package handler

import (
	"context"

	"github.com/hidenari-yuda/direct-ai-go/infra/database"
	"github.com/hidenari-yuda/direct-ai-go/infra/driver"
	"github.com/hidenari-yuda/direct-ai-go/pb"
	"github.com/hidenari-yuda/direct-ai-go/usecase"
	"github.com/hidenari-yuda/direct-ai-go/usecase/interactor"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement helloworld.GreeterServer.
type MediaServiceServer struct {
	pb.UnimplementedMediaServiceServer
	MediaInteractor interactor.MediaInteractor
	Db              *database.Db
	Firebase        usecase.Firebase
}

func NewMediaSercviceServer(userInteractor interactor.MediaInteractor) *MediaServiceServer {
	return &MediaServiceServer{
		MediaInteractor: userInteractor,
		Db:              database.NewDb(),
		Firebase:        driver.NewFirebaseImpl(),
	}
}

// create user
func (s *MediaServiceServer) Create(ctx context.Context, req *pb.Media) (*pb.Media, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, handleError(err)
	}

	res, err := s.MediaInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return res, nil

}

// update user
func (s *MediaServiceServer) Update(ctx context.Context, req *pb.Media) (*pb.MediaBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, handleError(err)
	}

	res, err := s.MediaInteractor.Update(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return &pb.MediaBoolResponse{Error: res}, nil
}

// delete user
func (s *MediaServiceServer) Delete(ctx context.Context, req *pb.MediaIdRequest) (*pb.MediaBoolResponse, error) {

	tx, err := s.Db.Begin()
	if err != nil {
		return nil, handleError(err)
	}

	res, err := s.MediaInteractor.Delete(req.Id)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return &pb.MediaBoolResponse{Error: res}, nil
}

// get user by id
func (s *MediaServiceServer) GetById(ctx context.Context, req *pb.MediaIdRequest) (*pb.Media, error) {

	res, err := s.MediaInteractor.GetById(req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	return res, nil
}

// get list by user id
func (s *MediaServiceServer) GetListByUser(ctx context.Context, req *pb.MediaIdRequest) (*pb.MediaList, error) {

	res, err := s.MediaInteractor.GetListByUser(req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.MediaList{MediaList: res}, nil
}

// get list by type
func (s *MediaServiceServer) GetListByType(ctx context.Context, req *pb.MediaIdRequest) (*pb.MediaList, error) {

	res, err := s.MediaInteractor.GetListByType(req.Id)
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.MediaList{MediaList: res}, nil
}

// admin
// get all user
func (s *MediaServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*pb.MediaList, error) {

	res, err := s.MediaInteractor.GetAll()
	if err != nil {
		return nil, handleError(err)
	}

	return &pb.MediaList{MediaList: res}, nil
}
