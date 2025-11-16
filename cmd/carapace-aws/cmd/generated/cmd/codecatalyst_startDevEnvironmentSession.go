package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_startDevEnvironmentSessionCmd = &cobra.Command{
	Use:   "start-dev-environment-session",
	Short: "Starts a session for a specified Dev Environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_startDevEnvironmentSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_startDevEnvironmentSessionCmd).Standalone()

		codecatalyst_startDevEnvironmentSessionCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment.")
		codecatalyst_startDevEnvironmentSessionCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_startDevEnvironmentSessionCmd.Flags().String("session-configuration", "", "")
		codecatalyst_startDevEnvironmentSessionCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_startDevEnvironmentSessionCmd.MarkFlagRequired("id")
		codecatalyst_startDevEnvironmentSessionCmd.MarkFlagRequired("project-name")
		codecatalyst_startDevEnvironmentSessionCmd.MarkFlagRequired("session-configuration")
		codecatalyst_startDevEnvironmentSessionCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_startDevEnvironmentSessionCmd)
}
