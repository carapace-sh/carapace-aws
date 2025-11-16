package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteHostedConfigurationVersionCmd = &cobra.Command{
	Use:   "delete-hosted-configuration-version",
	Short: "Deletes a version of a configuration from the AppConfig hosted configuration store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteHostedConfigurationVersionCmd).Standalone()

	appconfig_deleteHostedConfigurationVersionCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_deleteHostedConfigurationVersionCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
	appconfig_deleteHostedConfigurationVersionCmd.Flags().String("version-number", "", "The versions number to delete.")
	appconfig_deleteHostedConfigurationVersionCmd.MarkFlagRequired("application-id")
	appconfig_deleteHostedConfigurationVersionCmd.MarkFlagRequired("configuration-profile-id")
	appconfig_deleteHostedConfigurationVersionCmd.MarkFlagRequired("version-number")
	appconfigCmd.AddCommand(appconfig_deleteHostedConfigurationVersionCmd)
}
