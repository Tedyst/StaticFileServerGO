package config

import (
	"log"

	"github.com/namsral/flag"
)

var (
	Addr      = flag.String("addr", "localhost:8080", "TCP address to listen to")
	ByteRange = flag.Bool("byteRange", true, "Enables byte range requests if set to true")
	Compress  = flag.Bool("compress", true, "Enables transparent response compression if set to true")
	Dir       = flag.String("dir", "./serve", "Directory to serve static files from")
	APIHost   = flag.String("apihost", "localhost:8080", "API host from which to accept uploads and changes")
	KeyFile   = flag.String("keyFile", "keyfile.txt", "Where are the API keys stored")

	// Debug mode disables Yubico OTP checking
	Debug = flag.Bool("debug", false, "Enable Debugging Mode")

	ClientID       = flag.String("YubicoClientID", "57189", "Yubico Client ID")
	ClientSecret   = flag.String("YubicoClientSecret", "", "Yubico Client Secret")
	AllowedYubikey = flag.String("YubikeyOtpID", "", "Yubikey OTP ID")
)

func Init() {
	flag.Parse()
	log.Print(*AllowedYubikey)
}
