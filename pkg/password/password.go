package password

import "golang.org/x/crypto/bcrypt"

// Hasher interface
type Hasher interface {
	Hash(password string) string
	Compare(plainPassword, hashPassword string) bool
}

// BcryptHash struct
type BcryptHash struct{}

// NewBCryptHash return Hasher interface
func NewBCryptHash() Hasher {
	return &BcryptHash{}
}

// Hash Password function
func (b *BcryptHash) Hash(password string) string {
	hb, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hb)
}

// Compare Password function
func (b *BcryptHash) Compare(plainPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}
