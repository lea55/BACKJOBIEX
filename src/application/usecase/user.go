package usecase

import (
	"fmt"

	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/application/helpers"
	"github.com/lea55/BACKJOBIEX/src/application/mappers"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"
	"github.com/pkg/errors"
)

type User struct {
	repo                repository.User
	rolUc               *Rol
	mapper              *mappers.User
	pass                *helpers.Password
	projectCategoryRepo repository.ProjectCategoryRepo
	devSkillRepo        repository.DevSkill
}

// NewUser constructor de caso de uso de usuario
func NewUser(
	repository repository.User,
	rolRepo repository.UserRol,
	projectCategoryRepo repository.ProjectCategoryRepo,
	devSkillRepo repository.DevSkill,
) *User {
	return &User{
		repo:                repository,
		rolUc:               NewRol(rolRepo),
		mapper:              mappers.NewUser(),
		pass:                helpers.NewPassword(),
		projectCategoryRepo: projectCategoryRepo,
		devSkillRepo:        devSkillRepo,
	}
}

func (u User) UpdateSkills(userID string, devSkillIDs []string) error {
	devSkillList := make([]entity.DevSkill, 0)

	user, err := u.repo.FindByID(userID)
	if err != nil {
		return errors.New("Error en validación de usuario")
	}

	for _, v := range devSkillIDs {
		devSkill, skillErr := u.devSkillRepo.FindByID(v)
		if skillErr != nil {
			return errors.New("No se pudo validar la habilidad ingresada")
		}
		devSkillList = append(devSkillList, devSkill)
	}

	err = u.repo.UpdateSkills(user.ID, devSkillList)
	if err != nil {
		return errors.Wrap(err, "Error al actualizar valores en la base de datos")
	}

	return nil
}

func (u User) UpdateUserEducation(userID string, values []dto.RqUpdateUserEducation) error {
	userEducation := make([]entity.UserEducation, 0)

	user, err := u.repo.FindByID(userID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de usuario")
	}

	for _, v := range values {
		newItem := entity.UserEducation{
			Institution: v.Institution,
			DegreeType:  v.DegreeType,
			DegreeTitle: v.DegreeTitle,
			MonthStart:  v.MonthStart,
			YearStart:   v.YearStart,
			MonthEnd:    v.MonthEnd,
			YearEnd:     v.YearEnd,
		}
		userEducation = append(userEducation, newItem)
	}

	err = u.repo.UpdateEducation(user.ID, userEducation)
	if err != nil {
		return errors.Wrap(err, "Error en la actualización en la base de datos")
	}

	return nil
}

func (u User) UpdateUserProfile(model dto.RqUpdateUserProfile) error {

	user, err := u.repo.FindByID(model.UserID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de usuario")
	}

	cat, err := u.projectCategoryRepo.FindByID(model.CategoryID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de categoría")
	}

	founded, _ := u.repo.FindByNickname(model.NickName)
	if founded.ID != "" {
		return errors.New("No puede usar este nombre de usuario")
	}

	fmt.Println("saiguio el hpo")
	err = u.repo.UpdateProfile(user.ID, model.Names, model.Surnames, cat, model.CostPerHour, model.NickName)
	if err != nil {
		return errors.Wrap(err, "Error en actualización en la base de datos")
	}

	err = u.repo.UpdateLocation(user.ID, model.City, model.Country)
	if err != nil {
		return errors.Wrap(err, "Error en actualización de ubicación")
	}

	return nil

}

func (u User) UpdateProfileImage(imageUrl string, userId string) error {
	user, err := u.repo.FindByID(userId)
	if err != nil {
		return errors.Wrap(err, "Error en validación de usuarios en la base de datos")
	}

	err = u.repo.UpdateProfileImage(imageUrl, user.ID)
	if err != nil {
		return errors.Wrap(err, "Error en la actualización de imagen de perfil en la base de datos")
	}

	return nil
}

func (u User) UpdateCoverImage(imageUrl string, userId string) error {
	user, err := u.repo.FindByID(userId)
	if err != nil {
		return errors.Wrap(err, "Error en validación de usuarios en la base de datos")
	}

	err = u.repo.UpdateCoverImage(imageUrl, user.ID)
	if err != nil {
		return errors.Wrap(err, "Error en la actualización de imagen de portada en la base de datos")
	}

	return nil
}

func (u User) FindById(userID string) (entity.User, error) {
	user, err := u.repo.FindByID(userID)
	if err != nil {
		return user, errors.Wrap(err, "El usuario no existe")
	}
	user.Password = ""
	return user, nil
}

func (u User) FindPaginatedUsers(page uint16) ([]entity.User, error) {
	list := make([]entity.User, 0)

	if page < 0 {
		return list, errors.New("La página enviada no es válida")
	}

	limit := uint32(page * 10)
	from := limit - 10

	list, err := u.repo.FindPaginated(from, limit)
	if err != nil {
		return nil, errors.Wrap(err, "Error en la busqueda en la base de datos")
	}

	newList := make([]entity.User, 0)
	for _, v := range list {
		v.Password = ":)"
		newList = append(newList, v)
	}

	return newList, nil
}

//SimpleUserRegistration Registra en la aplicación un usuario con datos basicos
func (u User) SimpleUserRegistration(doc dto.RqRegisterUser) error {
	founded, _ := u.repo.FindByEmail(doc.Email)
	if founded.ID != "" {
		return errors.New("Ya existe una cuenta creada con este correo")
	}

	founded, _ = u.repo.FindByNickname(doc.UserName)
	if founded.ID != "" {
		return errors.New("Ya existe una cuenta creada con este usuario")
	}

	pass, err := u.pass.Encrypt(doc.Password)
	if err != nil {
		return err
	}

	rol, err := u.rolUc.FindByCode(entity.NewUserRol().User)
	if err != nil {
		return errors.Wrap(err, "Error en validación de rol")
	}

	user := u.mapper.GetToSimpleReg(doc, pass, rol)
	_, err = u.repo.Save(user)
	if err != nil {
		return errors.Wrap(err, "Error en registro en la base de datos")
	}

	return nil
}

//FindByEmail Consulta un entity.User en la base de datos segun un email
func (u User) FindByEmail(email string) (entity.User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return entity.User{}, errors.New("Este usuario no está registrado")
	}

	return user, nil
}

//FindByID Consulta un entity.User en la base de datos segun un ID
func (u User) FindByID(ID string) (entity.User, error) {
	user, err := u.repo.FindByID(ID)
	if err != nil {
		return entity.User{}, errors.New("El usuario no existe en la base de datos")
	}

	return user, nil
}

//UpdateDeviceID Busca un entity.User en la base de datos que coincida con un id y actualiza la propiedad device_id
func (u User) UpdateDeviceID(userID string, deviceID string) error {
	_, err := u.FindByID(userID)
	if err != nil {
		return err
	}

	err = u.repo.UpdateDeviceID(userID, deviceID)
	if err != nil {
		return errors.Wrap(err, "Error en la base de datos al actualizar device_id")
	}

	return nil
}

func (u User) UpdateImage(userID string, image string) error {
	_, err := u.FindByID(userID)
	if err != nil {
		return err
	}

	err = u.repo.UpdateImage(userID, image)
	if err != nil {
		return errors.Wrap(err, "Error en la base de datos al actualizar imagen")
	}

	return nil
}
