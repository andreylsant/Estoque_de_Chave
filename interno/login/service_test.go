package login

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLogin struct {
	mock.Mock
}

func (m *MockLogin) SaveLogin(login *Login) error {
	a := m.Called(login)
	return a.Error(0)
}

func (m *MockLogin) FindByEmail(email string) (*Login, error) {
	return &Login{}, nil
}

type MockToken struct {
	mock.Mock
}

func (m *MockToken) GerarToken(userID string) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

var (
		login = Login{
		Email: "email@chipcode.com",
		Senha: "1234567",
	}
)

func Test_CreateLogin_Sucesso(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(MockLogin)
	repositoryToken := new(MockToken)

	repositoryMock.On("SaveLogin", mock.Anything).Return(nil)
	repositoryToken.On("GerarToken", mock.Anything).Return("meu-token-jwt",nil)

	service := LoginService{
		RepositoryLogin: repositoryMock,
		Token:           repositoryToken,
	}

	tokenResultado, err := service.CreateLogin(&login)

	assert.NoError(err, "Não deveria retornar erro em cenário de sucesso")
	assert.Equal("meu-token-jwt", tokenResultado, "O token retornado deve ser o mesmo gerado pelo componente de Token")

	repositoryMock.AssertExpectations(t)
	repositoryToken.AssertExpectations(t)
}

func Test_CreateLogin_ErrorAoSalvarNoBanco(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(MockLogin)
	repositoryToken := new(MockToken)

	// Simulando erro no banco de dados
	repositoryMock.On("SaveLogin", mock.Anything).Return(errors.New("db error"))

	service := LoginService{
		RepositoryLogin: repositoryMock,
		Token:           repositoryToken,
	}

	tokenResultado, err := service.CreateLogin(&login)

	// Validações
	assert.Error(err)
	assert.Equal("[Error ao Salvar login]", err.Error(), "Deve retornar a mensagem customizada de erro do banco")
	assert.Empty(tokenResultado, "O token deve ser vazio se houver erro")

}
