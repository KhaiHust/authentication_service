package properties

import "github.com/golibs-starter/golib/config"

type TokenProperties struct {
	PrivateKey          string
	TokenExpired        int
	RefreshTokenExpired int
}

func (t TokenProperties) Prefix() string {
	return "app.services.security.http.jwt"
}
func NewTokenProperties(loader config.Loader) (*TokenProperties, error) {
	props := &TokenProperties{}
	err := loader.Bind(props)
	return props, err
}
