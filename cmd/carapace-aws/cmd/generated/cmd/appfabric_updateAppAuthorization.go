package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_updateAppAuthorizationCmd = &cobra.Command{
	Use:   "update-app-authorization",
	Short: "Updates an app authorization within an app bundle, which allows AppFabric to connect to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_updateAppAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appfabric_updateAppAuthorizationCmd).Standalone()

		appfabric_updateAppAuthorizationCmd.Flags().String("app-authorization-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app authorization to use for the request.")
		appfabric_updateAppAuthorizationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
		appfabric_updateAppAuthorizationCmd.Flags().String("credential", "", "Contains credentials for the application, such as an API key or OAuth2 client ID and secret.")
		appfabric_updateAppAuthorizationCmd.Flags().String("tenant", "", "Contains information about an application tenant, such as the application display name and identifier.")
		appfabric_updateAppAuthorizationCmd.MarkFlagRequired("app-authorization-identifier")
		appfabric_updateAppAuthorizationCmd.MarkFlagRequired("app-bundle-identifier")
	})
	appfabricCmd.AddCommand(appfabric_updateAppAuthorizationCmd)
}
