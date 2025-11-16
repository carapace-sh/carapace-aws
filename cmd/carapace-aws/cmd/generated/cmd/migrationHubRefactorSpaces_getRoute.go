package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationHubRefactorSpaces_getRouteCmd = &cobra.Command{
	Use:   "get-route",
	Short: "Gets an Amazon Web Services Migration Hub Refactor Spaces route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationHubRefactorSpaces_getRouteCmd).Standalone()

	migrationHubRefactorSpaces_getRouteCmd.Flags().String("application-identifier", "", "The ID of the application.")
	migrationHubRefactorSpaces_getRouteCmd.Flags().String("environment-identifier", "", "The ID of the environment.")
	migrationHubRefactorSpaces_getRouteCmd.Flags().String("route-identifier", "", "The ID of the route.")
	migrationHubRefactorSpaces_getRouteCmd.MarkFlagRequired("application-identifier")
	migrationHubRefactorSpaces_getRouteCmd.MarkFlagRequired("environment-identifier")
	migrationHubRefactorSpaces_getRouteCmd.MarkFlagRequired("route-identifier")
	migrationHubRefactorSpacesCmd.AddCommand(migrationHubRefactorSpaces_getRouteCmd)
}
