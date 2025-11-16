package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getConfigurationProfileCmd = &cobra.Command{
	Use:   "get-configuration-profile",
	Short: "Retrieves information about a configuration profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getConfigurationProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getConfigurationProfileCmd).Standalone()

		appconfig_getConfigurationProfileCmd.Flags().String("application-id", "", "The ID of the application that includes the configuration profile you want to get.")
		appconfig_getConfigurationProfileCmd.Flags().String("configuration-profile-id", "", "The ID of the configuration profile that you want to get.")
		appconfig_getConfigurationProfileCmd.MarkFlagRequired("application-id")
		appconfig_getConfigurationProfileCmd.MarkFlagRequired("configuration-profile-id")
	})
	appconfigCmd.AddCommand(appconfig_getConfigurationProfileCmd)
}
