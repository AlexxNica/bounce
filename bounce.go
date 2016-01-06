package bounce

import (
	"github.com/s-rah/go-ricochet"
	"time"
)

// Bounce is a prototype authentication system to allow indivudal logins
// to be backed by ricochet identities.
type Bounce struct {
	keyfile  string
	hostname string
	debugLog bool
	token    *Token
}

// Init sets up the Bounce Object.
func (b *Bounce) Init(keyfile string, hostname string, tokenKey [32]byte, tokenTtl time.Duration, debugLog bool) {
	b.keyfile = keyfile
	b.hostname = hostname
	b.debugLog = debugLog

	b.token = new(Token)
	b.token.Init(tokenKey, tokenTtl)
}

// SendToken, constructs and sends a token to a give ricochet address.
func (b *Bounce) SendToken(address string) {

	// We create an encrypted token with the address. We will use
	// this later on to validate the login.
	token := b.token.Generate(address)

	ricochet := new(goricochet.Ricochet)
	ricochet.Init(b.keyfile, b.debugLog)
	err := ricochet.Connect(b.hostname, address)

	if err != nil {
		return
	}

	if ricochet.IsKnownContact() == false {
		ricochet.SendContactRequest("Bounce", "Your token is: "+token)
	} else {
		go ricochet.ListenAndWait()
		ricochet.OpenChannel("im.ricochet.chat", 5)
		ricochet.SendMessage("Your token is: "+token, 5)
	}
}

// ValidateToken takes an encrypted token, and the supposed originating ricochet
// address. It validates the encrypted token against the address.
func (b *Bounce) ValidateToken(token string, address string) bool {
	return b.token.Validate(token, address)
}
