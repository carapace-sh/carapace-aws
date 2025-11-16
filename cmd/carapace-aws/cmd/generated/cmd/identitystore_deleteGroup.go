package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Delete a group within an identity store given `GroupId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_deleteGroupCmd).Standalone()

	identitystore_deleteGroupCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_deleteGroupCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_deleteGroupCmd.MarkFlagRequired("group-id")
	identitystore_deleteGroupCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_deleteGroupCmd)
}
