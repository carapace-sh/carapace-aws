package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listExtensionAssociationsCmd = &cobra.Command{
	Use:   "list-extension-associations",
	Short: "Lists all AppConfig extension associations in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listExtensionAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_listExtensionAssociationsCmd).Standalone()

		appconfig_listExtensionAssociationsCmd.Flags().String("extension-identifier", "", "The name, the ID, or the Amazon Resource Name (ARN) of the extension.")
		appconfig_listExtensionAssociationsCmd.Flags().String("extension-version-number", "", "The version number for the extension defined in the association.")
		appconfig_listExtensionAssociationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		appconfig_listExtensionAssociationsCmd.Flags().String("next-token", "", "A token to start the list.")
		appconfig_listExtensionAssociationsCmd.Flags().String("resource-identifier", "", "The ARN of an application, configuration profile, or environment.")
	})
	appconfigCmd.AddCommand(appconfig_listExtensionAssociationsCmd)
}
