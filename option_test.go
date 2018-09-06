package sophos_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esurdam/go-sophos"
)

func TestAutoResolveErrsMode(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"withResolve", args{r: httptest.NewRequest("GET", "/api", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sophos.AutoResolveErrsMode(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("AutoResolveErrsMode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.args.r.Header.Get(sophos.XRestdErrAck) != "all" {
					t.Errorf("AutoResolveErrsMode() error = %v, wantErr %v", "header should be all", tt.wantErr)
				}
			}
		})
	}
}

func TestCancelResolveErrsMode(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"withCancel", args{r: httptest.NewRequest("GET", "/api", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sophos.CancelResolveErrsMode(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CancelResolveErrsMode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.args.r.Header.Get(sophos.XRestdErrAck) != "none" {
					t.Errorf("CancelResolveErrsMode() error = %v, wantErr %v", "header should be none", tt.wantErr)
				}
			}
		})
	}
}

func TestWithRestdLockOverride(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"withLockOverride", args{r: httptest.NewRequest("GET", "/api", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sophos.WithRestdLockOverride(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("WithRestdLockOverride() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.args.r.Header.Get(sophos.XRestdLockOverride) != "yes" {
					t.Errorf("WithRestdLockOverride() error = %v, wantErr %v", "header should be yes", tt.wantErr)
				}
			}
		})
	}
}

func TestWithRestdInsert(t *testing.T) {

	type args struct {
		rule     string
		position int
	}
	tests := []struct {
		name      string
		args      args
		want      string
		shouldErr bool
	}{
		{"withRestdInsert", args{rule: "packetfilter.rules", position: 4}, "packetfilter.rules 4", false},
		{"withRestdInsertNoPos", args{rule: "packetfilter.rules", position: 0}, "packetfilter.rules", false},
		{"withRestdInsertErr", args{rule: "", position: 4}, " 4", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api", nil)
			opt := sophos.WithRestdInsert(tt.args.rule, tt.args.position)
			err := opt(r)

			if err == nil && tt.shouldErr {
				t.Errorf("WithRestdInsert() = should error, want %v", tt.want)
			}

			if err != nil && !tt.shouldErr {
				t.Errorf("WithRestdInsert() = should not error, %v, want %v", err, tt.want)
			}

			if err != nil && tt.shouldErr {
				return
			}

			if header := r.Header.Get(sophos.XRestdInsert); header != tt.want {
				t.Errorf("WithRestdInsert() = %v, want %v", header, tt.want)
			}
		})
	}
}

func TestWithSessionClose(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"withRestdSessionClose", args{r: httptest.NewRequest("GET", "/api", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sophos.WithSessionClose(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("WithSessionClose() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.args.r.Header.Get(sophos.XRestdSession) != "close" {
					t.Errorf("WithSessionClose() error = %v, wantErr %v", "header should be yes", tt.wantErr)
				}
			}
		})
	}
}

func TestWithBasicAuth(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"withBasicAuth", args{username: "abc123", password: ""}, "Basic YWJjMTIzOg==", false},
		{"withBasicAuth2", args{username: "abc123", password: "abc123"}, "Basic YWJjMTIzOmFiYzEyMw==", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api", nil)
			opt := sophos.WithBasicAuth(tt.args.username, tt.args.password)
			err := opt(r)
			if err == nil && tt.wantErr {
				t.Errorf("TestWithBasicAuth() = should error, want %v", tt.wantErr)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("TestWithBasicAuth() = should not error, want %v", tt.wantErr)
			}

			if err != nil && tt.wantErr {
				return
			}

			if header := r.Header.Get(sophos.Authorization); header != tt.want {
				t.Errorf("TestWithBasicAuth() = %v, want %v", header, tt.want)
			}
		})
	}
}

func TestWithApiToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"withApiToken", args{token: "abc123"}, "Basic dG9rZW46YWJjMTIz", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api", nil)
			opt := sophos.WithApiToken(tt.args.token)
			err := opt(r)
			if err == nil && tt.wantErr {
				t.Errorf("TestWithApiToken() = should error, want %v", tt.wantErr)
			}
			if err != nil && !tt.wantErr {
				t.Errorf("TestWithApiToken() = should not error, want %v", tt.wantErr)
			}

			if err != nil && tt.wantErr {
				return
			}

			if header := r.Header.Get(sophos.Authorization); header != tt.want {
				t.Errorf("TestWithApiToken() = %v, want %v", header, tt.want)
			}
		})
	}
}
