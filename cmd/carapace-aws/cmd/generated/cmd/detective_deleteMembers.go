package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_deleteMembersCmd = &cobra.Command{
	Use:   "delete-members",
	Short: "Removes the specified member accounts from the behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_deleteMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_deleteMembersCmd).Standalone()

		detective_deleteMembersCmd.Flags().String("account-ids", "", "The list of Amazon Web Services account identifiers for the member accounts to remove from the behavior graph.")
		detective_deleteMembersCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph to remove members from.")
		detective_deleteMembersCmd.MarkFlagRequired("account-ids")
		detective_deleteMembersCmd.MarkFlagRequired("graph-arn")
	})
	detectiveCmd.AddCommand(detective_deleteMembersCmd)
}
