package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mloves0824/enron/cmd/enron/internal/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the enron tools",
	Long:  "Upgrade the enron tools. Example: enron upgrade",
	Run:   Run,
}

// Run upgrade the enron tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"github.com/mloves0824/enron/cmd/enron@latest",
		"github.com/mloves0824/enron/cmd/protoc-gen-go-http@latest",
		"github.com/mloves0824/enron/cmd/protoc-gen-go-errors@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
