package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createWebAppCmd = &cobra.Command{
	Use:   "create-web-app",
	Short: "Creates a web app based on specified parameters, and returns the ID for the new web app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createWebAppCmd).Standalone()

	transfer_createWebAppCmd.Flags().String("access-endpoint", "", "The `AccessEndpoint` is the URL that you provide to your users for them to interact with the Transfer Family web app.")
	transfer_createWebAppCmd.Flags().String("identity-provider-details", "", "You can provide a structure that contains the details for the identity provider to use with your web app.")
	transfer_createWebAppCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for web apps.")
	transfer_createWebAppCmd.Flags().String("web-app-endpoint-policy", "", "Setting for the type of endpoint policy for the web app.")
	transfer_createWebAppCmd.Flags().String("web-app-units", "", "A union that contains the value for number of concurrent connections or the user sessions on your web app.")
	transfer_createWebAppCmd.MarkFlagRequired("identity-provider-details")
	transferCmd.AddCommand(transfer_createWebAppCmd)
}
