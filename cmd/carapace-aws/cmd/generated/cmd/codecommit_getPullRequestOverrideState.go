package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getPullRequestOverrideStateCmd = &cobra.Command{
	Use:   "get-pull-request-override-state",
	Short: "Returns information about whether approval rules have been set aside (overridden) for a pull request, and if so, the Amazon Resource Name (ARN) of the user or identity that overrode the rules and their requirements for the pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getPullRequestOverrideStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getPullRequestOverrideStateCmd).Standalone()

		codecommit_getPullRequestOverrideStateCmd.Flags().String("pull-request-id", "", "The ID of the pull request for which you want to get information about whether approval rules have been set aside (overridden).")
		codecommit_getPullRequestOverrideStateCmd.Flags().String("revision-id", "", "The system-generated ID of the revision for the pull request.")
		codecommit_getPullRequestOverrideStateCmd.MarkFlagRequired("pull-request-id")
		codecommit_getPullRequestOverrideStateCmd.MarkFlagRequired("revision-id")
	})
	codecommitCmd.AddCommand(codecommit_getPullRequestOverrideStateCmd)
}
