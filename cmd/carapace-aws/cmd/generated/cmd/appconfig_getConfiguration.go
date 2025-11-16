package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getConfigurationCmd = &cobra.Command{
	Use:   "get-configuration",
	Short: "(Deprecated) Retrieves the latest deployed configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getConfigurationCmd).Standalone()

		appconfig_getConfigurationCmd.Flags().String("application", "", "The application to get.")
		appconfig_getConfigurationCmd.Flags().String("client-configuration-version", "", "The configuration version returned in the most recent [GetConfiguration]() response.")
		appconfig_getConfigurationCmd.Flags().String("client-id", "", "The clientId parameter in the following command is a unique, user-specified ID to identify the client for the configuration.")
		appconfig_getConfigurationCmd.Flags().String("configuration", "", "The configuration to get.")
		appconfig_getConfigurationCmd.Flags().String("environment", "", "The environment to get.")
		appconfig_getConfigurationCmd.MarkFlagRequired("application")
		appconfig_getConfigurationCmd.MarkFlagRequired("client-id")
		appconfig_getConfigurationCmd.MarkFlagRequired("configuration")
		appconfig_getConfigurationCmd.MarkFlagRequired("environment")
	})
	appconfigCmd.AddCommand(appconfig_getConfigurationCmd)
}
