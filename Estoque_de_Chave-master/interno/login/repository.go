package login

type RepositoryLogin interface{
	SaveLogin(login *login) error 
}