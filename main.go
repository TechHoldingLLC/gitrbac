package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

var (
	org   = flag.String("org", "TechHoldingLLC", "")
	token = flag.String("token", "", "")
)

type ctxArgs string

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func OrgCheck(ctx context.Context, client *github.Client) {
	teams, _, err := client.Teams.ListTeams(ctx, *org, &github.ListOptions{})
	CheckErr(err)

	teamsJSON, err := json.Marshal(teams)
	CheckErr(err)

	fmt.Println(string(teamsJSON))
}

func OrgHandler(ctx context.Context, client *github.Client) {
	args := flag.Args()
	args = args[1:]

	if len(args) > 0 && args[0] == "check" {
		OrgCheck(ctx, client)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxArgs("args"), args)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// organization, _, err := client.Organizations.Get(ctx, *org)
	// CheckErr(err)

	if len(args) > 0 && args[0] == "org" {
		OrgHandler(ctx, client)
	}
}
