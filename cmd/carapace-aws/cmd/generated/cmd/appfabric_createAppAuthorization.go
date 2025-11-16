package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_createAppAuthorizationCmd = &cobra.Command{
	Use:   "create-app-authorization",
	Short: "Creates an app authorization within an app bundle, which allows AppFabric to connect to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_createAppAuthorizationCmd).Standalone()

	appfabric_createAppAuthorizationCmd.Flags().String("app", "", "The name of the application.")
	appfabric_createAppAuthorizationCmd.Flags().String("app-bundle-identifier", "", "The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.")
	appfabric_createAppAuthorizationCmd.Flags().String("auth-type", "", "The authorization type for the app authorization.")
	appfabric_createAppAuthorizationCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	appfabric_createAppAuthorizationCmd.Flags().String("credential", "", "Contains credentials for the application, such as an API key or OAuth2 client ID and secret.")
	appfabric_createAppAuthorizationCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
	appfabric_createAppAuthorizationCmd.Flags().String("tenant", "", "Contains information about an application tenant, such as the application display name and identifier.")
	appfabric_createAppAuthorizationCmd.MarkFlagRequired("app")
	appfabric_createAppAuthorizationCmd.MarkFlagRequired("app-bundle-identifier")
	appfabric_createAppAuthorizationCmd.MarkFlagRequired("auth-type")
	appfabric_createAppAuthorizationCmd.MarkFlagRequired("credential")
	appfabric_createAppAuthorizationCmd.MarkFlagRequired("tenant")
	appfabricCmd.AddCommand(appfabric_createAppAuthorizationCmd)
}
