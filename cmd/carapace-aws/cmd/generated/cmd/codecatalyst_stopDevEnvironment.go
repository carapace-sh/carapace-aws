package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_stopDevEnvironmentCmd = &cobra.Command{
	Use:   "stop-dev-environment",
	Short: "Pauses a specified Dev Environment and places it in a non-running state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_stopDevEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_stopDevEnvironmentCmd).Standalone()

		codecatalyst_stopDevEnvironmentCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment.")
		codecatalyst_stopDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_stopDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_stopDevEnvironmentCmd.MarkFlagRequired("id")
		codecatalyst_stopDevEnvironmentCmd.MarkFlagRequired("project-name")
		codecatalyst_stopDevEnvironmentCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_stopDevEnvironmentCmd)
}
