package login

type RepositoryLogin interface {
	SaveLogin(login *Login) error
	FindByEmail(email string) (*Login, error)
}
