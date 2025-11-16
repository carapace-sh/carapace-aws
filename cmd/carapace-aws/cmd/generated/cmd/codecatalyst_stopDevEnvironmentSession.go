package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_stopDevEnvironmentSessionCmd = &cobra.Command{
	Use:   "stop-dev-environment-session",
	Short: "Stops a session for a specified Dev Environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_stopDevEnvironmentSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_stopDevEnvironmentSessionCmd).Standalone()

		codecatalyst_stopDevEnvironmentSessionCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment.")
		codecatalyst_stopDevEnvironmentSessionCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_stopDevEnvironmentSessionCmd.Flags().String("session-id", "", "The system-generated unique ID of the Dev Environment session.")
		codecatalyst_stopDevEnvironmentSessionCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_stopDevEnvironmentSessionCmd.MarkFlagRequired("id")
		codecatalyst_stopDevEnvironmentSessionCmd.MarkFlagRequired("project-name")
		codecatalyst_stopDevEnvironmentSessionCmd.MarkFlagRequired("session-id")
		codecatalyst_stopDevEnvironmentSessionCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_stopDevEnvironmentSessionCmd)
}
