package main

import (
	"bytes"
	"log"

	"github.com/tedyst/staticfileservergo/api"
	"github.com/tedyst/staticfileservergo/config"

	"github.com/valyala/fasthttp"
)

func main() {
	// Parse command-line flags.
	config.Init()
	api.InitAPIKeys()
	// Setup FS handler
	fs := &fasthttp.FS{
		Root:            *config.Dir,
		IndexNames:      []string{"index.html"},
		Compress:        *config.Compress,
		AcceptByteRange: *config.ByteRange,
		PathRewrite:     fasthttp.NewVHostPathRewriter(0),
	}

	fsHandler := fs.NewRequestHandler()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		if bytes.Compare(ctx.Host(), []byte(*config.APIHost)) == 0 {
			api.APIHandler(ctx)
		} else {
			log.Printf("Loaded %q", ctx.URI())
			fsHandler(ctx)
		}
	}

	// Start HTTP server.
	if len(*config.Addr) > 0 {
		log.Printf("Starting HTTP server on %q", *config.Addr)
		go func() {
			if err := fasthttp.ListenAndServe(*config.Addr, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}

	log.Printf("Serving files from directory %q", *config.Dir)

	// Wait forever.
	select {}
}
