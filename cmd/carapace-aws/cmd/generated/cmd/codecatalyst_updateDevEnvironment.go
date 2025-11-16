package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_updateDevEnvironmentCmd = &cobra.Command{
	Use:   "update-dev-environment",
	Short: "Changes one or more values for a Dev Environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_updateDevEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_updateDevEnvironmentCmd).Standalone()

		codecatalyst_updateDevEnvironmentCmd.Flags().String("alias", "", "The user-specified alias for the Dev Environment.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("client-token", "", "A user-specified idempotency token.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("ides", "", "Information about the integrated development environment (IDE) configured for a Dev Environment.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("inactivity-timeout-minutes", "", "The amount of time the Dev Environment will run without any activity detected before stopping, in minutes.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("instance-type", "", "The Amazon EC2 instace type to use for the Dev Environment.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_updateDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_updateDevEnvironmentCmd.MarkFlagRequired("id")
		codecatalyst_updateDevEnvironmentCmd.MarkFlagRequired("project-name")
		codecatalyst_updateDevEnvironmentCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_updateDevEnvironmentCmd)
}
