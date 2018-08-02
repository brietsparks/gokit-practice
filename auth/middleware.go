package auth

import (
	"encoding/json"
	"net/http"
	"errors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"github.com/codegangsta/negroni"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type Env struct {
	Aud          string
	Iss          string
	JwksEndpoint string
}

func NewJwtChecker(env Env) *jwtmiddleware.JWTMiddleware {
	getValidationKey := makeGetValidationKey(env)

	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: getValidationKey,
		SigningMethod:       jwt.SigningMethodRS256,
	})
}

func NewAuthenticator(logger log.Logger) negroni.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		//user := req.Context().Value("user")
		//sub, _ := user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
		next.ServeHTTP(writer, req)
	}
}

func makeGetValidationKey(env Env) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(env.Aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(env.Iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token, env.JwksEndpoint)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}
}

func getPemCert(token *jwt.Token, jwksEndpoint string) (string, error) {
	cert := ""
	resp, err := http.Get(jwksEndpoint)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
