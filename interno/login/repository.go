package login

type RepositoryLogin interface {
	SaveLogin(login *Login) error 
	VerificaEmail(email string) (*Login, error)
}
