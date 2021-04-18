package service

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rGustave07/fbhcapi/api/model"
	"github.com/stretchr/testify/assert"
)

func TestNewPairFromUser(t *testing.T) {
	priv, _ := ioutil.ReadFile("../rsa_private_test.pem")
	privKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priv)

	pub, _ := ioutil.ReadFile("../rsa_public_test.pem")
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pub)

	secret := "TestSecret"

	tokenService := NewTokenService(&TSConfig{
		PrivKey:       privKey,
		PubKey:        pubKey,
		RefreshSecret: secret,
	})

	uid, _ := uuid.NewRandom()
	u := &model.User{
		UID:      uid,
		Email:    "bob@loblaw.com",
		Password: "testPass",
	}

	t.Run("Returns a token pair with values", func(t *testing.T) {
		ctx := context.TODO()
		tokenPair, err := tokenService.NewPairFromUser(ctx, u, "")
		assert.NoError(t, err)

		var s string
		assert.IsType(t, s, tokenPair.IDToken)

		// Decode base64URL encoded string
		// simpler to use JWT Library
		idTokenClaims := &IDTokenCustomClaims{}

		_, err = jwt.ParseWithClaims(tokenPair.IDToken, idTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})

		assert.NoError(t, err)

		// assert claims on idToken
		expectedClaims := []interface{}{
			u.UID,
			u.Email,
			u.Firstname,
			u.Lastname,
			u.ImageURL,
			u.Website,
		}
		actualIDClaims := []interface{}{
			idTokenClaims.User.UID,
			idTokenClaims.User.Email,
			idTokenClaims.User.Firstname,
			idTokenClaims.User.Lastname,
			idTokenClaims.User.ImageURL,
			idTokenClaims.User.Website,
		}

		assert.ElementsMatch(t, expectedClaims, actualIDClaims)
		assert.Empty(t, idTokenClaims.User.Password)

		idTokenExpiresAt := time.Unix(idTokenClaims.StandardClaims.ExpiresAt, 0)
		expectedExpiresAt := time.Now().Add(15 * time.Minute)

		assert.WithinDuration(t, expectedExpiresAt, idTokenExpiresAt, 5*time.Second)

		refreshTokenClaims := &RefreshTokenCustomClaims{}
		_, err = jwt.ParseWithClaims(tokenPair.RefreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		assert.IsType(t, s, tokenPair.RefreshToken)
		assert.NoError(t, err)
		assert.Equal(t, u.UID, refreshTokenClaims.UID)

		refreshTokenExpiresAt := time.Unix(refreshTokenClaims.StandardClaims.ExpiresAt, 0)
		expectedExpiresAt = time.Now().Add(3 * 24 * time.Hour)

		assert.WithinDuration(t, expectedExpiresAt, refreshTokenExpiresAt, 5*time.Second)
	})
}
