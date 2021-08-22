package scalr

import "context"

// Compile-time proof of interface implementation.
var _ Users = (*users)(nil)

// User represents a Scalr user.
type User struct {
	ID       string `jsonapi:"primary,users"`
	Email    string `jsonapi:"attr,email"`
	FullName string `jsonapi:"attr,fullname,omitempty"`
}

type Users interface {
	// ReadCurrent Gets details of the currently authenticated user.
	ReadCurrent(ctx context.Context) (*User, error)
}

// users implements Users.
type users struct {
	client *Client
}

func (s *users) ReadCurrent(ctx context.Context) (*User, error) {
	req, err := s.client.newRequest("GET", "me", nil)
	if err != nil {
		return nil, err
	}

	u := &User{}
	err = s.client.do(ctx, req, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
