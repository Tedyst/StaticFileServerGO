package auth

import (
	"strings"

	"github.com/GeertJohan/yubigo"
	"github.com/tedyst/staticfileservergo/config"
)

// YubikeyVerify verifies if an OTP is correct
func YubikeyVerify(otp string) bool {
	if *config.Debug == true {
		return true
	}
	if len(*config.AllowedYubikey) == 0 {
		return false
	}
	yubiAuth, err := yubigo.NewYubiAuth(*config.ClientID, *config.ClientSecret)
	if err != nil {
		return false
	}

	// verify an OTP string
	_, ok, err := yubiAuth.Verify(otp)
	if err != nil {
		return false
	}
	if !ok {
		return false
	}
	if !strings.HasPrefix(otp, *config.AllowedYubikey) {
		return false
	}
	return true
}
