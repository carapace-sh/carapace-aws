package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_getMembersCmd = &cobra.Command{
	Use:   "get-members",
	Short: "Returns the membership details for specified member accounts for a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_getMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_getMembersCmd).Standalone()

		detective_getMembersCmd.Flags().String("account-ids", "", "The list of Amazon Web Services account identifiers for the member account for which to return member details.")
		detective_getMembersCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph for which to request the member details.")
		detective_getMembersCmd.MarkFlagRequired("account-ids")
		detective_getMembersCmd.MarkFlagRequired("graph-arn")
	})
	detectiveCmd.AddCommand(detective_getMembersCmd)
}
