package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_listGroupMembershipsForMemberCmd = &cobra.Command{
	Use:   "list-group-memberships-for-member",
	Short: "For the specified member in the specified identity store, returns the list of all `GroupMembership` objects and returns results in paginated form.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_listGroupMembershipsForMemberCmd).Standalone()

	identitystore_listGroupMembershipsForMemberCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_listGroupMembershipsForMemberCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	identitystore_listGroupMembershipsForMemberCmd.Flags().String("member-id", "", "An object that contains the identifier of a group member.")
	identitystore_listGroupMembershipsForMemberCmd.Flags().String("next-token", "", "The pagination token used for the `ListUsers`, `ListGroups`, and `ListGroupMemberships` API operations.")
	identitystore_listGroupMembershipsForMemberCmd.MarkFlagRequired("identity-store-id")
	identitystore_listGroupMembershipsForMemberCmd.MarkFlagRequired("member-id")
	identitystoreCmd.AddCommand(identitystore_listGroupMembershipsForMemberCmd)
}
