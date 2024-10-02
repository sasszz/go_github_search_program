package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v65/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// getGitHubClient initializes and returns a GitHub client using an OAuth token loaded from the .env file.
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

// searchRepositories searches for GitHub repositories based on the provided query, sorted by stars, and prints the results.
func searchRepositories(client *github.Client, query string) {
    ctx := context.Background()
    opts := &github.SearchOptions{
        Sort:  "stars",
        Order: "desc",
    }

    result, _, err := client.Search.Repositories(ctx, query, opts)
    if err != nil {
        log.Fatalf("Error searching repositories: %v", err)
    }

    fmt.Printf("Found %d repositories:\n", result.GetTotal())
    for _, repo := range result.Repositories {
        fmt.Printf("Name: %s\nOwner: %s\nStars: %d\nDescription: %s\n\n",
            repo.GetName(), repo.Owner.GetLogin(), repo.GetStargazersCount(), repo.GetDescription())
    }
}

// main function parses the search term from command-line arguments, creates a GitHub client, and calls the search function.
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: gx <search_term>")
        return
    }

    searchTerm := os.Args[1]
    client := getGitHubClient()
    searchRepositories(client, searchTerm)
}
