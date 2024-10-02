# GitHub Repository Search Program

This program allows you to search for GitHub repositories using a keyword. It uses the GitHub API and displays repository details such as the name, owner, stars, and description. The program is written in Go and requires a GitHub personal access token to authenticate API requests.

<hr>

### NOTE
THis program is in beta. It currently only displays a Read Only printout og the Github API search results in the terminal. Next step would be to create an interactive text based elements to create a UI on the command line.

<hr>

### Technologies
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

<hr>

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- A GitHub personal access token with the necessary scopes (e.g., `public_repo` scope)

<hr>

### Program Overview

The program is a simple command-line tool that does the following:
1. Loads environment variables from a `.env` file using the `godotenv` package.
2. Authenticates with the GitHub API using a personal access token.
3. Searches GitHub repositories based on a query provided as a command-line argument.
4. Displays the repository name, owner, stars, and description for each result.

<hr>

### Getting Started

Install dependencies
 ```
go get github.com/joho/godotenv
go get github.com/google/go-github/v65/github
go get golang.org/x/oauth2
```

Create .env file with GH Access Token
```
touch .env
```
```
GITHUB_TOKEN=your_actual_token_here
```

Run the program in the terminal
```
go run main.go <search_term>
```

<hr>

### Example Output

```bash
$ go run main.go kubernetes

Loaded token from .env: your_actual_token_here
Found 1000 repositories:
Name: kubernetes
Owner: kubernetes
Stars: 84000
Description: Production-Grade Container Scheduling and Management

Name: minikube
Owner: kubernetes
Stars: 21000
Description: Run Kubernetes locally
...
```

<hr>

Source Inc 🚀 Lucie Chevreuil 👩‍💻 Wednesday, October 2nd 🎃


