package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomOTPRegister() (string, error) {

	maxNum := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, maxNum)
	if err != nil {
		return "", err
	}

	// Format it to a 6-digit string with leading zeros if necessary
	otp := fmt.Sprintf("%06d", n.Int64())
	return otp, nil
}
