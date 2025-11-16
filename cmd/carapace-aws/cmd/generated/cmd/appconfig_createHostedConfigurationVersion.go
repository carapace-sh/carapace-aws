package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createHostedConfigurationVersionCmd = &cobra.Command{
	Use:   "create-hosted-configuration-version",
	Short: "Creates a new configuration in the AppConfig hosted configuration store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createHostedConfigurationVersionCmd).Standalone()

	appconfig_createHostedConfigurationVersionCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("content", "", "The configuration data, as bytes.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("content-type", "", "A standard MIME type describing the format of the configuration content.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("description", "", "A description of the configuration.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("latest-version-number", "", "An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version.")
	appconfig_createHostedConfigurationVersionCmd.Flags().String("version-label", "", "An optional, user-defined label for the AppConfig hosted configuration version.")
	appconfig_createHostedConfigurationVersionCmd.MarkFlagRequired("application-id")
	appconfig_createHostedConfigurationVersionCmd.MarkFlagRequired("configuration-profile-id")
	appconfig_createHostedConfigurationVersionCmd.MarkFlagRequired("content")
	appconfig_createHostedConfigurationVersionCmd.MarkFlagRequired("content-type")
	appconfigCmd.AddCommand(appconfig_createHostedConfigurationVersionCmd)
}
