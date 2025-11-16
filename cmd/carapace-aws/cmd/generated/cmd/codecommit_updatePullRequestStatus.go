package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updatePullRequestStatusCmd = &cobra.Command{
	Use:   "update-pull-request-status",
	Short: "Updates the status of a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updatePullRequestStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updatePullRequestStatusCmd).Standalone()

		codecommit_updatePullRequestStatusCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_updatePullRequestStatusCmd.Flags().String("pull-request-status", "", "The status of the pull request.")
		codecommit_updatePullRequestStatusCmd.MarkFlagRequired("pull-request-id")
		codecommit_updatePullRequestStatusCmd.MarkFlagRequired("pull-request-status")
	})
	codecommitCmd.AddCommand(codecommit_updatePullRequestStatusCmd)
}
