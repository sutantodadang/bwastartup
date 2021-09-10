package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface{
	RegisterUser(inp RegisterUserInput) (User, error)
	LoginUser(inp LoginInput) (User, error)
	IsEmailAvail(inp CheckEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(inp RegisterUserInput) (User, error)  {
	user := User{}
	user.Name = inp.Name
	user.Email = inp.Email
	user.Occupation = inp.Occupation
	user.Role = "user"

	passwordHash, err  := bcrypt.GenerateFromPassword([]byte(inp.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)


	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}
	

	return newUser, nil
}

func (s *service) LoginUser(inp LoginInput) (User,error)  {
	email := inp.Email
	pass := inp.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(pass))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvail(inp CheckEmailInput) (bool, error) {
	email := inp.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		return true, nil
	}

	return false, nil
}

 func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {
	 user, err := s.repository.FindById(Id)
	 if err != nil {
		 return user, err
	 }

	 user.AvatarFileName = fileLocation

	 updateUser, err := s.repository.Update(user)

	 if err != nil {
		 return updateUser, err
	 }

	 return updateUser, nil
 } 
