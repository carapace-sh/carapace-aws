package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_isMemberInGroupsCmd = &cobra.Command{
	Use:   "is-member-in-groups",
	Short: "Checks the user's membership in all requested groups and returns if the member exists in all queried groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_isMemberInGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystore_isMemberInGroupsCmd).Standalone()

		identitystore_isMemberInGroupsCmd.Flags().String("group-ids", "", "A list of identifiers for groups in the identity store.")
		identitystore_isMemberInGroupsCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
		identitystore_isMemberInGroupsCmd.Flags().String("member-id", "", "An object containing the identifier of a group member.")
		identitystore_isMemberInGroupsCmd.MarkFlagRequired("group-ids")
		identitystore_isMemberInGroupsCmd.MarkFlagRequired("identity-store-id")
		identitystore_isMemberInGroupsCmd.MarkFlagRequired("member-id")
	})
	identitystoreCmd.AddCommand(identitystore_isMemberInGroupsCmd)
}
