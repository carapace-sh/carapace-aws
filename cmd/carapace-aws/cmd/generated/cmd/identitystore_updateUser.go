package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates the specified user metadata and attributes in the specified identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_updateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystore_updateUserCmd).Standalone()

		identitystore_updateUserCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
		identitystore_updateUserCmd.Flags().String("operations", "", "A list of `AttributeOperation` objects to apply to the requested user.")
		identitystore_updateUserCmd.Flags().String("user-id", "", "The identifier for a user in the identity store.")
		identitystore_updateUserCmd.MarkFlagRequired("identity-store-id")
		identitystore_updateUserCmd.MarkFlagRequired("operations")
		identitystore_updateUserCmd.MarkFlagRequired("user-id")
	})
	identitystoreCmd.AddCommand(identitystore_updateUserCmd)
}
