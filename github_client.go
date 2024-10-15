package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/google/go-github/v65/github"
	"golang.org/x/oauth2"
)

// GetGitHubClient initializes and returns a GitHub client using an OAuth token loaded from the .env file.
func getGitHubClient() *github.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Error: GITHUB_TOKEN is not set in the environment")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

// SearchRepositories searches for GitHub repositories based on the provided query.
func searchRepositories(client *github.Client, query string) ([]*github.Repository, error) {
	ctx := context.Background()
	opts := &github.SearchOptions{
		Sort:  "stars",
		Order: "desc",
	}

	result, _, err := client.Search.Repositories(ctx, query, opts)
	if err != nil {
		return nil, err
	}

	return result.Repositories, nil
}


