package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_startDevEnvironmentCmd = &cobra.Command{
	Use:   "start-dev-environment",
	Short: "Starts a specified Dev Environment and puts it into an active state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_startDevEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_startDevEnvironmentCmd).Standalone()

		codecatalyst_startDevEnvironmentCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment.")
		codecatalyst_startDevEnvironmentCmd.Flags().String("ides", "", "Information about the integrated development environment (IDE) configured for a Dev Environment.")
		codecatalyst_startDevEnvironmentCmd.Flags().String("inactivity-timeout-minutes", "", "The amount of time the Dev Environment will run without any activity detected before stopping, in minutes.")
		codecatalyst_startDevEnvironmentCmd.Flags().String("instance-type", "", "The Amazon EC2 instace type to use for the Dev Environment.")
		codecatalyst_startDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_startDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_startDevEnvironmentCmd.MarkFlagRequired("id")
		codecatalyst_startDevEnvironmentCmd.MarkFlagRequired("project-name")
		codecatalyst_startDevEnvironmentCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_startDevEnvironmentCmd)
}
