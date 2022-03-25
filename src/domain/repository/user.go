package repository

import "BACKJOBIEX/src/domain/entity"

type User interface {
	Save(user entity.User) (string, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID string) (entity.User, error)
	UpdateImage(userID string, image string) error
	UpdateDeviceID(userID string, deviceID string) error
	FindPaginated(from uint32, limit uint32) ([]entity.User, error)
	FindByNickname(nickname string) (entity.User, error)
	UpdateCoverImage(image string, ID string) error
	UpdateProfileImage(image string, ID string) error
	UpdateProfile(userID string, names string, surnames string, category entity.ProjectCategory, hourCost uint32, nickName string) error
	UpdateLocation(userID string, city string, country string) error
	UpdateEducation(userID string, values []entity.UserEducation) error
	UpdateSkills(userID string, values []entity.DevSkill) error
}
