package user

import "golang.org/x/crypto/bcrypt"

type Service interface{
	RegisterUser(inp RegisterUserInput) (User, error)
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