package bounce

import (
	"strconv"
	"strings"
	"time"
)

type Token struct {
	ic  Crypto
	ttl time.Duration
}

// Init produces a Token which encrypts tokens, with
// the given key, that have a ttl matching the one given.
func (token *Token) Init(key [32]byte, ttl time.Duration) {
	token.ic.Init(key)
	token.ttl = ttl
}

// Generate produces an encrpyted Token token comprising
// the discriminator and a timestamp representing when the token
// will expire.
// discriminator should be unique to the given user/session that is being checked.
func (token *Token) Generate(discriminator string) string {
	expiry := time.Now().Add(token.ttl).Unix()
	return token.ic.Encrypt(strconv.Itoa(int(expiry)) + "|" + discriminator)
}

// Validate decrypts encryptedToken and checks that it contains the
// given discriminator, and that the token is valid given the ttl, if all
// are true the function returns true, otherwise false.
func (token *Token) Validate(encryptedToken string, discriminator string) bool {
	plain, success := token.ic.Decrypt(encryptedToken)

	if !success {
		return false
	}

	s := strings.Split(plain, "|")
	if len(s) != 2 || s[1] != discriminator {
		return false
	}

	expiry, err := strconv.ParseInt(s[0], 0, 64)
	if err != nil {
		return false
	}

	if expiry < time.Now().Unix() {
		return false
	}
	return true
}
