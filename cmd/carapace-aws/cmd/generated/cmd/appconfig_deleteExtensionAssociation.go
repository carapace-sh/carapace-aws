package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteExtensionAssociationCmd = &cobra.Command{
	Use:   "delete-extension-association",
	Short: "Deletes an extension association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteExtensionAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_deleteExtensionAssociationCmd).Standalone()

		appconfig_deleteExtensionAssociationCmd.Flags().String("extension-association-id", "", "The ID of the extension association to delete.")
		appconfig_deleteExtensionAssociationCmd.MarkFlagRequired("extension-association-id")
	})
	appconfigCmd.AddCommand(appconfig_deleteExtensionAssociationCmd)
}
