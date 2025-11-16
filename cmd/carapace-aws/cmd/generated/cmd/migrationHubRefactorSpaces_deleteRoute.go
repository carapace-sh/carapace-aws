package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_deleteRouteCmd = &cobra.Command{
	Use:   "delete-route",
	Short: "Deletes an Amazon Web Services Migration Hub Refactor Spaces route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_deleteRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationHubRefactorSpaces_deleteRouteCmd).Standalone()

		migrationHubRefactorSpaces_deleteRouteCmd.Flags().String("application-identifier", "", "The ID of the application to delete the route from.")
		migrationHubRefactorSpaces_deleteRouteCmd.Flags().String("environment-identifier", "", "The ID of the environment to delete the route from.")
		migrationHubRefactorSpaces_deleteRouteCmd.Flags().String("route-identifier", "", "The ID of the route to delete.")
		migrationHubRefactorSpaces_deleteRouteCmd.MarkFlagRequired("application-identifier")
		migrationHubRefactorSpaces_deleteRouteCmd.MarkFlagRequired("environment-identifier")
		migrationHubRefactorSpaces_deleteRouteCmd.MarkFlagRequired("route-identifier")
	})
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_deleteRouteCmd)
}
