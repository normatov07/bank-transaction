package token

import "fmt"

const minSecretKeySize = 12

//JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil,fmt.Errorf("invalid key size: must be at least %d characters",minSecretKeySize)
	}

	return &JWTMaker{secretKey:secretKey}, nil
}

	//CreateToken creates a new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error){

	}

	//VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error){
		
	}