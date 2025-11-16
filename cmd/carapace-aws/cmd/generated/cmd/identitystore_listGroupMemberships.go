package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_listGroupMembershipsCmd = &cobra.Command{
	Use:   "list-group-memberships",
	Short: "For the specified group in the specified identity store, returns the list of all `GroupMembership` objects and returns results in paginated form.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_listGroupMembershipsCmd).Standalone()

	identitystore_listGroupMembershipsCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_listGroupMembershipsCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_listGroupMembershipsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	identitystore_listGroupMembershipsCmd.Flags().String("next-token", "", "The pagination token used for the `ListUsers`, `ListGroups` and `ListGroupMemberships` API operations.")
	identitystore_listGroupMembershipsCmd.MarkFlagRequired("group-id")
	identitystore_listGroupMembershipsCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_listGroupMembershipsCmd)
}
