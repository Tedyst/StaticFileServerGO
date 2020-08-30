package main

import (
	"bytes"
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr      = flag.String("addr", "localhost:8080", "TCP address to listen to")
	byteRange = flag.Bool("byteRange", true, "Enables byte range requests if set to true")
	compress  = flag.Bool("compress", true, "Enables transparent response compression if set to true")
	dir       = flag.String("dir", "./serve", "Directory to serve static files from")
	apihost   = flag.String("apihost", "localhost:8080", "API host from which to accept uploads and changes")
	keyfile   = flag.String("keyFile", "keyfile.txt", "Where are the API keys stored")
)

func main() {
	// Parse command-line flags.
	flag.Parse()
	initAPIKeys()

	// Setup FS handler
	fs := &fasthttp.FS{
		Root:            *dir,
		IndexNames:      []string{"index.html"},
		Compress:        *compress,
		AcceptByteRange: *byteRange,
		PathRewrite:     fasthttp.NewVHostPathRewriter(0),
	}

	fsHandler := fs.NewRequestHandler()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		if bytes.Compare(ctx.Host(), []byte(*apihost)) == 0 {
			apiHandler(ctx)
		} else {
			log.Printf("Loaded %q", ctx.URI())
			fsHandler(ctx)
		}
	}

	// Start HTTP server.
	if len(*addr) > 0 {
		log.Printf("Starting HTTP server on %q", *addr)
		go func() {
			if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}

	log.Printf("Serving files from directory %q", *dir)

	// Wait forever.
	select {}
}
