package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getHostedConfigurationVersionCmd = &cobra.Command{
	Use:   "get-hosted-configuration-version",
	Short: "Retrieves information about a specific configuration version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getHostedConfigurationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getHostedConfigurationVersionCmd).Standalone()

		appconfig_getHostedConfigurationVersionCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_getHostedConfigurationVersionCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
		appconfig_getHostedConfigurationVersionCmd.Flags().String("version-number", "", "The version.")
		appconfig_getHostedConfigurationVersionCmd.MarkFlagRequired("application-id")
		appconfig_getHostedConfigurationVersionCmd.MarkFlagRequired("configuration-profile-id")
		appconfig_getHostedConfigurationVersionCmd.MarkFlagRequired("version-number")
	})
	appconfigCmd.AddCommand(appconfig_getHostedConfigurationVersionCmd)
}
