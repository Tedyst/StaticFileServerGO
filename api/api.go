package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tedyst/staticfileservergo/config"

	"github.com/tedyst/staticfileservergo/auth"

	"github.com/google/uuid"

	"github.com/valyala/fasthttp"
)

var apiKeys = make(map[string]string)

func APIHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/keys/create":
		createAPIKeys(ctx)
	case "/keys/delete":
		createAPIKeys(ctx)
	default:
		invalidMethod(ctx)
	}
}

type createAPIStruct struct {
	OTP  string `json:"otp"`
	Path string `json:"path"`
}

type pathAPIResponse struct {
	Path string `json:"path"`
	Key  string `json:"key"`
}

func createAPIKeys(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		invalidRequest(ctx)
		return
	}
	var data createAPIStruct
	err := json.Unmarshal(ctx.PostBody(), &data)
	if err != nil {
		invalidRequest(ctx)
		return
	}
	if !auth.YubikeyVerify(data.OTP) {
		notAllowed(ctx)
		return
	}
	response := &pathAPIResponse{
		Path: data.Path,
	}

	if _, ok := apiKeys[data.Path]; ok {
		response.Key = apiKeys[data.Path]
	} else {
		key := uuid.New()
		apiKeys[data.Path] = key.String()
		response.Key = key.String()
		appendToFile(key.String(), data.Path)
		log.Printf("Added path %q with api key %q", data.Path, key.String())
	}

	jsonResponse, _ := json.Marshal(response)
	ctx.Write(jsonResponse)
}

func appendToFile(key string, path string) {
	f, err := os.OpenFile(*config.KeyFile,
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s %s\n", key, path)); err != nil {
		log.Panic(err)
	}
}

func deleteAPIKeys(ctx *fasthttp.RequestCtx) {

}

func InitAPIKeys() {
	f, err := os.Open(*config.KeyFile)
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
			apiKeys[split[1]] = split[0]
			log.Printf("Init path %q with api key %q", split[1], split[0])
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func invalidMethod(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(405)
	ctx.Write([]byte("405 Method Not Allowed"))
}

func invalidRequest(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(401)
	ctx.Write([]byte("401 "))
}

func notAllowed(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(401)
	ctx.Write([]byte("Not Allowed"))
}
