package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user within an identity store given `UserId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_deleteUserCmd).Standalone()

	identitystore_deleteUserCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_deleteUserCmd.Flags().String("user-id", "", "The identifier for a user in the identity store.")
	identitystore_deleteUserCmd.MarkFlagRequired("identity-store-id")
	identitystore_deleteUserCmd.MarkFlagRequired("user-id")
	identitystoreCmd.AddCommand(identitystore_deleteUserCmd)
}
