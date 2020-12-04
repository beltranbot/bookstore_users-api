package services

import (
	"github.com/beltranbot/bookstore_users-api/domain/users"
	"github.com/beltranbot/bookstore_users-api/utils/cryptoutils"
	"github.com/beltranbot/bookstore_users-api/utils/dateutils"
	"github.com/beltranbot/bookstore_utils-go/resterrors"
)

var (
	// UsersService instance
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	Get(int64) (*users.User, *resterrors.RestErr)
	Create(users.User) (*users.User, *resterrors.RestErr)
	Update(bool, users.User) (*users.User, *resterrors.RestErr)
	Delete(int64) *resterrors.RestErr
	Search(string) (users.Users, *resterrors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *resterrors.RestErr)
}

// Get func
func (s *usersService) Get(userID int64) (*users.User, *resterrors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// Create func
func (s *usersService) Create(user users.User) (*users.User, *resterrors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateutils.GetNowDBFormat()
	user.Password = cryptoutils.GetMD5(user.Password)
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

// Update func
func (s *usersService) Update(isPartial bool, user users.User) (*users.User, *resterrors.RestErr) {
	current, err := s.Get(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

// Delete func
func (s *usersService) Delete(userID int64) *resterrors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search func
func (s *usersService) Search(status string) (users.Users, *resterrors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *resterrors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: cryptoutils.GetMD5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
