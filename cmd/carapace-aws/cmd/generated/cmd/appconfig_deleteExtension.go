package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteExtensionCmd = &cobra.Command{
	Use:   "delete-extension",
	Short: "Deletes an AppConfig extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteExtensionCmd).Standalone()

	appconfig_deleteExtensionCmd.Flags().String("extension-identifier", "", "The name, ID, or Amazon Resource Name (ARN) of the extension you want to delete.")
	appconfig_deleteExtensionCmd.Flags().String("version-number", "", "A specific version of an extension to delete.")
	appconfig_deleteExtensionCmd.MarkFlagRequired("extension-identifier")
	appconfigCmd.AddCommand(appconfig_deleteExtensionCmd)
}
