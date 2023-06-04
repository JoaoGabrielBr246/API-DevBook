package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //omitempty -> se tiver em branco o id, não vai passar para o json
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email, omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm, omitempty"`
}
