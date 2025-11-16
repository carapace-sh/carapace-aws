package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateEnvironmentCmd = &cobra.Command{
	Use:   "update-environment",
	Short: "Updates an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_updateEnvironmentCmd).Standalone()

		appconfig_updateEnvironmentCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_updateEnvironmentCmd.Flags().String("description", "", "A description of the environment.")
		appconfig_updateEnvironmentCmd.Flags().String("environment-id", "", "The environment ID.")
		appconfig_updateEnvironmentCmd.Flags().String("monitors", "", "Amazon CloudWatch alarms to monitor during the deployment process.")
		appconfig_updateEnvironmentCmd.Flags().String("name", "", "The name of the environment.")
		appconfig_updateEnvironmentCmd.MarkFlagRequired("application-id")
		appconfig_updateEnvironmentCmd.MarkFlagRequired("environment-id")
	})
	appconfigCmd.AddCommand(appconfig_updateEnvironmentCmd)
}
