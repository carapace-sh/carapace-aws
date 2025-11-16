package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createExtensionAssociationCmd = &cobra.Command{
	Use:   "create-extension-association",
	Short: "When you create an extension or configure an Amazon Web Services authored extension, you associate the extension with an AppConfig application, environment, or configuration profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createExtensionAssociationCmd).Standalone()

	appconfig_createExtensionAssociationCmd.Flags().String("extension-identifier", "", "The name, the ID, or the Amazon Resource Name (ARN) of the extension.")
	appconfig_createExtensionAssociationCmd.Flags().String("extension-version-number", "", "The version number of the extension.")
	appconfig_createExtensionAssociationCmd.Flags().String("parameters", "", "The parameter names and values defined in the extensions.")
	appconfig_createExtensionAssociationCmd.Flags().String("resource-identifier", "", "The ARN of an application, configuration profile, or environment.")
	appconfig_createExtensionAssociationCmd.Flags().String("tags", "", "Adds one or more tags for the specified extension association.")
	appconfig_createExtensionAssociationCmd.MarkFlagRequired("extension-identifier")
	appconfig_createExtensionAssociationCmd.MarkFlagRequired("resource-identifier")
	appconfigCmd.AddCommand(appconfig_createExtensionAssociationCmd)
}
