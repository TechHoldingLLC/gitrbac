package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

var (
	org   = flag.String("org", "TechHoldingLLC", "")
	token = flag.String("token", "", "")
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	
}
