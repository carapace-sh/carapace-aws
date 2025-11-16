package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identitystore_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Retrieves the user metadata and attributes from the `UserId` in an identity store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identitystore_describeUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(identitystore_describeUserCmd).Standalone()

		identitystore_describeUserCmd.Flags().String("identity-store-id", "", "The globally unique identifier for the identity store, such as `d-1234567890`.")
		identitystore_describeUserCmd.Flags().String("user-id", "", "The identifier for a user in the identity store.")
		identitystore_describeUserCmd.MarkFlagRequired("identity-store-id")
		identitystore_describeUserCmd.MarkFlagRequired("user-id")
	})
	identitystoreCmd.AddCommand(identitystore_describeUserCmd)
}
