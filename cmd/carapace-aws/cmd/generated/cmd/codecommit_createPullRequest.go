package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createPullRequestCmd = &cobra.Command{
	Use:   "create-pull-request",
	Short: "Creates a pull request in the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createPullRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_createPullRequestCmd).Standalone()

		codecommit_createPullRequestCmd.Flags().String("client-request-token", "", "A unique, client-generated idempotency token that, when provided in a request, ensures the request cannot be repeated with a changed parameter.")
		codecommit_createPullRequestCmd.Flags().String("description", "", "A description of the pull request.")
		codecommit_createPullRequestCmd.Flags().String("targets", "", "The targets for the pull request, including the source of the code to be reviewed (the source branch) and the destination where the creator of the pull request intends the code to be merged after the pull request is closed (the destination branch).")
		codecommit_createPullRequestCmd.Flags().String("title", "", "The title of the pull request.")
		codecommit_createPullRequestCmd.MarkFlagRequired("targets")
		codecommit_createPullRequestCmd.MarkFlagRequired("title")
	})
	codecommitCmd.AddCommand(codecommit_createPullRequestCmd)
}
