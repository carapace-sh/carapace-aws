package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_disassociateMembershipCmd = &cobra.Command{
	Use:   "disassociate-membership",
	Short: "Removes the member account from the specified behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_disassociateMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_disassociateMembershipCmd).Standalone()

		detective_disassociateMembershipCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph to remove the member account from.")
		detective_disassociateMembershipCmd.MarkFlagRequired("graph-arn")
	})
	detectiveCmd.AddCommand(detective_disassociateMembershipCmd)
}
