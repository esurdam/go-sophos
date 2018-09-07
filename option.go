package sophos

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
)

// Option is a functional config that is used to modify outgoing requests.
type Option func(r *http.Request) error

// AutoResolveErrsMode is an Option which sets the X-Restd-Err-Ack header to 'all' which will auto resolve all
// non-fatal errors. e.g. deleting sub-dependencies.
func AutoResolveErrsMode(r *http.Request) error {
	r.Header.Set(http.CanonicalHeaderKey(XRestdErrAck), "all")
	return nil
}

// CancelResolveErrsMode is an Option which sets the X-Restd-Err-Ack header to 'none' which will automatically acknowledge
// the error and cancel the operation. Use this setting only when you’re not deleting important data or objects.
func CancelResolveErrsMode(r *http.Request) error {
	r.Header.Set(http.CanonicalHeaderKey(XRestdErrAck), "none")
	return nil
}

// WithAPIToken is an Option which sets the Authorization header to the provided token
func WithAPIToken(token string) Option {
	return func(r *http.Request) error {
		r.Header.Set(Authorization, "Basic "+base64.StdEncoding.EncodeToString([]byte("token:"+token)))
		return nil
	}
}

// WithBasicAuth is an Option which sets the Authorization header to the provided username and password.
func WithBasicAuth(username, password string) Option {
	return func(r *http.Request) error {
		r.Header.Set(Authorization, "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
		return nil
	}
}

// WithContext is an Option which sets the provided context to the the client's request.
func WithContext(ctx context.Context) Option {
	return func(r *http.Request) error {
		*r = *r.WithContext(ctx)
		return nil
	}
}

// WithRestdInsert is an Option which adds the XRestdInsert header to insert a reference at the given position inside the node.
// X-Restd-Insert: packetfilter.rules 4
func WithRestdInsert(rule string, position int) Option {
	return func(r *http.Request) error {
		if rule == "" {
			return fmt.Errorf("invalid rule name for X-Restd-Insert header: %s", rule)
		}
		if position == 0 {
			r.Header.Set(http.CanonicalHeaderKey(XRestdInsert), fmt.Sprintf("%s", rule))
		} else {
			r.Header.Set(http.CanonicalHeaderKey(XRestdInsert), fmt.Sprintf("%s %d", rule, position))
		}
		return nil
	}
}

// WithRestdLockOverride is an Option which sets the X-Restd-Lock-Override to yes.
func WithRestdLockOverride(r *http.Request) error {
	r.Header.Set(http.CanonicalHeaderKey(XRestdLockOverride), "yes")
	return nil
}

// WithSessionClose is an Option which sets the X-Restd-Session to close, be sure to only send this command with the last request.
func WithSessionClose(r *http.Request) error {
	r.Header.Set(http.CanonicalHeaderKey(XRestdSession), "close")
	return nil
}

func evaluateOpts(r *http.Request, oo []Option) error {
	for _, o := range oo {
		if err := o(r); err != nil {
			return err
		}
	}
	return nil
}
