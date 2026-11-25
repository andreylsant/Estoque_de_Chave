package gormrepo

type Chave struct {
	Id         string `json:"id"`
	Marca      string `json:"marca"`
	Modelo     string `json:"modelo"`
	Quantidade int    `json:"quantidade"`
}

type Login struct {
	Id    string `json:"id"`
	Email string `josn:"email"`
	Senha string `json:"senha"`
}
