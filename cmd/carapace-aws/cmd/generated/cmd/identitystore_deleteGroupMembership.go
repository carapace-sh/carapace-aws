package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_deleteGroupMembershipCmd = &cobra.Command{
	Use:   "delete-group-membership",
	Short: "Delete a membership within a group given `MembershipId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_deleteGroupMembershipCmd).Standalone()

	identitystore_deleteGroupMembershipCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_deleteGroupMembershipCmd.Flags().String("membership-id", "", "The identifier for a `GroupMembership` in an identity store.")
	identitystore_deleteGroupMembershipCmd.MarkFlagRequired("identity-store-id")
	identitystore_deleteGroupMembershipCmd.MarkFlagRequired("membership-id")
	identitystoreCmd.AddCommand(identitystore_deleteGroupMembershipCmd)
}
