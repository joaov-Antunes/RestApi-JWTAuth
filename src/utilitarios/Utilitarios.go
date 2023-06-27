package utilitarios

import (
	"crypto/sha256"
	"fmt"
)

func GerarHash(s string) string {
	encriptador := sha256.New()
	encriptador.Write([]byte(s))

	senhaEncript := encriptador.Sum(nil)

	senha := fmt.Sprintf("%x", senhaEncript)

	return senha
}
