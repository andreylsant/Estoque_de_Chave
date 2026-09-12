package handle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreylsant/estoque_de_chave/interno/login"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLogin struct {
	mock.Mock
}

func (m *MockLogin) SaveLogin(login *login.Login) error {
	a := m.Called(login)
	return a.Error(0)
}

func (m *MockLogin) VerificaEmail(email string) (*login.Login, error) {
	return &login.Login{}, nil
}

type MockToken struct {
	mock.Mock
}

func (m *MockToken) GerarToken(userID string) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

var (
	loginUse = login.Login{
		Email: "email@chipcode.com",
		Senha: "1234567",
	}
)

/*
	Criar tests golang colocar minhas depêndencias
		handler := HandleKey{
		LoginService: login.LoginService{
			RepositoryLogin: repositoryMock,
			Token:           repositoryToken,
		},
	}

	Se tiver utilizando o pacote gin deve deve se modelo do gin que ira utilizar 
	gin.SetMode(gin.TestMode)
*/
func Test_PostLogin_CriandoNovoLogin(t *testing.T) {
	assert := assert.New(t)
	gin.SetMode(gin.TestMode)

	jsonValou, _ := json.Marshal(loginUse)
	CorpoDarequest := bytes.NewBuffer(jsonValou)

	repositoryMock := new(MockLogin)
	repositoryToken := new(MockToken)

	repositoryMock.On("SaveLogin", mock.Anything).Return(nil)
	repositoryToken.On("GerarToken", mock.Anything).Return("meu-token-jwt", nil)

	// 3. Montando a estrutura do Handler com a dependência mockada
	handler := HandleKey{
		LoginService: login.LoginService{
			RepositoryLogin: repositoryMock,
			Token:           repositoryToken,
		},
	}

	recorder := httptest.NewRecorder()
	context, router := gin.CreateTestContext(recorder)

	// Registra a rota apontando para o método do seu handler
	router.POST("/login", handler.Post_Login)

	req, _ := http.NewRequest("POST", "/login", CorpoDarequest)
	req.Header.Set("Content-Type", "application/json") // Avisa ao Gin que é um JSON

	// Executa a requisição passando o contexto
	context.Request = req
	router.ServeHTTP(recorder, req)

	// 5. Validações do resultado HTTP
	assert.Equal(http.StatusOK, recorder.Code, "O status de retorno deve ser 200 OK")

	// Captura e valida o JSON que o seu controller respondeu
	var respostaMap map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &respostaMap)
	
	assert.NoError(err)
	assert.Equal("meu-token-jwt", respostaMap["token"])
	assert.Equal("Login recebido", respostaMap["mensagem"])

}
