package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteConfigurationProfileCmd = &cobra.Command{
	Use:   "delete-configuration-profile",
	Short: "Deletes a configuration profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteConfigurationProfileCmd).Standalone()

	appconfig_deleteConfigurationProfileCmd.Flags().String("application-id", "", "The application ID that includes the configuration profile you want to delete.")
	appconfig_deleteConfigurationProfileCmd.Flags().String("configuration-profile-id", "", "The ID of the configuration profile you want to delete.")
	appconfig_deleteConfigurationProfileCmd.Flags().String("deletion-protection-check", "", "A parameter to configure deletion protection.")
	appconfig_deleteConfigurationProfileCmd.MarkFlagRequired("application-id")
	appconfig_deleteConfigurationProfileCmd.MarkFlagRequired("configuration-profile-id")
	appconfigCmd.AddCommand(appconfig_deleteConfigurationProfileCmd)
}
