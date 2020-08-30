package main

import (
	"github.com/GeertJohan/yubigo"
)

func yubikeyVerify(otp string) bool {
	if *debug == true {
		return true
	}
	yubiAuth, err := yubigo.NewYubiAuth(*clientID, *clientSecret)
	if err != nil {
		return false
	}

	// verify an OTP string
	_, ok, err := yubiAuth.Verify(otp)
	if err != nil {
		return false
	}

	return ok
}

func apiVerify(otp string) bool {
	return true
}
