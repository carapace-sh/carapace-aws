package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listConfigurationProfilesCmd = &cobra.Command{
	Use:   "list-configuration-profiles",
	Short: "Lists the configuration profiles for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listConfigurationProfilesCmd).Standalone()

	appconfig_listConfigurationProfilesCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_listConfigurationProfilesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	appconfig_listConfigurationProfilesCmd.Flags().String("next-token", "", "A token to start the list.")
	appconfig_listConfigurationProfilesCmd.Flags().String("type", "", "A filter based on the type of configurations that the configuration profile contains.")
	appconfig_listConfigurationProfilesCmd.MarkFlagRequired("application-id")
	appconfigCmd.AddCommand(appconfig_listConfigurationProfilesCmd)
}
