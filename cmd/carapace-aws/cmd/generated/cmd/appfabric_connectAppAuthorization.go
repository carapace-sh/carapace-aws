package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_connectAppAuthorizationCmd = &cobra.Command{
	Use:   "connect-app-authorization",
	Short: "Establishes a connection between Amazon Web Services AppFabric and an application, which allows AppFabric to call the APIs of the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_connectAppAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_connectAppAuthorizationCmd).Standalone()

		appfabric_connectAppAuthorizationCmd.Flags().String("app-authorization-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app authorization to use for the request.")
		appfabric_connectAppAuthorizationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle that contains the app authorization to use for the request.")
		appfabric_connectAppAuthorizationCmd.Flags().String("auth-request", "", "Contains OAuth2 authorization information.")
		appfabric_connectAppAuthorizationCmd.MarkFlagRequired("app-authorization-identifier")
		appfabric_connectAppAuthorizationCmd.MarkFlagRequired("app-bundle-identifier")
	})
	appfabricCmd.AddCommand(appfabric_connectAppAuthorizationCmd)
}
