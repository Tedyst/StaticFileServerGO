package main

import (
	"flag"
	"log"
	"os"

	"github.com/GeertJohan/yubigo"
)

var (
	clientID     = flag.String("YubicoClientID", "57189", "Yubico Client ID")
	clientSecret = flag.String("YubicoClientSecret", os.Getenv("YUBICO_KEY"), "Yubico Client Secret")
)

func yubikeyVerify(otp string) bool {
	yubiAuth, err := yubigo.NewYubiAuth(*clientID, *clientSecret)
	if err != nil {
		log.Fatalln(err)
	}

	// verify an OTP string
	_, ok, err := yubiAuth.Verify(otp)
	if err != nil {
		log.Fatalln(err)
	}

	return ok
}

func apiVerify(otp string) bool {
	return true
}
