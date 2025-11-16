package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updatePullRequestTitleCmd = &cobra.Command{
	Use:   "update-pull-request-title",
	Short: "Replaces the title of a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updatePullRequestTitleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updatePullRequestTitleCmd).Standalone()

		codecommit_updatePullRequestTitleCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_updatePullRequestTitleCmd.Flags().String("title", "", "The updated title of the pull request.")
		codecommit_updatePullRequestTitleCmd.MarkFlagRequired("pull-request-id")
		codecommit_updatePullRequestTitleCmd.MarkFlagRequired("title")
	})
	codecommitCmd.AddCommand(codecommit_updatePullRequestTitleCmd)
}
