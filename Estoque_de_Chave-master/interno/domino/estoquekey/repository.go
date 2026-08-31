package estoquekey

type Repository interface{
	Save(key *Key) error
	Upadete(key *Key)
	Get() (*Key, error)
	GetKeyID(id string) *Key
}