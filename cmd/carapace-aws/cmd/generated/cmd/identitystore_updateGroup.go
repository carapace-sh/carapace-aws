package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates the specified group metadata and attributes in the specified identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_updateGroupCmd).Standalone()

	identitystore_updateGroupCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_updateGroupCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_updateGroupCmd.Flags().String("operations", "", "A list of `AttributeOperation` objects to apply to the requested group.")
	identitystore_updateGroupCmd.MarkFlagRequired("group-id")
	identitystore_updateGroupCmd.MarkFlagRequired("identity-store-id")
	identitystore_updateGroupCmd.MarkFlagRequired("operations")
	identitystoreCmd.AddCommand(identitystore_updateGroupCmd)
}
