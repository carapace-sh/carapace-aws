package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getPullRequestCmd = &cobra.Command{
	Use:   "get-pull-request",
	Short: "Gets information about a pull request in a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getPullRequestCmd).Standalone()

	codecommit_getPullRequestCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
	codecommit_getPullRequestCmd.MarkFlagRequired("pull-request-id")
	codecommitCmd.AddCommand(codecommit_getPullRequestCmd)
}
