package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateExtensionCmd = &cobra.Command{
	Use:   "update-extension",
	Short: "Updates an AppConfig extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateExtensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_updateExtensionCmd).Standalone()

		appconfig_updateExtensionCmd.Flags().String("actions", "", "The actions defined in the extension.")
		appconfig_updateExtensionCmd.Flags().String("description", "", "Information about the extension.")
		appconfig_updateExtensionCmd.Flags().String("extension-identifier", "", "The name, the ID, or the Amazon Resource Name (ARN) of the extension.")
		appconfig_updateExtensionCmd.Flags().String("parameters", "", "One or more parameters for the actions called by the extension.")
		appconfig_updateExtensionCmd.Flags().String("version-number", "", "The extension version number.")
		appconfig_updateExtensionCmd.MarkFlagRequired("extension-identifier")
	})
	appconfigCmd.AddCommand(appconfig_updateExtensionCmd)
}
