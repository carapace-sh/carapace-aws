package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createExtensionCmd = &cobra.Command{
	Use:   "create-extension",
	Short: "Creates an AppConfig extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createExtensionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_createExtensionCmd).Standalone()

		appconfig_createExtensionCmd.Flags().String("actions", "", "The actions defined in the extension.")
		appconfig_createExtensionCmd.Flags().String("description", "", "Information about the extension.")
		appconfig_createExtensionCmd.Flags().String("latest-version-number", "", "You can omit this field when you create an extension.")
		appconfig_createExtensionCmd.Flags().String("name", "", "A name for the extension.")
		appconfig_createExtensionCmd.Flags().String("parameters", "", "The parameters accepted by the extension.")
		appconfig_createExtensionCmd.Flags().String("tags", "", "Adds one or more tags for the specified extension.")
		appconfig_createExtensionCmd.MarkFlagRequired("actions")
		appconfig_createExtensionCmd.MarkFlagRequired("name")
	})
	appconfigCmd.AddCommand(appconfig_createExtensionCmd)
}
