package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_validateConfigurationCmd = &cobra.Command{
	Use:   "validate-configuration",
	Short: "Uses the validators in a configuration profile to validate a configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_validateConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_validateConfigurationCmd).Standalone()

		appconfig_validateConfigurationCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_validateConfigurationCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
		appconfig_validateConfigurationCmd.Flags().String("configuration-version", "", "The version of the configuration to validate.")
		appconfig_validateConfigurationCmd.MarkFlagRequired("application-id")
		appconfig_validateConfigurationCmd.MarkFlagRequired("configuration-profile-id")
		appconfig_validateConfigurationCmd.MarkFlagRequired("configuration-version")
	})
	appconfigCmd.AddCommand(appconfig_validateConfigurationCmd)
}
