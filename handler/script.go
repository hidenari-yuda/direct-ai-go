package handler

import (
	"context"

	"github.com/hidenari-yuda/direct-ai-go/infra/database"
	"github.com/hidenari-yuda/direct-ai-go/pb"
	"github.com/hidenari-yuda/direct-ai-go/usecase"
	"github.com/hidenari-yuda/direct-ai-go/usecase/interactor"
)

type ScriptServiceServer struct {
	pb.UnimplementedScriptServiceServer
	ScriptInteractor interactor.ScriptInteractor
	Db               *database.Db
	Firebase         usecase.Firebase
}

func (s *ScriptServiceServer) Create(ctx context.Context, req *pb.Script) (*pb.Script, error) {

	tx, _ := s.Db.Begin()
	res, err := s.ScriptInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return res, nil
}

func (s *ScriptServiceServer) CreateScriptByKeyword(ctx context.Context, req *pb.Script) (*pb.Script, error) {

	tx, _ := s.Db.Begin()
	res, err := s.ScriptInteractor.Create(req)
	if err != nil {
		tx.Rollback()
		return nil, handleError(err)
	}

	tx.Commit()

	return res, nil
}
