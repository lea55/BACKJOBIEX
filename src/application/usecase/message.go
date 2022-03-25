package usecase

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/application/mappers"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"
	"github.com/pkg/errors"
)

type Message struct {
	messageRepo          repository.Message
	userConversationRepo repository.UserConversation
	userRepo             repository.User
	mapper               *mappers.Message
}

func NewMessage(
	messageRepo repository.Message,
	userRepo repository.User,
	userConversationRepo repository.UserConversation,
) *Message {
	return &Message{
		userConversationRepo: userConversationRepo,
		messageRepo:          messageRepo,
		userRepo:             userRepo,
		mapper:               mappers.NewMessage(),
	}
}

func (m Message) GetMessages(userID string, userSenderID string, page uint32) ([]entity.Message, error) {
	from := page * 10
	limit := uint8(from + 10)

	list, err := m.messageRepo.FindByUser(userID, userSenderID, from, limit)
	if err != nil {
		return list, err
	}

	listFromOther, err := m.messageRepo.FindByUser(userSenderID, userID, from, limit)
	if err != nil {
		return list, err
	}

	resultList := make([]entity.Message, 0)

	resultList = append(resultList, list...)
	resultList = append(resultList, listFromOther...)

	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].CreatedAt.Before(resultList[j].CreatedAt)
	})
	return resultList, err
}

func (m Message) CreateNewMessage(dtoModel dto.RqSaveMessage) (string, error) {
	userFrom, err := m.userRepo.FindByID(dtoModel.UserFromID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de usuario")
	}

	userTo, err := m.userRepo.FindByID(dtoModel.UserToID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de usuario")
	}

	message := m.mapper.FromCreateRequest(userFrom, userTo, dtoModel.Content)

	err = m.messageRepo.Save(message)
	if err != nil {
		return "", errors.Wrap(err, "Error al guardar en la base de datos")
	}

	go m.updateUserConversation(message)

	return message.ID, nil
}

func (m Message) updateUserConversation(message entity.Message) {
	cvnList := make([]entity.UserConversationItem, 0)

	_, userCvn, _ := m.userConversationRepo.FindByRelatedUser(message.UserFrom.ID)

	userCvnItem := entity.UserConversationItem{
		DestUser:    message.UserTo,
		LastMessage: message.Content,
		MessageDate: message.CreatedAt,
		Read:        false,
	}

	if userCvn.ID == "" {
		cvnList = append(cvnList, userCvnItem)

		newItem := entity.UserConversation{
			ID:            uuid.New().String(),
			RelatedUser:   message.UserFrom,
			UpdatedAt:     time.Now().UTC(),
			CreatedAt:     time.Now().UTC(),
			Conversations: cvnList,
		}
		_ = m.userConversationRepo.Save(newItem)
		return
	}

	for _, v := range userCvn.Conversations {
		if v.DestUser.ID == message.UserTo.ID {
			v.Read = false
			v.MessageDate = time.Now().UTC()
			v.LastMessage = message.Content
		}
		cvnList = append(cvnList, v)
	}

	_ = m.userConversationRepo.UpdateItem(userCvn.ID, cvnList)
}
