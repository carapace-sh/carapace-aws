package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getExtensionCmd = &cobra.Command{
	Use:   "get-extension",
	Short: "Returns information about an AppConfig extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getExtensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getExtensionCmd).Standalone()

		appconfig_getExtensionCmd.Flags().String("extension-identifier", "", "The name, the ID, or the Amazon Resource Name (ARN) of the extension.")
		appconfig_getExtensionCmd.Flags().String("version-number", "", "The extension version number.")
		appconfig_getExtensionCmd.MarkFlagRequired("extension-identifier")
	})
	appconfigCmd.AddCommand(appconfig_getExtensionCmd)
}
