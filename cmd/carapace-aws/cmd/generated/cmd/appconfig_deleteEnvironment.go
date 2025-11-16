package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_deleteEnvironmentCmd).Standalone()

		appconfig_deleteEnvironmentCmd.Flags().String("application-id", "", "The application ID that includes the environment that you want to delete.")
		appconfig_deleteEnvironmentCmd.Flags().String("deletion-protection-check", "", "A parameter to configure deletion protection.")
		appconfig_deleteEnvironmentCmd.Flags().String("environment-id", "", "The ID of the environment that you want to delete.")
		appconfig_deleteEnvironmentCmd.MarkFlagRequired("application-id")
		appconfig_deleteEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	appconfigCmd.AddCommand(appconfig_deleteEnvironmentCmd)
}
