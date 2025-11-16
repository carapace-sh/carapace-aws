package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getFederationTokenCmd = &cobra.Command{
	Use:   "get-federation-token",
	Short: "Supports SAML sign-in for Amazon Connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getFederationTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getFederationTokenCmd).Standalone()

		connect_getFederationTokenCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getFederationTokenCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_getFederationTokenCmd)
}
