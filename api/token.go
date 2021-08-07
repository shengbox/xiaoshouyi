package api

import (
	"sync"
	"time"
)

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	IssuedAt    int64  `json:"issued_at"`
}

type Tokener struct {
	mutex *sync.RWMutex
	TokenInfo
	tokenFetcher func() (token string, expiresIn int64)
}

func (t *Tokener) Token() string {
	if !t.isValidToken() {
		t.mutex.Lock()
		defer t.mutex.Unlock()
		accessToken, issuedAt := t.tokenFetcher()
		t.AccessToken = accessToken
		t.IssuedAt = issuedAt
	}
	return t.AccessToken
}

func (t *Tokener) RefreshToken() error {
	accessToken, expiresIn := t.tokenFetcher()
	t.IssuedAt = expiresIn
	t.AccessToken = accessToken
	return nil
}

func (t *Tokener) isValidToken() bool {
	now := time.Now().Unix() * 1000
	if now >= t.IssuedAt || t.AccessToken == "" {
		return false
	}
	return true
}
