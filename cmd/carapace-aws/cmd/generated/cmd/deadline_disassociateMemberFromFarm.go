package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_disassociateMemberFromFarmCmd = &cobra.Command{
	Use:   "disassociate-member-from-farm",
	Short: "Disassociates a member from a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_disassociateMemberFromFarmCmd).Standalone()

	deadline_disassociateMemberFromFarmCmd.Flags().String("farm-id", "", "The farm ID of the farm to disassociate from the member.")
	deadline_disassociateMemberFromFarmCmd.Flags().String("principal-id", "", "A member's principal ID to disassociate from a farm.")
	deadline_disassociateMemberFromFarmCmd.MarkFlagRequired("farm-id")
	deadline_disassociateMemberFromFarmCmd.MarkFlagRequired("principal-id")
	deadlineCmd.AddCommand(deadline_disassociateMemberFromFarmCmd)
}
