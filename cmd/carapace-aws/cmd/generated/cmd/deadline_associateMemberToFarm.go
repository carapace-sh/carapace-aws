package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_associateMemberToFarmCmd = &cobra.Command{
	Use:   "associate-member-to-farm",
	Short: "Assigns a farm membership level to a member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_associateMemberToFarmCmd).Standalone()

	deadline_associateMemberToFarmCmd.Flags().String("farm-id", "", "The ID of the farm to associate with the member.")
	deadline_associateMemberToFarmCmd.Flags().String("identity-store-id", "", "The identity store ID of the member to associate with the farm.")
	deadline_associateMemberToFarmCmd.Flags().String("membership-level", "", "The principal's membership level for the associated farm.")
	deadline_associateMemberToFarmCmd.Flags().String("principal-id", "", "The member's principal ID to associate with the farm.")
	deadline_associateMemberToFarmCmd.Flags().String("principal-type", "", "The principal type of the member to associate with the farm.")
	deadline_associateMemberToFarmCmd.MarkFlagRequired("farm-id")
	deadline_associateMemberToFarmCmd.MarkFlagRequired("identity-store-id")
	deadline_associateMemberToFarmCmd.MarkFlagRequired("membership-level")
	deadline_associateMemberToFarmCmd.MarkFlagRequired("principal-id")
	deadline_associateMemberToFarmCmd.MarkFlagRequired("principal-type")
	deadlineCmd.AddCommand(deadline_associateMemberToFarmCmd)
}
