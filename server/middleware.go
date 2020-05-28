package server

import (
	"encoding/json"
	"errors"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var (
	errMiddlewareVerifyAudience = errors.New("server: invalid token audience")
	errMiddlewareVerifyIssuer   = errors.New("server: invalid token issuer")
	errMiddlewareParsePublicKey = errors.New("server: error parsing public key from pem")

	errCertGetJWKS   = errors.New("server: error getting jwks from provider")
	errCertParseJWKS = errors.New("server: error parsing jwks")
	errCertGetKey    = errors.New("server: unable to find key")
)

func newMiddleware(p params) func(http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			audience := p.Auth0APIAudience
			ok := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
			if !ok {
				return token, errMiddlewareVerifyAudience
			}

			issuer := "https://" + p.Auth0Domain
			ok = token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
			if !ok {
				return token, errMiddlewareVerifyIssuer
			}

			cert, err := getPEMCert(issuer, token)
			if err != nil {
				return nil, err
			}

			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, errMiddlewareParsePublicKey
			}

			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	return jwtMiddleware.Handler
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// domain:     forstmeier.auth0.com
// identifier: https://forstmeier.auth0.com/api/v2/
//
// https://auth0.com/docs/quickstart/backend/golang/01-authorization#validate-access-tokens

func getPEMCert(domainURL string, token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(domainURL + "/.well-known/jwks.json")
	if err != nil {
		return cert, errCertGetJWKS
	}
	defer resp.Body.Close()

	jwks := JWKS{}
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return cert, errCertParseJWKS
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errCertGetKey
	}

	return cert, nil
}
