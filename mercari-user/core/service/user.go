package service

import (
	"context"
	"fmt"
	"sagungw/mercari/core/datastore"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Register(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (string, error)
	Authorize(ctx context.Context, token string) (*datastore.User, bool)
}

type userService struct {
	userDatastore         datastore.UserDatastore
	loginHistoryDatastore datastore.HistoryDatastore
	secret                []byte
}

func NewUserService(userDatastore datastore.UserDatastore, loginHistoryDatastore datastore.HistoryDatastore) *userService {
	return &userService{
		userDatastore:         userDatastore,
		loginHistoryDatastore: loginHistoryDatastore,
		secret:                []byte("itsasecret"),
	}
}

func (u *userService) Register(ctx context.Context, email, password string) error {
	return u.userDatastore.Register(ctx, email, password)
}

func (u *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.userDatastore.Login(ctx, email, password)
	if err != nil {
		fmt.Println("SINI")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID.Hex(),
		"user_email": user.Email,
	})

	tokenString, err := token.SignedString(u.secret)
	if err != nil {
		fmt.Println("SINIA", err)
		return "", err
	}

	err = u.loginHistoryDatastore.SaveLoginHistory(ctx, user.ID.Hex(), time.Now())
	if err != nil {
		fmt.Println("SINIB")
		return "", err
	}

	return tokenString, nil
}

func (u *userService) Authorize(ctx context.Context, tokenString string) (*datastore.User, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return u.secret, nil
	})
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, _ := primitive.ObjectIDFromHex(claims["user_id"].(string))
		return &datastore.User{
			ID:    userID,
			Email: claims["user_email"].(string),
		}, true
	}

	return nil, false
}
