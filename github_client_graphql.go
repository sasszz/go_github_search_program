package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	graphql "github.com/hasura/go-graphql-client"
)

func GetGraphQLGithub() *graphql.Client {
	// get dotenv GH token
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// := declare + assigns a variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Error: GITHUB_TOKEN is not set in the environment")
	}

	// create auth client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// create graphql client
	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// we want the graphql client returned
	return client
}

func ExecuteGraphQL() error {

	client := GetGraphQLGithub();

	query:= `query {
		search(query:"hello", type:REPOSITORY first:10) {
			edges {
				node {
					... on Repository {
						name
						owner {
							login
						}
						primaryLanguage {
							name
						}
						description
						stargazerCount
					}
				}
			}
		}
	}`
	
	var output struct {
		Search struct {
			Edges []struct {
				Node struct {
					Repository struct {
						Name string
						Owner struct {
							Login string
						}
						PrimaryLanguage struct {
							Name string
						}
						Description string
						StargazerCount int
					} `graphql:"... on Repository"`
				}
			}
		} 
	}

	err := client.Exec(context.Background(), query, &output, nil,)
	if err != nil {
		panic(err)
	}

	fmt.Println("above")
	print(output)
	fmt.Println("below")

	return nil
}

// print pretty prints v to stdout. It panics on any error.
func print(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}