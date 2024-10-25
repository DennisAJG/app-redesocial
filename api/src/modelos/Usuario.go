package modelos

import "time"

// Usuario reperesenta um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // não deixa que o json insira o id em branco
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}
