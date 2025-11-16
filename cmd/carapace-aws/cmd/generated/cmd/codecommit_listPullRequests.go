package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listPullRequestsCmd = &cobra.Command{
	Use:   "list-pull-requests",
	Short: "Returns a list of pull requests for a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listPullRequestsCmd).Standalone()

	codecommit_listPullRequestsCmd.Flags().String("author-arn", "", "Optional.")
	codecommit_listPullRequestsCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_listPullRequestsCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_listPullRequestsCmd.Flags().String("pull-request-status", "", "Optional.")
	codecommit_listPullRequestsCmd.Flags().String("repository-name", "", "The name of the repository for which you want to list pull requests.")
	codecommit_listPullRequestsCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_listPullRequestsCmd)
}
