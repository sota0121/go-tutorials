package ghpra

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

const (
	rootDir = "features/ghpra"
)

type ConfigElement struct {
	SourceOwner string `yaml:"source_owner"`
	SourceRepo  string `yaml:"source_repo"`
	BaseBranch  string `yaml:"base_branch"`
}

type Config struct {
	Elements []ConfigElement `yaml:"elements"`
}

// Main is the entry point of the feature
func Main() {
	fmt.Println(">> Start GitHub Pull Request Aggregation!")
	// Load config file
	config, err := loadConfig(rootDir + "/config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Let the user select the config
	fmt.Println("Select the config:")
	for i, element := range config.Elements {
		fmt.Println(i, ":", element.SourceOwner, "/", element.SourceRepo, ":", element.BaseBranch)
	}
	var selected int
	fmt.Scan(&selected)
	if selected < 0 || selected >= len(config.Elements) {
		fmt.Println("Invalid selection")
		return
	}
	sourceOwner := config.Elements[selected].SourceOwner
	sourceRepo := config.Elements[selected].SourceRepo
	baseBranch := config.Elements[selected].BaseBranch

	// Load .env file
	err = godotenv.Load(rootDir + "/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	ght := os.Getenv("GITHUB_ACCESS_TOKEN")

	// Authentication
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: ght,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// get my user
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Authenticated user:", *user.Login)

	// list recent pull requests for the authenticated user
	prs, _, err := listClosedPullRequestsByClient(ctx, client, sourceOwner, sourceRepo, baseBranch, *user.Login)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pull Requests of the authenticated user (", len(prs), ") :")
	for _, pr := range prs {
		fmt.Println(*pr.Title)
	}

	// // list all repositories for the authenticated user
	// repos, _, err := listRepos(ctx, client)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Repositories of the authenticated user:")
	// for _, repo := range repos {
	// 	fmt.Println(*repo.Name)
	// }
}

func loadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func listClosedPullRequestsByClient(ctx context.Context, client *github.Client, owner, repo, baseBranch, username string) ([]*github.PullRequest, *github.Response, error) {
	// Create PullRequestListOptions
	// https://godoc.org/github.com/google/go-github/github#PullRequestListOptions
	opt := &github.PullRequestListOptions{
		State: "closed",
		ListOptions: github.ListOptions{
			PerPage: 30,
		},
		Base: baseBranch,
		// Head: username + ":" + "develop",
	}

	// list all pull requests
	var prs []*github.PullRequest
	for {
		curprs, res, err := client.PullRequests.List(ctx, owner, repo, opt)
		if err != nil {
			return nil, res, err
		}
		prs = append(prs, curprs...)
		fmt.Print("*")

		if res.NextPage == 0 {
			fmt.Println("")
			break
		}
		opt.ListOptions.Page = res.NextPage
	}

	// extract pull requests which owned by the authenticated user
	var myprs []*github.PullRequest
	for _, pr := range prs {
		if *pr.User.Login == username {
			myprs = append(myprs, pr)
		}
	}

	return myprs, nil, nil
}

func listRepos(ctx context.Context, client *github.Client) ([]*github.Repository, *github.Response, error) {
	repos, res, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		return nil, nil, err
	}
	return repos, res, nil
}
