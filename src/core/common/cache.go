package common

import "fmt"

var (
	CachePrefixUserInfoEmail = "user::email::%s"
	CachePrefixOTPRegister   = "otp::register::%s"
	CachePrefixGroupRoleCode = "group_role::code::%s"
)

func BuildCacheKeyGetUserInfoByEmail(email string) string {
	return fmt.Sprintf(CachePrefixUserInfoEmail, email)
}

func BuildCacheKeyOTPRegister(email string) string {
	return fmt.Sprintf(CachePrefixOTPRegister, email)
}
func BuildCacheKeyGroupRoleCode(code string) string {
	return fmt.Sprintf(CachePrefixGroupRoleCode, code)
}
