package tokenservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		expired time.Duration
		want    *tokenService
	}{
		{
			name:    "TEST OK",
			key:     "test",
			expired: 10 * time.Second,
			want:    &tokenService{key: "test", expired: 10 * time.Second},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, New(tt.key, tt.expired), tt.want)
		})
	}
}

func TestTokenService(t *testing.T) {
	type fields struct {
		key     string
		expired time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		userID  uint
		wantErr bool
	}{
		{
			name:    "TEST OK userID 1 with key test",
			fields:  fields{key: "test", expired: 10 * time.Second},
			userID:  1,
			wantErr: false,
		},
		{
			name:    "TEST OK userID 2 with key test2",
			fields:  fields{key: "test2", expired: 10 * time.Second},
			userID:  2,
			wantErr: false,
		},
		{
			name:    "TEST OK userID 55 with key test",
			fields:  fields{key: "test", expired: 10 * time.Second},
			userID:  55,
			wantErr: false,
		},
		{
			name:    "TEST OK userID 56 with key test2",
			fields:  fields{key: "test2", expired: 10 * time.Second},
			userID:  56,
			wantErr: false,
		},
		{
			name:    "TEST OK userID 56 with key test2",
			fields:  fields{key: "test2", expired: 10 * time.Second},
			userID:  56,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := New(tt.fields.key, tt.fields.expired)
			token, err := ts.Encrypt(tt.userID)
			print(token)
			assert.Equal(t, tt.wantErr, err != nil)
			userID, err := ts.Decrypt(token)
			assert.Equal(t, tt.userID, userID)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestTokenServiceDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
		userID  uint
	}{
		{
			name:    "TEST error",
			token:   "test",
			wantErr: true,
			userID:  0,
		},
		{
			name:    "TEST with expired token",
			token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQyMjkyODIsIlVzZXJJRCI6NTZ9.n9lhAFN7tx9lF3P3stsuQOxz4Brp_XoqbLXZwh_MR_M",
			wantErr: true,
			userID:  0,
		},
		{
			name:    "TEST with expired token",
			token:   "eyJhbGciOiIxMjMxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQyMzAwODMsIlVzZXJJRCI6NTZ9.V20A3x5H3pOL1ktzA7Ei2eEXD7FkZXCvqdZ70RC7qBU=",
			wantErr: true,
			userID:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := New("", time.Second)
			userID, err := ts.Decrypt(tt.token)
			assert.Equal(t, tt.userID, userID)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
