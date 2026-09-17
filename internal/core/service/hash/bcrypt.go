package core_service_hash

import "golang.org/x/crypto/bcrypt"

type Hasher interface{
	Hash(password string) (string, error)
	Compare(hash string, password string) error
}

type BCryptHasher struct {
	times int
}

func NewBCryptHasher(times int) *BCryptHasher{
	return &BCryptHasher{
		times: times,
	}
}

func (h *BCryptHasher) Hash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), h.times)
	return string(b), err
}

func (h *BCryptHasher) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}