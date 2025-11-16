package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Lists all groups in the identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystore_listGroupsCmd).Standalone()

		identitystore_listGroupsCmd.Flags().String("filters", "", "A list of `Filter` objects, which is used in the `ListUsers` and `ListGroups` requests.")
		identitystore_listGroupsCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store, such as `d-1234567890`.")
		identitystore_listGroupsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		identitystore_listGroupsCmd.Flags().String("next-token", "", "The pagination token used for the `ListUsers` and `ListGroups` API operations.")
		identitystore_listGroupsCmd.MarkFlagRequired("identity-store-id")
	})
	identitystoreCmd.AddCommand(identitystore_listGroupsCmd)
}
