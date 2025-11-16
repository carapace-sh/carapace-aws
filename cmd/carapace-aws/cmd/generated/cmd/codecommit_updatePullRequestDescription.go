package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updatePullRequestDescriptionCmd = &cobra.Command{
	Use:   "update-pull-request-description",
	Short: "Replaces the contents of the description of a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updatePullRequestDescriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updatePullRequestDescriptionCmd).Standalone()

		codecommit_updatePullRequestDescriptionCmd.Flags().String("description", "", "The updated content of the description for the pull request.")
		codecommit_updatePullRequestDescriptionCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_updatePullRequestDescriptionCmd.MarkFlagRequired("description")
		codecommit_updatePullRequestDescriptionCmd.MarkFlagRequired("pull-request-id")
	})
	codecommitCmd.AddCommand(codecommit_updatePullRequestDescriptionCmd)
}
