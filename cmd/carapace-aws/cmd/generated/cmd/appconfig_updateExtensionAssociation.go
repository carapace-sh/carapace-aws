package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateExtensionAssociationCmd = &cobra.Command{
	Use:   "update-extension-association",
	Short: "Updates an association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateExtensionAssociationCmd).Standalone()

	appconfig_updateExtensionAssociationCmd.Flags().String("extension-association-id", "", "The system-generated ID for the association.")
	appconfig_updateExtensionAssociationCmd.Flags().String("parameters", "", "The parameter names and values defined in the extension.")
	appconfig_updateExtensionAssociationCmd.MarkFlagRequired("extension-association-id")
	appconfigCmd.AddCommand(appconfig_updateExtensionAssociationCmd)
}
