package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_associateMemberToFleetCmd = &cobra.Command{
	Use:   "associate-member-to-fleet",
	Short: "Assigns a fleet membership level to a member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_associateMemberToFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_associateMemberToFleetCmd).Standalone()

		deadline_associateMemberToFleetCmd.Flags().String("farm-id", "", "The farm ID of the fleet to associate with the member.")
		deadline_associateMemberToFleetCmd.Flags().String("fleet-id", "", "The ID of the fleet to associate with a member.")
		deadline_associateMemberToFleetCmd.Flags().String("identity-store-id", "", "The member's identity store ID to associate with the fleet.")
		deadline_associateMemberToFleetCmd.Flags().String("membership-level", "", "The principal's membership level for the associated fleet.")
		deadline_associateMemberToFleetCmd.Flags().String("principal-id", "", "The member's principal ID to associate with a fleet.")
		deadline_associateMemberToFleetCmd.Flags().String("principal-type", "", "The member's principal type to associate with the fleet.")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("farm-id")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("fleet-id")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("identity-store-id")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("membership-level")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("principal-id")
		deadline_associateMemberToFleetCmd.MarkFlagRequired("principal-type")
	})
	deadlineCmd.AddCommand(deadline_associateMemberToFleetCmd)
}
