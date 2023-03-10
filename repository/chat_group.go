package repository

import (
	"time"

	"github.com/hidenari-yuda/direct-ai-go/domain/utils"
	"github.com/hidenari-yuda/direct-ai-go/pb"
	"github.com/hidenari-yuda/direct-ai-go/usecase"
)

type ChatGroupRepositoryImpl struct {
	Name     string
	executer SQLExecuter
}

func NewChatGroupRepositoryImpl(ex SQLExecuter) usecase.ChatGroupRepository {
	return &ChatGroupRepositoryImpl{
		Name:     "ChatGroupRepository",
		executer: ex,
	}
}

/***** Create *****/
func (r *ChatGroupRepositoryImpl) Create(param *pb.ChatGroup) error {
	now := time.Now()

	_, err := r.executer.Exec(
		r.Name+"Create",
		`INSERT INTO chat_groups (
			uuid,
			user_id,
			name,
			desc,
			created_at,
			updated_at,
			is_deleted,
			) VALUES (
				?,
				?, 
				?,
				?,
				?,
				?,
				?
		)`,
		utils.CreateUuid(),
		param.UserId,
		param.Name,
		param.Desc,
		now,
		now,
		false,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Update *****/
func (r *ChatGroupRepositoryImpl) Update(param *pb.ChatGroup) error {
	_, err := r.executer.Exec(
		r.Name+"Update",
		`UPDATE chat_groups SET
			name = ?,
			desc = ?,
			updated_at = ?`,
		param.Name,
		param.Desc,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Delete *****/
func (r *ChatGroupRepositoryImpl) Delete(id uint) error {
	_, err := r.executer.Exec(
		r.Name+"Delete",
		"UPDATE chat_groups SET is_deleted = ? WHERE id = ?",
		true,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

/***** Get *****/
func (r *ChatGroupRepositoryImpl) GetById(id uint) (ChatGroup *pb.ChatGroup, err error) {
	err = r.executer.Get(
		r.Name+"SignIn",
		ChatGroup,
		"SELECT * FROM chat_groups WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	return ChatGroup, nil
}

// get list by user id
func (r *ChatGroupRepositoryImpl) GetListByUser(userId uint) (ChatGroups []*pb.ChatGroup, err error) {
	err = r.executer.Select(
		r.Name+"GetListByUser",
		&ChatGroups,
		`
		SELECT * 
		FROM chat_groups 
		WHERE user_id = ? AND is_deleted = ?
		ORDER BY id DESC
		`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	return ChatGroups, nil
}
