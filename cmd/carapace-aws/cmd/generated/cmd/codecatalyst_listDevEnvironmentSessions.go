package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listDevEnvironmentSessionsCmd = &cobra.Command{
	Use:   "list-dev-environment-sessions",
	Short: "Retrieves a list of active sessions for a Dev Environment in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listDevEnvironmentSessionsCmd).Standalone()

	codecatalyst_listDevEnvironmentSessionsCmd.Flags().String("dev-environment-id", "", "The system-generated unique ID of the Dev Environment.")
	codecatalyst_listDevEnvironmentSessionsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
	codecatalyst_listDevEnvironmentSessionsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
	codecatalyst_listDevEnvironmentSessionsCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_listDevEnvironmentSessionsCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_listDevEnvironmentSessionsCmd.MarkFlagRequired("dev-environment-id")
	codecatalyst_listDevEnvironmentSessionsCmd.MarkFlagRequired("project-name")
	codecatalyst_listDevEnvironmentSessionsCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_listDevEnvironmentSessionsCmd)
}
