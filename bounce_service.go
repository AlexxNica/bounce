package bounce

import (
	"fmt"
	"github.com/s-rah/go-ricochet"
	"log"
	"net/url"
	"time"
)

type BounceService struct {
	goricochet.StandardRicochetService
	hostname string
	token    *Token
	tokenTtl time.Duration
}

// InitTokenService sets up all the internal token generation logic.
func (bs *BounceService) InitTokenService(hostname string, tokenKey [32]byte, tokenTtl time.Duration) {
	bs.hostname = hostname
	bs.token = new(Token)
	bs.token.Init(tokenKey, tokenTtl)
	bs.tokenTtl = tokenTtl
}

// SendToken, constructs and sends a token to a give ricochet address.
func (bs *BounceService) SendToken(address string) {
	err := bs.Connect(address)
	if err != nil {
		log.Printf("Could not connect to ricochet service:  %v", err)
	}
}

// ValidateToken takes an encrypted token, and the supposed originating ricochet
// address. It validates the encrypted token against the address.
func (bs *BounceService) ValidateToken(token string, address string) bool {
	return bs.token.Validate(token, address)
}

// Always Accept Confirmed Contacts
func (bs *BounceService) IsKnownContact(hostname string) bool {
	return true
}

// OnAuthenticationResult accepted the authentication, and sends messages depending on the state of the connection.
func (bs *BounceService) OnAuthenticationResult(oc *goricochet.OpenConnection, channelID int32, result bool, isKnownContact bool) {
	bs.StandardRicochetService.OnAuthenticationResult(oc, channelID, result, isKnownContact)

	if result == false {
		return
	}

	// We create an encrypted token with the address. We will use
	// this later on to validate the login.
	token := bs.token.Generate(oc.OtherHostname)
	tokenEscaped := url.QueryEscape(token)

	tokenMessage := fmt.Sprintf("Visit this link to login: %s/bounce?address=%s&token=%s", bs.hostname, oc.OtherHostname, tokenEscaped)
	ttlMessage := fmt.Sprintf("Your token expires at: %s", bs.tokenTtl)

	if isKnownContact == false {
		oc.SendContactRequest(3, "Bounce", tokenMessage)
	} else {
		oc.OpenChatChannel(5)
		oc.SendMessage(5, tokenMessage)
		oc.SendMessage(5, ttlMessage)
	}
}
