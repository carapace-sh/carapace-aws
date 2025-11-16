package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_disassociateMemberFromFleetCmd = &cobra.Command{
	Use:   "disassociate-member-from-fleet",
	Short: "Disassociates a member from a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_disassociateMemberFromFleetCmd).Standalone()

	deadline_disassociateMemberFromFleetCmd.Flags().String("farm-id", "", "The farm ID of the fleet to disassociate a member from.")
	deadline_disassociateMemberFromFleetCmd.Flags().String("fleet-id", "", "The fleet ID of the fleet to from which to disassociate a member.")
	deadline_disassociateMemberFromFleetCmd.Flags().String("principal-id", "", "A member's principal ID to disassociate from a fleet.")
	deadline_disassociateMemberFromFleetCmd.MarkFlagRequired("farm-id")
	deadline_disassociateMemberFromFleetCmd.MarkFlagRequired("fleet-id")
	deadline_disassociateMemberFromFleetCmd.MarkFlagRequired("principal-id")
	deadlineCmd.AddCommand(deadline_disassociateMemberFromFleetCmd)
}
