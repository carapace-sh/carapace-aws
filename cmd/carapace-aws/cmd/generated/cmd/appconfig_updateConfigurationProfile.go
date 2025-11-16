package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateConfigurationProfileCmd = &cobra.Command{
	Use:   "update-configuration-profile",
	Short: "Updates a configuration profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateConfigurationProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_updateConfigurationProfileCmd).Standalone()

		appconfig_updateConfigurationProfileCmd.Flags().String("application-id", "", "The application ID.")
		appconfig_updateConfigurationProfileCmd.Flags().String("configuration-profile-id", "", "The ID of the configuration profile.")
		appconfig_updateConfigurationProfileCmd.Flags().String("description", "", "A description of the configuration profile.")
		appconfig_updateConfigurationProfileCmd.Flags().String("kms-key-identifier", "", "The identifier for a Key Management Service key to encrypt new configuration data versions in the AppConfig hosted configuration store.")
		appconfig_updateConfigurationProfileCmd.Flags().String("name", "", "The name of the configuration profile.")
		appconfig_updateConfigurationProfileCmd.Flags().String("retrieval-role-arn", "", "The ARN of an IAM role with permission to access the configuration at the specified `LocationUri`.")
		appconfig_updateConfigurationProfileCmd.Flags().String("validators", "", "A list of methods for validating the configuration.")
		appconfig_updateConfigurationProfileCmd.MarkFlagRequired("application-id")
		appconfig_updateConfigurationProfileCmd.MarkFlagRequired("configuration-profile-id")
	})
	appconfigCmd.AddCommand(appconfig_updateConfigurationProfileCmd)
}
