package usecase

import (
	"strings"
	"time"

	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/application/helpers"
	"github.com/lea55/BACKJOBIEX/src/application/mappers"
	"github.com/lea55/BACKJOBIEX/src/domain/config"
	"github.com/lea55/BACKJOBIEX/src/domain/platforms"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type Auth struct {
	userUc      *User
	passHelp    *helpers.Password
	tokenHelp   *helpers.AuthToken
	sessionRepo repository.UserSession
	facebook    platforms.Facebook
	google      platforms.Google
	mapper      *mappers.Auth
}

func NewAuth(userRepo repository.User, rolRepo repository.UserRol, config *config.AppConfig,
	sessionRepo repository.UserSession, facebook platforms.Facebook, google platforms.Google,
	pCatRepo repository.ProjectCategoryRepo, devSkillRepo repository.DevSkill) *Auth {
	return &Auth{
		userUc:      NewUser(userRepo, rolRepo, pCatRepo, devSkillRepo),
		passHelp:    helpers.NewPassword(),
		tokenHelp:   helpers.NewAuthToken(config),
		sessionRepo: sessionRepo,
		facebook:    facebook,
		google:      google,
		mapper:      mappers.NewAuth(),
	}
}

func (a Auth) InternalAuth(rq dto.RqAuthentication) (dto.AuthenticatedUser, error) {
	var authenticated dto.AuthenticatedUser

	user, err := a.userUc.FindByEmail(rq.Email)
	if err != nil {
		return authenticated, errors.New("Usuario y/o contraseña incorrectos")
	}

	_, err = a.passHelp.Validate(rq.Password, user.Password)
	if err != nil {
		return authenticated, errors.New("Usuario y/o contraseña incorrectos")
	}

	token, err := a.tokenHelp.GenerateUserToken(user)
	if err != nil {
		return authenticated, err
	}

	newSession := a.mapper.BuildSessionToSave(rq.CnnIP, token, user)

	sessionID, err := a.sessionRepo.Save(newSession)
	if err != nil {
		return authenticated, errors.Wrap(err, "Error en almacenamiento de sesión")
	}

	err = a.userUc.UpdateDeviceID(user.ID, rq.DeviceID)
	if err != nil {
		return authenticated, err
	}

	authenticated = a.mapper.BuildUser(user, token, sessionID)

	return authenticated, nil
}

//RefreshToken Validaciones segun requerimientos para un inicio de sesión
func (a Auth) RefreshToken(rq dto.RqRefreshToken) dto.AuthenticatedUser {
	var authenticated dto.AuthenticatedUser

	user, err := a.userUc.FindByEmail(rq.Email)
	if err != nil {
		return authenticated
	}

	session, err := a.sessionRepo.FindByID(rq.TokenID)
	if err != nil {
		return authenticated
	}

	if user.DeviceID != rq.DeviceID || rq.TokenID != session.TokenID {
		return authenticated
	}

	token, err := a.tokenHelp.GenerateUserToken(user)
	if err != nil {
		return authenticated
	}

	newSession := a.mapper.BuildSessionToSave(rq.CnnIP, token, user)
	sessionID, err := a.sessionRepo.Save(newSession)
	if err != nil {
		return authenticated
	}

	err = a.userUc.UpdateDeviceID(user.ID, rq.DeviceID)
	if err != nil {
		return authenticated
	}

	authenticated = a.mapper.BuildUser(user, token, sessionID)
	return authenticated
}

//SocialMediaAuth Punto de entrada para autenticación con redes sociales
func (a Auth) SocialMediaAuth(model dto.RqSocialMediaAuth) (dto.AuthenticatedUser, error) {
	var validUser dto.AuthenticatedUser
	var err error

	if model.Type == "facebook" {
		validUser, err = a.validateFacebookUser(model)
		if err != nil {
			return validUser, err
		}
	}

	if model.Type == "google" {
		validUser, err = a.validateGoogleUser(model)
		if err != nil {
			return validUser, err
		}
	}

	_ = a.userUc.UpdateImage(validUser.ID, model.Image)

	err = a.userUc.UpdateDeviceID(validUser.ID, model.DeviceID)
	if err != nil {
		return validUser, err
	}

	return validUser, nil
}

//validateFacebookUser valida datos enviados por cliente en la plataforma facebook para saber si la autenticación fue
//exitosa
func (a Auth) validateFacebookUser(rq dto.RqSocialMediaAuth) (dto.AuthenticatedUser, error) {
	var validUser dto.AuthenticatedUser

	err := a.facebook.ValidateUser(rq.PlatformToken, rq.PlatformUserID)
	if err != nil {
		return validUser, err
	}

	user, err := a.userUc.FindByEmail(rq.Email)
	if err != nil {
		newUserName := strings.Split(rq.Email, "@")
		err = a.userUc.SimpleUserRegistration(dto.RqRegisterUser{
			Names:    rq.Names,
			Surnames: "",
			Email:    rq.Email,
			Password: time.Now().String(),
			UserName: newUserName[0],
		})

		if err != nil {
			return validUser, err
		}

		user, err = a.userUc.FindByEmail(rq.Email)
		if err != nil {
			return validUser, err
		}
	}

	token, err := a.tokenHelp.GenerateUserToken(user)
	if err != nil {
		return validUser, err
	}

	newSession := a.mapper.BuildSessionToSave(rq.CnnIP, token, user)
	sessionID, err := a.sessionRepo.Save(newSession)
	if err != nil {
		return validUser, err
	}

	validUser = a.mapper.BuildUser(user, token, sessionID)
	return validUser, nil
}

//validateGoogleUser valida datos enviados por cliente en la plataforma platforms.Google para saber si la autenticación
//fue exitosa
func (a Auth) validateGoogleUser(rq dto.RqSocialMediaAuth) (dto.AuthenticatedUser, error) {
	var validUser dto.AuthenticatedUser

	err := a.google.ValidateUser(rq.PlatformToken, rq.PlatformUserID)
	if err != nil {
		return validUser, err
	}

	user, err := a.userUc.FindByEmail(rq.Email)
	if err != nil {
		newUserName := strings.Split(rq.Email, "@")
		err = a.userUc.SimpleUserRegistration(dto.RqRegisterUser{
			Names:    rq.Names,
			Surnames: "",
			Email:    rq.Email,
			Password: time.Now().String(),
			UserName: newUserName[0],
		})

		if err != nil {
			return validUser, err
		}

		user, err = a.userUc.FindByEmail(rq.Email)
		if err != nil {
			return validUser, err
		}
	}

	token, err := a.tokenHelp.GenerateUserToken(user)
	if err != nil {
		return validUser, err
	}

	newSession := a.mapper.BuildSessionToSave(rq.CnnIP, token, user)
	sessionID, err := a.sessionRepo.Save(newSession)
	if err != nil {
		return validUser, err
	}

	validUser = a.mapper.BuildUser(user, token, sessionID)
	return validUser, nil
}
