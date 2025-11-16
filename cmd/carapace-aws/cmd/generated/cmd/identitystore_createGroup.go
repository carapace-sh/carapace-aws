package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a group within the specified identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_createGroupCmd).Standalone()

	identitystore_createGroupCmd.Flags().String("description", "", "A string containing the description of the group.")
	identitystore_createGroupCmd.Flags().String("display-name", "", "A string containing the name of the group.")
	identitystore_createGroupCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store.")
	identitystore_createGroupCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_createGroupCmd)
}
