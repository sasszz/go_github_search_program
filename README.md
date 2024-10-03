# GitHub Repository Search Program

This program allows you to search for GitHub repositories using a keyword. It uses the GitHub API and displays repository details such as the name, owner, stars, and description. The program is written in Go and requires a GitHub personal access token to authenticate API requests.

<hr>

![fasterdemo](https://github.com/user-attachments/assets/4f48cefd-68a0-4d97-8186-c92bdc01d49d)

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
3. Opens a Text Based User Interface (TUI) in the terminal with a search box and button.
4. Searches GitHub repositories based on the search term provided in the TUI.
5. Displays the repository name, owner, stars, and description for each result.

<hr>

### Getting Started

Install dependencies
 ```
go get github.com/joho/godotenv
go get github.com/gdamore/tcell/v2
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
go run main.go github_client.go
```

<hr>

Source Inc 🚀 Lucie Chevreuil 👩‍💻 Wednesday, October 2nd 🎃


