package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_describeGroupCmd = &cobra.Command{
	Use:   "describe-group",
	Short: "Retrieves the group metadata and attributes from `GroupId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_describeGroupCmd).Standalone()

	identitystore_describeGroupCmd.Flags().String("group-id", "", "The identifier for a group in the identity store.")
	identitystore_describeGroupCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store, such as `d-1234567890`.")
	identitystore_describeGroupCmd.MarkFlagRequired("group-id")
	identitystore_describeGroupCmd.MarkFlagRequired("identity-store-id")
	identitystoreCmd.AddCommand(identitystore_describeGroupCmd)
}
