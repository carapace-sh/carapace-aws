package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Lists all users in the identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_listUsersCmd).Standalone()

	identitystore_listUsersCmd.Flags().String("filters", "", "A list of `Filter` objects, which is used in the `ListUsers` and `ListGroups` requests.")
	identitystore_listUsersCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store, such as `d-1234567890`.")
	identitystore_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	identitystore_listUsersCmd.Flags().String("next-token", "", "The pagination token used for the `ListUsers` and `ListGroups` API operations.")
	identitystore_listUsersCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_listUsersCmd)
}
