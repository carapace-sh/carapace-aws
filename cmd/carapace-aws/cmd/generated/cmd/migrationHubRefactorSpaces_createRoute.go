package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_createRouteCmd = &cobra.Command{
	Use:   "create-route",
	Short: "Creates an Amazon Web Services Migration Hub Refactor Spaces route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_createRouteCmd).Standalone()

	migrationHubRefactorSpaces_createRouteCmd.Flags().String("application-identifier", "", "The ID of the application within which the route is being created.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("default-route", "", "Configuration for the default route type.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("environment-identifier", "", "The ID of the environment in which the route is created.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("route-type", "", "The route type of the route.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("service-identifier", "", "The ID of the service in which the route is created.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("tags", "", "The tags to assign to the route.")
	migrationHubRefactorSpaces_createRouteCmd.Flags().String("uri-path-route", "", "The configuration for the URI path route type.")
	migrationHubRefactorSpaces_createRouteCmd.MarkFlagRequired("application-identifier")
	migrationHubRefactorSpaces_createRouteCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpaces_createRouteCmd.MarkFlagRequired("route-type")
	migrationHubRefactorSpaces_createRouteCmd.MarkFlagRequired("service-identifier")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_createRouteCmd)
}
