package jwe

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var privateKey *rsa.PrivateKey
var signingKey jose.SigningKey

type (
	PrivateClaims struct {
		UserID int `json:"user_id"`
	}
)

func (c PrivateClaims) Validate(e PrivateClaims) error {
	if e.UserID == 0 {
		return errors.New("No credentials attached in request")
	}
	return nil
}

func Encode(pri PrivateClaims) (token string, err error) {
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	signingKey = jose.SigningKey{Algorithm: jose.RS256, Key: privateKey}

	signerOpts := (&jose.SignerOptions{}).WithContentType("JWT")
	rsaSigner, err := jose.NewSigner(signingKey, signerOpts)
	if err != nil {
		return
	}

	enc, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{
			Algorithm: jose.RSA_OAEP,
			Key:       &privateKey.PublicKey,
		},
		(&jose.EncrypterOptions{}).WithType("JWT").WithContentType("JWT"))
	if err != nil {
		return
	}
	t := time.Now()
	pub := jwt.Claims{
		Expiry:    jwt.NewNumericDate(t.Add(24 * 30 * 12)),
		IssuedAt:  jwt.NewNumericDate(t),
		NotBefore: jwt.NewNumericDate(t),
	}
	token, err = jwt.SignedAndEncrypted(rsaSigner, enc).Claims(pub).Claims(pri).CompactSerialize()
	if err != nil {
		return
	}
	return
}

func Decode(token string) (pub jwt.Claims, pri PrivateClaims, err error) {
	var parsed *jwt.NestedJSONWebToken
	parsed, err = jwt.ParseSignedAndEncrypted(token)
	if err != nil {
		return
	}

	decrypted, err := parsed.Decrypt(privateKey)
	if err != nil {
		return
	}

	if err = decrypted.Claims(&privateKey.PublicKey, &pub, &pri); err != nil {
		return
	}
	return
}
