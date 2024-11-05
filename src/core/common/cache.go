package common

import "fmt"

var (
	CachePrefixUserInfoEmail = "user::email::%s"
	CachePrefixOTPRegister   = "otp::register::%s"
)

func BuildCacheKeyGetUserInfoByEmail(email string) string {
	return fmt.Sprintf(CachePrefixUserInfoEmail, email)
}

func BuildCacheKeyOTPRegister(email string) string {
	return fmt.Sprintf(CachePrefixOTPRegister, email)
}
