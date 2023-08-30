package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/mloves0824/enron/cmd/enron/internal/change"
	"github.com/mloves0824/enron/cmd/enron/internal/project"
	"github.com/mloves0824/enron/cmd/enron/internal/proto"
	"github.com/mloves0824/enron/cmd/enron/internal/run"
	"github.com/mloves0824/enron/cmd/enron/internal/upgrade"
)

var rootCmd = &cobra.Command{
	Use:     "enron",
	Short:   "Enron: An elegant toolkit for Go microservices.",
	Long:    `Enron: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
