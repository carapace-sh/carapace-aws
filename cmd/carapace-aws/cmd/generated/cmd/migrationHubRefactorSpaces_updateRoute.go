package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_updateRouteCmd = &cobra.Command{
	Use:   "update-route",
	Short: "Updates an Amazon Web Services Migration Hub Refactor Spaces route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_updateRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_updateRouteCmd).Standalone()

		migrationHubRefactorSpaces_updateRouteCmd.Flags().String("activation-state", "", "If set to `ACTIVE`, traffic is forwarded to this route’s service after the route is updated.")
		migrationHubRefactorSpaces_updateRouteCmd.Flags().String("application-identifier", "", "The ID of the application within which the route is being updated.")
		migrationHubRefactorSpaces_updateRouteCmd.Flags().String("environment-identifier", "", "The ID of the environment in which the route is being updated.")
		migrationHubRefactorSpaces_updateRouteCmd.Flags().String("route-identifier", "", "The unique identifier of the route to update.")
		migrationHubRefactorSpaces_updateRouteCmd.MarkFlagRequired("activation-state")
		migrationHubRefactorSpaces_updateRouteCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_updateRouteCmd.MarkFlagRequired("environment-identifier")
		migrationHubRefactorSpaces_updateRouteCmd.MarkFlagRequired("route-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_updateRouteCmd)
}
