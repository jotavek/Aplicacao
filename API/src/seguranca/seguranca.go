package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca uma hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

}

// VerificarSenha compara uma senha e um hash e verifica se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
	