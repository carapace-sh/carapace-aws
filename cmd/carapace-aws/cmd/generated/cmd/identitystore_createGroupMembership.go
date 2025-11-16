package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_createGroupMembershipCmd = &cobra.Command{
	Use:   "create-group-membership",
	Short: "Creates a relationship between a member and a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_createGroupMembershipCmd).Standalone()

	identitystore_createGroupMembershipCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_createGroupMembershipCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_createGroupMembershipCmd.Flags().String("member-id", "", "An object that contains the identifier of a group member.")
	identitystore_createGroupMembershipCmd.MarkFlagRequired("group-id")
	identitystore_createGroupMembershipCmd.MarkFlagRequired("identity-store-id")
	identitystore_createGroupMembershipCmd.MarkFlagRequired("member-id")
	identitystoreCmd.AddCommand(identitystore_createGroupMembershipCmd)
}
