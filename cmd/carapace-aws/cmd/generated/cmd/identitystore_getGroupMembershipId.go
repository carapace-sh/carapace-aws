package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_getGroupMembershipIdCmd = &cobra.Command{
	Use:   "get-group-membership-id",
	Short: "Retrieves the `MembershipId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_getGroupMembershipIdCmd).Standalone()

	identitystore_getGroupMembershipIdCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_getGroupMembershipIdCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_getGroupMembershipIdCmd.Flags().String("member-id", "", "An object that contains the identifier of a group member.")
	identitystore_getGroupMembershipIdCmd.MarkFlagRequired("group-id")
	identitystore_getGroupMembershipIdCmd.MarkFlagRequired("identity-store-id")
	identitystore_getGroupMembershipIdCmd.MarkFlagRequired("member-id")
	identitystoreCmd.AddCommand(identitystore_getGroupMembershipIdCmd)
}
