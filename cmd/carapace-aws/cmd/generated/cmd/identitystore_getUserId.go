package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_getUserIdCmd = &cobra.Command{
	Use:   "get-user-id",
	Short: "Retrieves the `UserId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_getUserIdCmd).Standalone()

	identitystore_getUserIdCmd.Flags().String("alternate-identifier", "", "A unique identifier for a user or group that is not the primary identifier.")
	identitystore_getUserIdCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_getUserIdCmd.MarkFlagRequired("alternate-identifier")
	identitystore_getUserIdCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_getUserIdCmd)
}
