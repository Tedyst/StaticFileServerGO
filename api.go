package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var keys = make(map[string]string)

func apiHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/keys/create":
		createAPIKeys(ctx)
	case "/keys/delete":
		createAPIKeys(ctx)
	default:
		invalidRequest(ctx)
	}
}

func invalidRequest(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(405)
	ctx.Write([]byte("405 Method Not Allowed"))
}

func createAPIKeys(ctx *fasthttp.RequestCtx) {

}

func deleteAPIKeys(ctx *fasthttp.RequestCtx) {

}

func initAPIKeys() {
	f, err := os.Open(*keyfile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		split := strings.Split(s.Text(), " ")
		if len(split) == 2 {
			keys[split[0]] = split[1]
			log.Printf("Init path %q with api key %q", split[1], split[0])
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
