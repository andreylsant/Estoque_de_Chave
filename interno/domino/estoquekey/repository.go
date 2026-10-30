package estoquekey

type Repository interface{
	Save(chave *Chave) error
	Upadete(chave *Chave) error
	Get() ([]Chave, error)
	GetKeyID(id string) (*Chave, error)
	DeleteID(id string) error
}