package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createConfigurationProfileCmd = &cobra.Command{
	Use:   "create-configuration-profile",
	Short: "Creates a configuration profile, which is information that enables AppConfig to access the configuration source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createConfigurationProfileCmd).Standalone()

	appconfig_createConfigurationProfileCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_createConfigurationProfileCmd.Flags().String("description", "", "A description of the configuration profile.")
	appconfig_createConfigurationProfileCmd.Flags().String("kms-key-identifier", "", "The identifier for an Key Management Service key to encrypt new configuration data versions in the AppConfig hosted configuration store.")
	appconfig_createConfigurationProfileCmd.Flags().String("location-uri", "", "A URI to locate the configuration.")
	appconfig_createConfigurationProfileCmd.Flags().String("name", "", "A name for the configuration profile.")
	appconfig_createConfigurationProfileCmd.Flags().String("retrieval-role-arn", "", "The ARN of an IAM role with permission to access the configuration at the specified `LocationUri`.")
	appconfig_createConfigurationProfileCmd.Flags().String("tags", "", "Metadata to assign to the configuration profile.")
	appconfig_createConfigurationProfileCmd.Flags().String("type", "", "The type of configurations contained in the profile.")
	appconfig_createConfigurationProfileCmd.Flags().String("validators", "", "A list of methods for validating the configuration.")
	appconfig_createConfigurationProfileCmd.MarkFlagRequired("application-id")
	appconfig_createConfigurationProfileCmd.MarkFlagRequired("location-uri")
	appconfig_createConfigurationProfileCmd.MarkFlagRequired("name")
	appconfigCmd.AddCommand(appconfig_createConfigurationProfileCmd)
}
