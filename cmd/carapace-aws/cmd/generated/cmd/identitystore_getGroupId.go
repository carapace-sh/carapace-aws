package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_getGroupIdCmd = &cobra.Command{
	Use:   "get-group-id",
	Short: "Retrieves `GroupId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_getGroupIdCmd).Standalone()

	identitystore_getGroupIdCmd.Flags().String("alternate-identifier", "", "A unique identifier for a user or group that is not the primary identifier.")
	identitystore_getGroupIdCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_getGroupIdCmd.MarkFlagRequired("alternate-identifier")
	identitystore_getGroupIdCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_getGroupIdCmd)
}
