package security

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificarSenha(senhaHash, senhaString string) error {
	// TODO: tratar a mensagem de saída de erro de senhas
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
