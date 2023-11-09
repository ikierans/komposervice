package types

import "example/komposervice/internal/schema"

type IAuthService interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
}
