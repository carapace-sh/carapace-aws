package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateWebAppCmd = &cobra.Command{
	Use:   "update-web-app",
	Short: "Assigns new properties to a web app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateWebAppCmd).Standalone()

	transfer_updateWebAppCmd.Flags().String("access-endpoint", "", "The `AccessEndpoint` is the URL that you provide to your users for them to interact with the Transfer Family web app.")
	transfer_updateWebAppCmd.Flags().String("identity-provider-details", "", "Provide updated identity provider values in a `WebAppIdentityProviderDetails` object.")
	transfer_updateWebAppCmd.Flags().String("web-app-id", "", "Provide the identifier of the web app that you are updating.")
	transfer_updateWebAppCmd.Flags().String("web-app-units", "", "A union that contains the value for number of concurrent connections or the user sessions on your web app.")
	transfer_updateWebAppCmd.MarkFlagRequired("web-app-id")
	transferCmd.AddCommand(transfer_updateWebAppCmd)
}
