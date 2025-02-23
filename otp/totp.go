package otp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateTotp(secret string) string {
	otp, codeErr := totp.GenerateCode(secret, time.Now())

	if codeErr != nil {
		panic(codeErr)
	}
	fmt.Println("- T-OTP Code Generated ✅")
	return otp
}
