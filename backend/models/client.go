package models

type Client struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Telephone      string `json:"telephone"`
	Email          string `json:"email"`
	DataNascimento string `json:"data_nascimento"`
	Sexo           string `json:"sexo"`
	DataCadastro   string `json:"data_cadastro"`
	Observacoes    string `json:"observacoes"`
}
