package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLogin struct{
	mock.Mock
}

func (m *MockLogin) SaveLogin(login *Login) error {
	a:= m.Called(login)
	return a.Error(0)
}

func (m *MockLogin) FindByEmail(email string) (*Login, error){
	return &Login{}, nil
}

func Test_CreateLogin(t *testing.T) {
	assert:=  assert.New(t)
	//instanciando meu login 
	login:= Login{
		Email: "email@gmail.com",
		Senha: "1234567",
	}
	//
	repositoryMock:= new(MockLogin)
	service:= LoginService{
		RepositoryLogin: repositoryMock,
	}

	err:= service.CreateLogin(&login)

	assert.Nil(err)

}

func Test_CreateLogin_SaveLogin(t *testing.T) {
	assert:=  assert.New(t)
	//instanciando meu login 
	login:= Login{
		Email: "email@gmail.com",
		Senha: "1234567",
	}
	//
	repositoryMock:= new(MockLogin)

	repositoryMock.On("SaveLogin", mock.AnythingOfType("*login")).Return(nil)

	service:= LoginService{
		RepositoryLogin: repositoryMock,
	}

	err:= service.CreateLogin(&login)

	assert.NoError(err)

	repositoryMock.AssertExpectations(t)

}