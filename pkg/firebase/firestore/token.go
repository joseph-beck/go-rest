package firestore

import (
	"errors"
	"reflect"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

func token(info oauth2.TokenInfo) (*models.Token, error) {
	empty := func(info oauth2.TokenInfo) bool {
		if info == nil {
			return true
		}
		if v := reflect.ValueOf(info); v.IsNil() {
			return true
		}
		return reflect.DeepEqual(info, info.New())
	}(info)

	if empty {
		return nil, errors.New("invalid TokenInfo")
	}

	return &models.Token{
		ClientID:         info.GetClientID(),
		UserID:           info.GetUserID(),
		RedirectURI:      info.GetRedirectURI(),
		Scope:            info.GetScope(),
		Code:             info.GetCode(),
		CodeCreateAt:     info.GetCodeCreateAt(),
		CodeExpiresIn:    info.GetCodeExpiresIn(),
		Access:           info.GetAccess(),
		AccessCreateAt:   info.GetAccessCreateAt(),
		AccessExpiresIn:  info.GetAccessExpiresIn(),
		Refresh:          info.GetRefresh(),
		RefreshCreateAt:  info.GetRefreshCreateAt(),
		RefreshExpiresIn: info.GetRefreshExpiresIn(),
	}, nil
}
