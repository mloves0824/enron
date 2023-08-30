package change

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CmdChange is enron change log tool
var CmdChange = &cobra.Command{
	Use:   "changelog",
	Short: "Get a enron change log",
	Long:  "Get a enron release or commits info. Example: enron changelog dev or enron changelog {version}",
	Run:   run,
}

var (
	token   string
	repoURL string
)

func init() {
	if repoURL = os.Getenv("ENRON_REPO"); repoURL == "" {
		repoURL = "https://github.com/mloves0824/enron.git"
	}
	CmdChange.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "github repo")
	token = os.Getenv("GITHUB_TOKEN")
}

func run(_ *cobra.Command, args []string) {
	owner, repo := ParseGithubURL(repoURL)
	api := GithubAPI{Owner: owner, Repo: repo, Token: token}
	version := "latest"
	if len(args) > 0 {
		version = args[0]
	}
	if version == "dev" {
		info := api.GetCommitsInfo()
		fmt.Print(ParseCommitsInfo(info))
		return
	}
	info := api.GetReleaseInfo(version)
	fmt.Print(ParseReleaseInfo(info))
}
