package token

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := New(tt.fields.key, tt.fields.expired)
			token, err := ts.Crypt(tt.userID)
			userId, err := ts.Decrypt(token)
			assert.Equal(t, tt.userID, userId)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
