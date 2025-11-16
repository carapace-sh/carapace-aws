package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_describeGroupMembershipCmd = &cobra.Command{
	Use:   "describe-group-membership",
	Short: "Retrieves membership metadata and attributes from `MembershipId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_describeGroupMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystore_describeGroupMembershipCmd).Standalone()

		identitystore_describeGroupMembershipCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
		identitystore_describeGroupMembershipCmd.Flags().String("membership-id", "", "The identifier for a `GroupMembership` in an identity store.")
		identitystore_describeGroupMembershipCmd.MarkFlagRequired("identity-store-id")
		identitystore_describeGroupMembershipCmd.MarkFlagRequired("membership-id")
	})
	identitystoreCmd.AddCommand(identitystore_describeGroupMembershipCmd)
}
