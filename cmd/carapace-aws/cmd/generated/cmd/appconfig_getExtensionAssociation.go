package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getExtensionAssociationCmd = &cobra.Command{
	Use:   "get-extension-association",
	Short: "Returns information about an AppConfig extension association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getExtensionAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getExtensionAssociationCmd).Standalone()

		appconfig_getExtensionAssociationCmd.Flags().String("extension-association-id", "", "The extension association ID to get.")
		appconfig_getExtensionAssociationCmd.MarkFlagRequired("extension-association-id")
	})
	appconfigCmd.AddCommand(appconfig_getExtensionAssociationCmd)
}
