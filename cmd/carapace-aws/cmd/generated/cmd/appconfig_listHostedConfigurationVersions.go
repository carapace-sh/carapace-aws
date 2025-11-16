package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listHostedConfigurationVersionsCmd = &cobra.Command{
	Use:   "list-hosted-configuration-versions",
	Short: "Lists configurations stored in the AppConfig hosted configuration store by version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listHostedConfigurationVersionsCmd).Standalone()

	appconfig_listHostedConfigurationVersionsCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_listHostedConfigurationVersionsCmd.Flags().String("configuration-profile-id", "", "The configuration profile ID.")
	appconfig_listHostedConfigurationVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	appconfig_listHostedConfigurationVersionsCmd.Flags().String("next-token", "", "A token to start the list.")
	appconfig_listHostedConfigurationVersionsCmd.Flags().String("version-label", "", "An optional filter that can be used to specify the version label of an AppConfig hosted configuration version.")
	appconfig_listHostedConfigurationVersionsCmd.MarkFlagRequired("application-id")
	appconfig_listHostedConfigurationVersionsCmd.MarkFlagRequired("configuration-profile-id")
	appconfigCmd.AddCommand(appconfig_listHostedConfigurationVersionsCmd)
}
